package commit_test

import (
	"errors"
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

var errMock = errors.New("error")

type MockApply struct {
	commit repository.Commit
	err    error
}

func (a *MockApply) Apply() func(c repository.Commit, opts ...repository.CommitOptions) error {
	return func(c repository.Commit, opts ...repository.CommitOptions) error {
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
				options: commit.Options{
					Apply: true,
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
			name: "apply",
			args: args{
				options: commit.Options{
					Apply: true,
				},
			},
			want: want{
				dryRun: false,
			},
		},
		{
			name: "dry_run",
			args: args{
				options: commit.Options{
					Apply: false,
				},
			},
			want: want{
				dryRun: true,
			},
		},
		{
			name: "amend_apply",
			args: args{
				options: commit.Options{
					Amend: true,
					Apply: true,
				},
			},
			want: want{
				amend:  true,
				dryRun: false,
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

			c := commit.Commit{
				Request: commit.Request{
					Emoji:   tt.args.emoji,
					Summary: tt.args.summary,
					Body:    tt.args.body,
					Footer:  tt.args.footer,
					Author:  tt.args.author,
				},
				Options: tt.args.options,
				Applier: a.Apply(),
			}

			err := c.Apply()
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
