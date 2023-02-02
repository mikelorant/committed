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
	commit repository.Commit

	openErr  error
	descErr  error
	applyErr error
}

func (r *MockRepository) Open() error {
	return r.openErr
}

func (r *MockRepository) Describe() (repository.Description, error) {
	return repository.Description{}, r.descErr
}

func (r *MockRepository) Apply(c repository.Commit, opts ...func(c *repository.Commit)) error {
	r.commit = c

	for _, o := range opts {
		o(&r.commit)
	}

	if r.applyErr != nil {
		return r.applyErr
	}

	return nil
}

type MockConfig struct {
	cfg  config.Config
	file config.Config

	loadErr error
	saveErr error
}

func (c *MockConfig) Load(fh io.Reader) (config.Config, error) {
	return c.cfg, c.loadErr
}

func (c *MockConfig) Save(fh io.WriteCloser, cfg config.Config) error {
	c.file = cfg

	return c.saveErr
}

type DiscardCloser struct {
	io.Writer
}

func (d DiscardCloser) Close() error {
	return nil
}

func MockNewEmoji(opts ...func(*emoji.Set)) *emoji.Set {
	return &emoji.Set{}
}

func MockOpen(err error) func(string) (io.Reader, error) {
	return func(file string) (io.Reader, error) {
		return strings.NewReader(""), err
	}
}

func MockCreate(err error) func(file string) (io.WriteCloser, error) {
	return func(file string) (io.WriteCloser, error) {
		if err != nil {
			return DiscardCloser{}, err
		}

		return DiscardCloser{}, nil
	}
}

var errMock = errors.New("error")

func TestConfigure(t *testing.T) {
	type args struct {
		opts        commit.Options
		cfg         config.Config
		repoOpenErr error
		repoDescErr error
		configErr   error
		openErr     error
		createErr   error
		loadErr     error
		saveErr     error
	}

	type want struct {
		state commit.State
		cfg   config.Config
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
			name: "dryrun",
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
			name: "save",
			args: args{
				cfg: config.Config{
					View: config.View{
						Focus: config.FocusAuthor,
					},
				},
			},
			want: want{
				state: commit.State{
					Placeholders: testPlaceholders(),
					Config: config.Config{
						View: config.View{
							Focus: config.FocusAuthor,
						},
					},
					Emojis: &emoji.Set{},
				},
				cfg: config.Config{
					View: config.View{
						Focus: config.FocusAuthor,
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
		{
			name: "create_error",
			args: args{
				createErr: errMock,
			},
			want: want{
				err: "unable to set config: unable to create config: error",
			},
		},
		{
			name: "save_error",
			args: args{
				saveErr: errMock,
			},
			want: want{
				err: "unable to set config: unable to save config: error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := MockConfig{
				cfg:     tt.args.cfg,
				loadErr: tt.args.configErr,
				saveErr: tt.args.saveErr,
			}

			repo := MockRepository{
				openErr: tt.args.repoOpenErr,
				descErr: tt.args.repoDescErr,
			}

			c := commit.Commit{
				Repoer:   &repo,
				Configer: &cfg,
				Emojier:  MockNewEmoji,
				Creator:  MockCreate(tt.args.createErr),
				Opener:   MockOpen(tt.args.openErr),
			}

			state, err := c.Configure(tt.args.opts)
			if tt.want.err != "" {
				assert.NotNil(t, err)
				assert.Equal(t, tt.want.err, err.Error())
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, &tt.want.state, state)
			assert.Equal(t, tt.want.cfg, cfg.file)
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		apply     bool
		emoji     string
		summary   string
		body      string
		footer    string
		author    repository.User
		amend     bool
		options   commit.Options
		createErr error
		saveErr   error
		applyErr  error
		nilReq    bool
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
				apply:   true,
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
				apply: true,
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
				apply: true,
				amend: true,
			},
			want: want{
				amend: true,
			},
		},
		{
			name: "amend_dryrun",
			args: args{
				apply: true,
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
			name: "save",
			args: args{
				apply: false,
			},
		},
		{
			name: "skip_apply",
			args: args{
				apply: false,
			},
		},
		{
			name: "no_request",
			args: args{
				nilReq: true,
			},
		},
		{
			name: "invalid",
			args: args{
				apply:    true,
				applyErr: errMock,
			},
			want: want{
				err: errMock,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := MockRepository{
				applyErr: tt.args.applyErr,
			}

			req := &commit.Request{
				Apply:   tt.args.apply,
				Emoji:   tt.args.emoji,
				Summary: tt.args.summary,
				Body:    tt.args.body,
				Footer:  tt.args.footer,
				Author:  tt.args.author,
				Amend:   tt.args.amend,
			}

			if tt.args.nilReq {
				req = nil
			}

			c := commit.Commit{
				Options: tt.args.options,
				Repoer:  &repo,
				Creator: MockCreate(tt.args.createErr),
			}

			err := c.Apply(req)
			if tt.want.err != nil {
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want.author, repo.commit.Author)
			assert.Equal(t, tt.want.subject, repo.commit.Subject)
			assert.Equal(t, tt.want.body, repo.commit.Body)
			assert.Equal(t, tt.want.footer, repo.commit.Footer)
			assert.Equal(t, tt.want.amend, repo.commit.Amend)
			assert.Equal(t, tt.want.dryRun, repo.commit.DryRun)
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
