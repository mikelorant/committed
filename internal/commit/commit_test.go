package commit_test

import (
	"errors"
	"testing"

	"github.com/mikelorant/committed/internal/commit"
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

var errMock = errors.New("error")

func TestConfigure(t *testing.T) {
	type args struct {
		opts    commit.Options
		openErr error
		descErr error
	}

	type want struct {
		cfg commit.Config
		err string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			want: want{
				cfg: commit.Config{
					Placeholders: testPlaceholders(),
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
				cfg: commit.Config{
					Placeholders: testPlaceholders(),
					Amend:        true,
				},
			},
		},
		{
			name: "open_error",
			args: args{
				openErr: errMock,
			},
			want: want{
				err: "unable to open repository: error",
			},
		},
		{
			name: "describe_error",
			args: args{
				descErr: errMock,
			},
			want: want{
				err: "unable to describe repository: error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := MockRepository{
				openErr: tt.args.openErr,
				descErr: tt.args.descErr,
			}

			c := commit.Commit{
				Repoer:  &repo,
				Emojier: MockNewEmoji,
			}

			cfg, err := c.Configure(tt.args.opts)
			if tt.want.err != "" {
				assert.NotNil(t, err)
				assert.Equal(t, tt.want.err, err.Error())
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, &tt.want.cfg, cfg)
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
				options: commit.Options{
					Amend: true,
				},
			},
			want: want{
				amend: true,
			},
		},
		{
			name: "amend_dryrun",
			args: args{
				options: commit.Options{
					Amend:  true,
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
