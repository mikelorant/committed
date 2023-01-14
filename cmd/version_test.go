package cmd_test

import (
	"bytes"
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/cmd"
)

func TestNewVersionCmd(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "version_arg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver := cmd.NewVersionCmd()
			var buf bytes.Buffer

			ver.SetOut(&buf)
			ver.SetErr(&buf)
			ver.SetArgs([]string{})

			ver.Execute()
			autogold.ExpectFile(t, autogold.Raw(buf.String()), autogold.Name(tt.name))
		})
	}
}