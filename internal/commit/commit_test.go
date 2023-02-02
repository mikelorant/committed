package commit_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"

	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	openErr error
	descErr error
}

func (r *MockRepository) Open() error {
	return r.openErr
}

func (r *MockRepository) Describe() (repository.Description, error) {
	return repository.Description{}, r.descErr
}

type MockConfig struct {
	err error
}

func (c *MockConfig) Load(fh io.Reader) (config.Config, error) {
	return config.Config{}, c.err
}

type MockApply struct {
	commit repository.Commit
	err    error
}

func (a *MockApply) Apply() func(c repository.Commit, opts ...func(c *repository.Commit)) error {
	return func(c repository.Commit, opts ...func(c *repository.Commit)) error {
		a.commit = c

		for _, o := range opts {
			o(&a.commit)
		}

		if a.err != nil {
			return a.err
		}

		return nil
	}
}

func MockNewRepository(err error) func() (*repository.Repository, error) {
	return func() (*repository.Repository, error) {
		return nil, err
	}
}

func MockNewEmoji(opts ...func(*emoji.Set)) *emoji.Set {
	return &emoji.Set{}
}

func MockOpen(err error) func(string) (io.Reader, error) {
	return func(file string) (io.Reader, error) {
		return strings.NewReader(""), err
	}
}

var errMock = errors.New("error")

func TestConfigure(t *testing.T) {
	type args struct {
		opts        commit.Options
		repoOpenErr error
		repoDescErr error
		configErr   error
		openErr     error
	}

	type want struct {
		state commit.State
		err   string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			want: want{
				state: commit.State{
					Placeholders: testPlaceholders(),
					Config:       config.Config{},
					Emojis:       &emoji.Set{},
				},
			},
		},
		{
			name: "amend",
			args: args{
				opts: commit.Options{
					Amend: true,
				},
			},
			want: want{
				state: commit.State{
					Placeholders: testPlaceholders(),
					Config:       config.Config{},
					Emojis:       &emoji.Set{},
					Options: commit.Options{
						Amend: true,
					},
				},
			},
		},
		{
			name: "config_file",
			args: args{
				opts: commit.Options{
					ConfigFile: "test",
				},
			},
			want: want{
				state: commit.State{
					Placeholders: testPlaceholders(),
					Config:       config.Config{},
					Emojis:       &emoji.Set{},
					Options: commit.Options{
						ConfigFile: "test",
					},
				},
			},
		},
		{
			name: "open_error",
			args: args{
				repoOpenErr: errMock,
			},
			want: want{
				err: "unable to get repository: unable to open repository: error",
			},
		},
		{
			name: "describe_error",
			args: args{
				repoDescErr: errMock,
			},
			want: want{
				err: "unable to get repository: unable to describe repository: error",
			},
		},
		{
			name: "open_error",
			args: args{
				opts: commit.Options{
					ConfigFile: "test",
				},
				openErr: errMock,
			},
			want: want{
				err: "unable to get config: unable to open config file: test: error",
			},
		},
		{
			name: "config_error",
			args: args{
				configErr: errMock,
			},
			want: want{
				err: "unable to get config: unable to load config file: error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := MockConfig{
				err: tt.args.configErr,
			}

			repo := MockRepository{
				openErr: tt.args.repoOpenErr,
				descErr: tt.args.repoDescErr,
			}

			c := commit.Commit{
				Repoer:  &repo,
				Loader:  config.Load,
				Emojier: MockNewEmoji,
				Opener:  MockOpen(tt.args.openErr),
			}

			state, err := c.Configure(tt.args.opts)
			if tt.want.err != "" {
				assert.NotNil(t, err)
				assert.Equal(t, tt.want.err, err.Error())
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, &tt.want.state, state)
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		emoji   string
		summary string
		body    string
		footer  string
		author  repository.User
		amend   bool
		options commit.Options
	}

	type want struct {
		author  string
		subject string
		body    string
		footer  string
		amend   bool
		dryRun  bool
		err     error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "normal",
			args: args{
				emoji:   ":art:",
				summary: "summary",
				body:    "body",
				footer:  "Signed-off-by: John Doe <john.doe@example.com>",
				author: repository.User{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				},
			},
			want: want{
				author:  "John Doe <john.doe@example.com>",
				subject: ":art: summary",
				body:    "body",
				footer:  "Signed-off-by: John Doe <john.doe@example.com>",
			},
		},
		{
			name: "dryrun",
			args: args{
				options: commit.Options{
					DryRun: true,
				},
			},
			want: want{
				dryRun: true,
			},
		},
		{
			name: "amend",
			args: args{
				amend: true,
			},
			want: want{
				amend: true,
			},
		},
		{
			name: "amend_dryrun",
			args: args{
				amend: true,
				options: commit.Options{
					DryRun: true,
				},
			},
			want: want{
				amend:  true,
				dryRun: true,
			},
		},
		{
			name: "invalid",
			want: want{
				err: errMock,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := MockApply{
				err: tt.want.err,
			}

			req := &commit.Request{
				Emoji:   tt.args.emoji,
				Summary: tt.args.summary,
				Body:    tt.args.body,
				Footer:  tt.args.footer,
				Author:  tt.args.author,
				Amend:   tt.args.amend,
			}

			c := commit.Commit{
				Options: tt.args.options,
				Applier: a.Apply(),
			}

			err := c.Apply(req)
			if tt.want.err != nil {
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want.author, a.commit.Author)
			assert.Equal(t, tt.want.subject, a.commit.Subject)
			assert.Equal(t, tt.want.body, a.commit.Body)
			assert.Equal(t, tt.want.footer, a.commit.Footer)
			assert.Equal(t, tt.want.amend, a.commit.Amend)
			assert.Equal(t, tt.want.dryRun, a.commit.DryRun)
		})
	}
}

func testPlaceholders() commit.Placeholders {
	return commit.Placeholders{
		Hash:    commit.PlaceholderHash,
		Summary: commit.PlaceholderSummary,
		Body:    commit.PlaceholderMessage,
	}
}
