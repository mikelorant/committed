package repository_test

import (
	"errors"
	"io"
	"os"
	"strings"
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

type MockOpenFile struct {
	filename     string
	mockFilename string
	err          error
	close        bool
}

func (o *MockOpenFile) OpenFile() func(string, int, os.FileMode) (*os.File, error) {
	return func(filename string, flag int, perm os.FileMode) (*os.File, error) {
		if o.err != nil {
			return nil, o.err
		}

		fh, _ := os.CreateTemp("", "")

		o.filename = filename
		o.mockFilename = fh.Name()

		if o.close {
			fh.Close()
		}

		return fh, nil
	}
}

var errMock = errors.New("error")

func TestApply(t *testing.T) {
	t.Parallel()

	type args struct {
		commit      repository.Commit
		opts        []func(c *repository.Commit)
		filename    string
		runErr      error
		openFileErr error
		close       bool
	}

	type want struct {
		cmd          string
		args         []string
		data         string
		mockFilename string
		err          string
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
			name: "hook",
			args: args{
				commit: repository.Commit{
					Hook:        true,
					MessageFile: "test",
				},
			},
			want: want{
				mockFilename: "test",
			},
		},
		{
			name: "hook_summary",
			args: args{
				commit: repository.Commit{
					Subject: "summary",
					Hook:    true,
				},
			},
			want: want{
				data: "summary",
			},
		},
		{
			name: "hook_summary_body",
			args: args{
				commit: repository.Commit{
					Subject: "summary",
					Body:    "body",
					Hook:    true,
				},
			},
			want: want{
				data: "summary\n\nbody",
			},
		},
		{
			name: "hook_summary_footer",
			args: args{
				commit: repository.Commit{
					Subject: "summary",
					Footer:  "Signed-off-by: John Doe <john.doe@example.com>",
					Hook:    true,
				},
			},
			want: want{
				data: "summary\n\nSigned-off-by: John Doe <john.doe@example.com>",
			},
		},
		{
			name: "hook_summary_body_footer",
			args: args{
				commit: repository.Commit{
					Subject: "summary",
					Body:    "body",
					Footer:  "Signed-off-by: John Doe <john.doe@example.com>",
					Hook:    true,
				},
			},
			want: want{
				data: "summary\n\nbody\n\nSigned-off-by: John Doe <john.doe@example.com>",
			},
		},
		{
			name: "run_error",
			args: args{
				runErr: errMock,
			},
			want: want{
				err: "unable to run command: error",
			},
		},
		{
			name: "hook_error",
			args: args{
				commit: repository.Commit{
					Hook: true,
				},
				openFileErr: errMock,
			},
			want: want{
				err: "unble to open file: error",
			},
		},
		{
			name: "hook_error_close",
			args: args{
				commit: repository.Commit{
					Hook: true,
				},
				close: true,
			},
			want: want{
				err: "unable to write file: unable to close file: close",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			shell := MockShell{
				err: tt.args.runErr,
			}

			openFile := MockOpenFile{
				filename: tt.args.filename,
				err:      tt.args.openFileErr,
				close:    tt.args.close,
			}

			repo := repository.Repository{
				Runner:    shell.Run(),
				OpenFiler: openFile.OpenFile(),
			}

			err := repo.Apply(tt.args.commit)
			if tt.want.err != "" {
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}
			assert.Nil(t, err)

			assert.Equal(t, tt.want.cmd, shell.command)
			assert.Equal(t, tt.want.args, shell.args)

			if tt.args.commit.Hook {
				out, _ := os.ReadFile(openFile.mockFilename)
				assert.Equal(t, tt.want.mockFilename, openFile.filename)
				assert.Equal(t, tt.want.data, strings.TrimSpace(string(out)))
			}
		})
	}
}
