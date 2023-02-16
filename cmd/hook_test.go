package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/cmd"
	"github.com/mikelorant/committed/internal/hook"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

type MockHook struct {
	opts hook.Options
	err  error
}

func (h *MockHook) Do(opts hook.Options) error {
	h.opts = opts

	if h.err != nil {
		return h.err
	}

	return nil
}

func TestHookCmd(t *testing.T) {
	type args struct {
		args []string
		err  error
	}

	type want struct {
		opts hook.Options
		err  string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "help",
		},
		{
			name: "install",
			args: args{
				args: []string{"--install"},
			},
			want: want{
				opts: hook.Options{
					Install: true,
				},
			},
		},
		{
			name: "uninstall",
			args: args{
				args: []string{"--uninstall"},
			},
			want: want{
				opts: hook.Options{
					Uninstall: true,
				},
			},
		},
		{
			name: "hook_invalid",
			args: args{
				args: []string{"--invalid"},
			},
		},
		{
			name: "error",
			args: args{
				args: []string{"--install"},
				err:  errMock,
			},
			want: want{
				err: "error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			mlog := NewMockLogger(&buf)

			h := MockHook{
				err: tt.args.err,
			}

			a := cmd.App{
				Hooker: &h,
				Logger: mlog,
			}

			hook := cmd.NewHookCmd(a)

			hook.SetOut(&buf)
			hook.SetErr(&buf)
			hook.SetArgs(tt.args.args)

			hook.Execute()
			if tt.want.err != "" {
				assert.Error(t, h.err)
				assert.ErrorContains(t, h.err, tt.want.err)
				output := stripString(buf.String())
				autogold.ExpectFile(t, autogold.Raw(output), autogold.Name(tt.name))
				return
			}
			assert.NoError(t, h.err)

			assert.Equal(t, tt.want.opts, h.opts)

			output := stripString(buf.String())
			autogold.ExpectFile(t, autogold.Raw(output), autogold.Name(tt.name))
		})
	}
}

func TestNewHookCmdFlags(t *testing.T) {
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
			name: "install_flag",
			args: "--install",
			want: want{
				flags: map[string]flag{
					"install": {
						shorthand:   "",
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
			name: "uninstall_flag",
			args: "--uninstall",
			want: want{
				flags: map[string]flag{
					"uninstall": {
						shorthand:   "",
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
			name: "hook_invalid",
			args: "--invalid",
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			var a cmd.App

			hcmd := cmd.NewHookCmd(a)

			hcmd.SetOut(&buf)
			hcmd.SetErr(&buf)

			args := strings.Split(tt.args, " ")
			hcmd.SetArgs(args)

			hcmd.RunE = func(cmd *cobra.Command, args []string) error {
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

			err := hcmd.Execute()

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
