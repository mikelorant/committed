package shell_test

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"testing"

	"github.com/mikelorant/committed/internal/shell"
	"github.com/stretchr/testify/assert"
)

type badBuffer struct {
	buf bytes.Buffer
	err error
}

func (b *badBuffer) Write(p []byte) (int, error) {
	if b.err != nil {
		return 0, b.err
	}
	return b.buf.Write(p)
}

func (b *badBuffer) Read(p []byte) (int, error) {
	return b.buf.Read(p)
}

var (
	errMockRun           = errors.New("unable to exec command")
	errMockRead          = errors.New("unable to copy commit output")
	errMockReadPathError = &fs.PathError{
		Path: "/dev/pmtx",
		Err:  errMockRead,
	}
)

func TestRun(t *testing.T) {
	type args struct {
		command string
		args    []string
	}

	type want struct {
		output  string
		runErr  error
		readErr error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid",
			args: args{
				command: "echo",
				args:    []string{"hello world"},
			},
			want: want{
				output: "hello world\r\n",
			},
		},
		{
			name: "empty",
			want: want{
				runErr: errMockRun,
			},
		},
		{
			name: "run_error",
			args: args{
				command: "no_command",
			},
			want: want{
				runErr: errMockRun,
			},
		},
		{
			name: "copy_error",
			args: args{
				command: "echo",
				args:    []string{"hello world"},
			},
			want: want{
				readErr: errMockRead,
			},
		},
		{
			name: "copy_path_error",
			args: args{
				command: "echo",
				args:    []string{"hello world"},
			},
			want: want{
				readErr: errMockReadPathError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := badBuffer{
				err: tt.want.readErr,
			}

			err := shell.Run(&w, tt.args.command, tt.args.args)
			switch {
			case tt.want.runErr != nil:
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.want.runErr.Error())
				return
			case tt.want.readErr != nil:
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.want.readErr.Error())
				return
			}
			assert.Nil(t, err)

			b, err := io.ReadAll(&w)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.output, string(b))
		})
	}
}
