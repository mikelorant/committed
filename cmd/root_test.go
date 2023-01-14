package cmd_test

import (
	"bytes"
	"strings"
	"testing"
	"unicode"

	"github.com/acarl005/stripansi"
	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestNewRootCmd(t *testing.T) {
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
			root := cmd.NewRootCmd()
			var buf bytes.Buffer

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
