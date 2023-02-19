package cmd_test

import (
	"bytes"
	"testing"

	"github.com/mikelorant/committed/cmd"

	"github.com/hexops/autogold/v2"
)

func TestListCmd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "list_arg",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer

			list := cmd.NewListCmd(&buf)
			list.SetOut(&buf)
			list.SetErr(&buf)
			list.SetArgs([]string{})

			list.Execute()
			autogold.ExpectFile(t, autogold.Raw(buf.String()), autogold.Name(tt.name))
		})
	}
}

func TestListEmojiCmd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "list_emoji_arg",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer

			list := cmd.NewListEmojiProfilesCmd(&buf)
			list.SetOut(&buf)
			list.SetErr(&buf)
			list.SetArgs([]string{})

			list.Execute()
			autogold.ExpectFile(t, autogold.Raw(buf.String()), autogold.Name(tt.name))
		})
	}
}

func TestListThemeCmd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "list_theme_arg",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer

			list := cmd.NewListThemesCmd(&buf)
			list.SetOut(&buf)
			list.SetErr(&buf)
			list.SetArgs([]string{})

			list.Execute()
			autogold.ExpectFile(t, autogold.Raw(buf.String()), autogold.Name(tt.name))
		})
	}
}
