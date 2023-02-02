package commit_test

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/stretchr/testify/assert"
)

func TestFileOpen(t *testing.T) {
	type args struct {
		create bool
		env    bool
	}

	type want struct {
		err        error
		returnType io.Reader
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			args: args{
				create: true,
			},
			want: want{
				returnType: &os.File{},
			},
		},
		{
			name: "default_env",
			args: args{
				create: true,
				env:    true,
			},
			want: want{
				returnType: &os.File{},
			},
		},
		{
			name: "path_error",
			want: want{
				returnType: &strings.Reader{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				file string
				dir  string
			)

			const env = "TEST_ENV"

			tmpDir := t.TempDir()

			if tt.args.env {
				t.Setenv(env, tmpDir)
			}

			switch {
			case tt.args.create && tt.args.env:
				dir = fmt.Sprintf("$%v", env)
			case tt.args.create:
				dir = tmpDir
			}

			file = path.Join(dir, tt.name)

			if tt.args.create {
				f := path.Join(tmpDir, tt.name)
				if _, err := os.Create(os.ExpandEnv(f)); err != nil {
					t.Fail()
					return
				}
			}

			r, err := commit.FileOpen()(file)
			if tt.want.err != nil {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.IsType(t, tt.want.returnType, r)
		})
	}
}

func TestFileCreate(t *testing.T) {
	type args struct {
		truncate bool
		env      bool
		err      bool
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
			name: "create",
		},
		{
			name: "create_env",
			args: args{
				env: true,
			},
		},
		{
			name: "create_truncate",
			args: args{
				truncate: true,
			},
		},
		{
			name: "create_error",
			args: args{
				err: true,
			},
			want: want{
				err: "unable to create file",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				file string
				dir  string
			)

			const env = "TEST_ENV"

			tmpDir := t.TempDir()

			dir = tmpDir
			if tt.args.env {
				t.Setenv(env, tmpDir)
				dir = fmt.Sprintf("$%v", env)
			}

			file = path.Join(dir, tt.name)

			if tt.args.truncate || tt.args.err {
				w, err := os.Create(os.ExpandEnv(file))
				if err != nil {
					t.Fail()
				}

				if tt.args.err {
					os.Chmod(file, 0o000)
				}

				if _, err := io.WriteString(w, tt.name); err != nil {
					t.Fail()
				}
			}

			_, err := commit.FileCreate()(file)
			if tt.want.err != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}
			assert.NoError(t, err)

			data, _ := os.ReadFile(file)
			assert.Empty(t, string(data))
		})
	}
}

func TestFileExists(t *testing.T) {
	type args struct {
		env    bool
		create bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "missing",
		},
		{
			name: "exists",
			args: args{
				create: true,
			},
			want: true,
		},
		{
			name: "exists_env",
			args: args{
				create: true,
				env:    true,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				file string
				dir  string
			)

			const env = "TEST_ENV"

			tmpDir := t.TempDir()

			if tt.args.env {
				t.Setenv(env, tmpDir)
			}

			switch {
			case tt.args.create && tt.args.env:
				dir = fmt.Sprintf("$%v", env)
			case tt.args.create:
				dir = tmpDir
			}

			file = path.Join(dir, tt.name)

			if tt.args.create {
				if _, err := os.Create(os.ExpandEnv(file)); err != nil {
					t.Fail()
				}
			}

			got := commit.FileExists(file)
			assert.Equal(t, tt.want, got)
		})
	}
}
