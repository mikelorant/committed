package cmd_test

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
	"testing"
	"unicode"

	"github.com/acarl005/stripansi"
	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/cmd"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

type MockCommit struct {
	configErr error
	applyErr  error
}

type MockUI struct {
	err error
}

type MockLogger struct {
	logger *log.Logger
	rw     io.ReadWriter
}

func (m MockLogger) Fatalf(format string, v ...any) {
	m.logger.Printf(format, v...)
}

func (m MockLogger) String() string {
	out, _ := io.ReadAll(m.rw)

	return strings.TrimSpace(string(out))
}

func (m *MockCommit) Configure(opts commit.Options) (*commit.Config, error) {
	return nil, m.configErr
}

func (m *MockCommit) Apply(req *commit.Request) error {
	return m.applyErr
}

func (m *MockUI) Configure(cfg *commit.Config) {}

func (m *MockUI) Start() (*commit.Request, error) {
	return nil, m.err
}

var errMock = errors.New("error")

func NewMockLogger(rw io.ReadWriter) MockLogger {
	return MockLogger{
		rw:     rw,
		logger: log.New(rw, "", 0),
	}
}

func TestNewRootCmd(t *testing.T) {
	type args struct {
		configErr error
		applyErr  error
		startErr  error
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
			name: "config_error",
			args: args{
				configErr: errMock,
			},
			want: want{
				err: "unable to init commit: error",
			},
		},
		{
			name: "repository_error",
			args: args{
				configErr: repository.NotFoundError(),
			},
			want: want{
				err: "No git repository found.",
			},
		},
		{
			name: "start_error",
			args: args{
				startErr: errMock,
			},
			want: want{
				err: "unable to start ui: error",
			},
		},
		{
			name: "apply_error",
			args: args{
				applyErr: errMock,
			},
			want: want{
				err: "unable to apply commit: error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			mlog := NewMockLogger(&buf)

			root := cmd.NewRootCmd(cmd.App{
				Commiter: &MockCommit{
					configErr: tt.args.configErr,
					applyErr:  tt.args.applyErr,
				},
				UIer: &MockUI{
					err: tt.args.startErr,
				},
				Logger: mlog,
			})

			root.SetOut(io.Discard)
			root.SetErr(io.Discard)

			err := root.Execute()
			if tt.want.err != "" {
				assert.NotNil(t, err)
				assert.Equal(t, tt.want.err, mlog.String())
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestNewRootCmdFlags(t *testing.T) {
	type flag struct {
		shorthand   string
		value       string
		defValue    string
		changed     bool
		noOptDefVal string
	}

	type want struct {
		flags  map[string]flag
		output string
		err    bool
	}

	tests := []struct {
		name string
		args string
		want want
	}{
		{
			name: "default",
			want: want{
				err: false,
			},
		},
		{
			name: "help_arg",
			args: "help",
			want: want{
				err: false,
			},
		},
		{
			name: "help_flag",
			args: "--help",
			want: want{
				err: false,
			},
		},
		{
			name: "version_flag",
			args: "--version",
			want: want{
				err: false,
			},
		},
		{
			name: "amend_flag",
			args: "--amend",
			want: want{
				flags: map[string]flag{
					"amend": {
						shorthand:   "a",
						value:       "true",
						defValue:    "false",
						changed:     true,
						noOptDefVal: "true",
					},
				},
				err: false,
			},
		},
		{
			name: "dry-run_flag",
			args: "--dry-run",
			want: want{
				flags: map[string]flag{
					"dry-run": {
						shorthand:   "",
						value:       "true",
						defValue:    "true",
						changed:     true,
						noOptDefVal: "true",
					},
				},
				err: false,
			},
		},
		{
			name: "invalid",
			args: "--invalid",
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			root := cmd.NewRootCmd(cmd.App{
				Commiter: &MockCommit{},
				UIer:     &MockUI{},
			})

			root.SetOut(&buf)
			root.SetErr(&buf)

			args := strings.Split(tt.args, " ")
			root.SetArgs(args)

			root.RunE = func(cmd *cobra.Command, args []string) error {
				cmd.Flags().Visit(func(f *pflag.Flag) {
					if _, ok := tt.want.flags[f.Name]; !ok {
						t.FailNow()
					}
					wf := tt.want.flags[f.Name]

					assert.Equal(t, wf.shorthand, f.Shorthand)
					assert.Equal(t, wf.value, f.Value.String())
					assert.Equal(t, wf.changed, f.Changed)
					assert.Equal(t, wf.defValue, f.DefValue)
					assert.Equal(t, wf.noOptDefVal, f.NoOptDefVal)
				})

				return nil
			}

			err := root.Execute()

			if tt.want.err == true {
				assert.NotNil(t, err)
				output := stripString(buf.String())
				autogold.ExpectFile(t, autogold.Raw(output), autogold.Name(tt.name))
				return
			}
			assert.Nil(t, err)

			output := stripString(buf.String())
			autogold.ExpectFile(t, autogold.Raw(output), autogold.Name(tt.name))
		})
	}
}

func stripString(str string) string {
	s := stripansi.Strip(str)
	ss := strings.Split(s, "\n")

	var lines []string
	for _, l := range ss {
		trim := strings.TrimRightFunc(l, unicode.IsSpace)
		lines = append(lines, trim)
	}

	return strings.Join(lines, "\n")
}
