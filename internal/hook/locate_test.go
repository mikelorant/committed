package hook_test

import (
	"errors"
	"io"
	"testing"

	"github.com/mikelorant/committed/internal/hook"
	"github.com/stretchr/testify/assert"
)

type MockRun struct {
	idx int

	glob    string
	globErr error
	repo    string
	repoErr error

	cmd  string
	args []string
}

var errMock = errors.New("error")

func (r *MockRun) Run() func(io.Writer, string, []string) error {
	return func(w io.Writer, cmd string, args []string) error {
		r.cmd = cmd
		r.args = args

		if r.idx == 0 && r.glob != "" {
			if r.globErr != nil {
				return r.globErr
			}

			io.WriteString(w, r.glob)

			return nil
		}

		if r.idx == 1 && r.repo != "" {
			if r.repoErr != nil {
				return r.repoErr
			}

			io.WriteString(w, r.repo)

			return nil
		}

		r.idx++

		return nil
	}
}

func TestLocate(t *testing.T) {
	t.Parallel()

	type args struct {
		glob    string
		globErr error

		repo    string
		repoErr error

		output string
	}

	type want struct {
		cmd    string
		args   []string
		err    string
		output string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			want: want{
				err: "no hook location",
			},
		},
		{
			name: "global",
			args: args{
				glob: "test",
			},
			want: want{
				cmd:    "git",
				args:   []string{"config", "--get", "core.hooksPath"},
				output: "test",
			},
		},
		{
			name: "repo",
			args: args{
				repo: "test",
			},
			want: want{
				cmd:    "git",
				args:   []string{"rev-parse", "--absolute-git-dir"},
				output: "test",
			},
		},
		{
			name: "global_error",
			args: args{
				glob:    "test",
				globErr: errMock,
			},
			want: want{
				err: "unable to check global hook: unable to run command: error",
			},
		},
		{
			name: "repo_error",
			args: args{
				repo:    "test",
				repoErr: errMock,
			},
			want: want{
				err: "unable to check repository hook: unable to run command: error",
			},
		},
		{
			name: "spaces_at_end",
			args: args{
				glob: "test    ",
			},
			want: want{
				cmd:    "git",
				args:   []string{"config", "--get", "core.hooksPath"},
				output: "test",
			},
		},
		{
			name: "spaces_at_beginning",
			args: args{
				glob: "    test",
			},
			want: want{
				cmd:    "git",
				args:   []string{"config", "--get", "core.hooksPath"},
				output: "test",
			},
		},
		{
			name: "spaces_at_beginning_end",
			args: args{
				glob: "    test    ",
			},
			want: want{
				cmd:    "git",
				args:   []string{"config", "--get", "core.hooksPath"},
				output: "test",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := MockRun{
				glob:    tt.args.glob,
				globErr: tt.args.globErr,
				repo:    tt.args.repo,
				repoErr: tt.args.repoErr,
			}

			got, err := hook.Locate(r.Run())
			if tt.want.err != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.cmd, r.cmd)
			assert.Equal(t, tt.want.args, r.args)
			assert.Equal(t, tt.want.output, got)
		})
	}
}
