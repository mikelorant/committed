package hook_test

import (
	"path"
	"testing"

	"github.com/mikelorant/committed/internal/hook"

	"github.com/stretchr/testify/assert"
)

type MockDelete struct {
	delFile string
	err     error
}

func (d *MockDelete) Delete() func(string) error {
	return func(file string) error {
		d.delFile = file

		if d.err != nil {
			return d.err
		}

		return nil
	}
}

func TestUninstall(t *testing.T) {
	type args struct {
		data     string
		emptyLoc bool
		openErr  error
		locErr   error
		delErr   error
		runErr   error
	}

	type want struct {
		delFile string
		err     string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			want: want{
				delFile: "prepare-commit-msg",
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
			name: "managed",
			args: args{
				data: hook.Marker,
			},
			want: want{
				delFile: "prepare-commit-msg",
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
			name: "locate_error",
			args: args{
				locErr: errMock,
			},
			want: want{
				err: "unable to determine hook location: error",
			},
		},
		{
			name: "delete_error",
			args: args{
				delErr: errMock,
			},
			want: want{
				err: "unable to determine managed state: unable to delete file: error",
			},
		},
		{
			name: "open_error",
			args: args{
				openErr: errMock,
				data:    hook.Marker,
			},
			want: want{
				err: "unable to determine managed state: unable to open file: error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			del := MockDelete{
				err: tt.args.delErr,
			}

			h := hook.Hook{
				Deleter: del.Delete(),
				Locater: MockLocater(t, tt.args.emptyLoc, tt.args.data, tt.args.locErr),
				Opener:  MockOpen(tt.args.openErr),
				Runner:  MockRun(tt.args.data, tt.args.runErr),
				Stater:  MockStat(),
			}

			err := h.Uninstall()
			if tt.want.err != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want.delFile, path.Base(del.delFile))
		})
	}
}
