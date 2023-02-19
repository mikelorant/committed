package hook_test

import (
	"errors"
	"io"
	"os"
	"path"
	"testing"

	"github.com/mikelorant/committed/internal/hook"

	"github.com/stretchr/testify/assert"
)

type (
	Runner  func(io.Writer, string, []string) error
	Stater  func(string) (os.FileInfo, error)
	Opener  func(string) (*os.File, error)
	Creator func(name string, flag int, perm os.FileMode) (*os.File, error)
	Locater func(run Runner) (string, error)
)

var errMock = errors.New("error")

func MockCreate(err error) func(name string, flag int, perm os.FileMode) (*os.File, error) {
	return func(name string, flag int, perm os.FileMode) (*os.File, error) {
		if err != nil {
			return nil, err
		}

		fh, _ := os.CreateTemp("", "")

		return fh, nil
	}
}

func MockLocater(t *testing.T, emptyLoc bool, data string, err error) func(run hook.Runner) (string, error) {
	return func(run hook.Runner) (string, error) {
		if err != nil {
			return "", err
		}

		tmpDir := t.TempDir()
		file := path.Join(tmpDir, hook.GitHook)

		_ = os.MkdirAll(path.Dir(file), 0o755)

		if emptyLoc {
			return "", nil
		}

		if data == "" {
			return tmpDir, nil
		}

		_ = os.WriteFile(file, []byte(data), 0o755)

		return tmpDir, nil
	}
}

func MockOpen(err error) func(string) (*os.File, error) {
	return func(file string) (*os.File, error) {
		if err != nil {
			return nil, err
		}

		fh, _ := os.Open(file)

		return fh, nil
	}
}

func MockRun(str string, err error) func(w io.Writer, str string, ss []string) error {
	return func(w io.Writer, str string, ss []string) error {
		if err != nil {
			return err
		}

		w.Write([]byte(str))

		return nil
	}
}

func MockStat() func(string) (os.FileInfo, error) {
	return func(file string) (os.FileInfo, error) {
		st, err := os.Stat(file)

		return st, err
	}
}

func TestInstall(t *testing.T) {
	t.Parallel()

	type args struct {
		data      string
		emptyLoc  bool
		createErr error
		locateErr error
		openErr   error
		runErr    error
	}

	type want struct {
		err string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
		},
		{
			name: "managed",
			args: args{
				data: hook.Marker,
			},
		},
		{
			name: "unmanaged",
			args: args{
				data: "unmanaged",
			},
			want: want{
				err: "hook file unmanaged",
			},
		},
		{
			name: "no_location",
			args: args{
				emptyLoc: true,
			},
			want: want{
				err: "no hook location found",
			},
		},
		{
			name: "create_error",
			args: args{
				createErr: errMock,
			},
			want: want{
				err: "unable to determine managed state: unable to create file: error",
			},
		},
		{
			name: "location_error",
			args: args{
				locateErr: errMock,
			},
			want: want{
				err: "unable to determine hook location: error",
			},
		},
		{
			name: "open_error",
			args: args{
				data:    hook.Marker,
				openErr: errMock,
			},
			want: want{
				err: "unable to determine managed state: unable to open file: error",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := hook.Hook{
				Creator: MockCreate(tt.args.createErr),
				Opener:  MockOpen(tt.args.openErr),
				Locater: MockLocater(t, tt.args.emptyLoc, tt.args.data, tt.args.locateErr),
				Runner:  MockRun(tt.args.data, tt.args.runErr),
				Stater:  MockStat(),
			}

			err := h.Install()
			if tt.want.err != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
