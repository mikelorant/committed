package repository_test

import (
	"errors"
	"io"
	"testing"

	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MockShell struct {
	command string
	args    []string

	err error
}

func (r *MockShell) Run() func(w io.Writer, command string, args []string) error {
	return func(w io.Writer, command string, args []string) error {
		r.command = command
		r.args = args

		if r.err != nil {
			return r.err
		}

		return nil
	}
}

var errMock = errors.New("error")

func TestApply(t *testing.T) {
	type args struct {
		commit repository.Commit
		opts   []func(c *repository.Commit)
		err    error
	}

	type want struct {
		cmd  string
		args []string
		err  error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "full",
			args: args{
				commit: repository.Commit{
					Author:  "John Doe <john.doe@example.com",
					Subject: ":art: summary",
					Body:    "body",
					Footer:  "Signed-off-by: John Doe <john.doe@example.com",
				},
			},
			want: want{
				cmd: "git",
				args: []string{
					"commit",
					"--author", "John Doe <john.doe@example.com",
					"--message", ":art: summary",
					"--message", "body",
					"--message", "Signed-off-by: John Doe <john.doe@example.com",
				},
			},
		},
		{
			name: "no_body",
			args: args{
				commit: repository.Commit{
					Author:  "John Doe <john.doe@example.com",
					Subject: ":art: summary",
					Footer:  "Signed-off-by: John Doe <john.doe@example.com",
				},
			},
			want: want{
				cmd: "git",
				args: []string{
					"commit",
					"--author", "John Doe <john.doe@example.com",
					"--message", ":art: summary",
					"--message", "Signed-off-by: John Doe <john.doe@example.com",
				},
			},
		},
		{
			name: "no_footer",
			args: args{
				commit: repository.Commit{
					Author:  "John Doe <john.doe@example.com",
					Subject: ":art: summary",
					Body:    "body",
				},
			},
			want: want{
				cmd: "git",
				args: []string{
					"commit",
					"--author", "John Doe <john.doe@example.com",
					"--message", ":art: summary",
					"--message", "body",
				},
			},
		},
		{
			name: "amend",
			args: args{
				commit: repository.Commit{
					Author:  "John Doe <john.doe@example.com",
					Subject: ":art: summary",
					Amend:   true,
				},
			},
			want: want{
				cmd: "git",
				args: []string{
					"commit",
					"--author", "John Doe <john.doe@example.com",
					"--message", ":art: summary",
					"--amend",
				},
			},
		},
		{
			name: "dryrun",
			args: args{
				commit: repository.Commit{
					Author:  "John Doe <john.doe@example.com",
					Subject: ":art: summary",
					DryRun:  true,
				},
			},
			want: want{
				cmd: "git",
				args: []string{
					"commit",
					"--author", "John Doe <john.doe@example.com",
					"--message", ":art: summary",
					"--dry-run",
				},
			},
		},
		{
			name: "error",
			args: args{
				err: errMock,
			},
			want: want{
				err: errMock,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var repo repository.Repository

			shell := MockShell{
				err: tt.want.err,
			}

			repo.Runner = shell.Run()

			err := repo.Apply(tt.args.commit)
			if tt.want.err != nil {
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want.cmd, shell.command)
			assert.Equal(t, tt.want.args, shell.args)
		})
	}
}
