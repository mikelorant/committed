package config_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestUnmarshallYAMLFocus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  config.Focus
	}{
		{name: "empty", input: "", want: config.FocusUnset},
		{name: "author", input: "author", want: config.FocusAuthor},
		{name: "emoji", input: "emoji", want: config.FocusEmoji},
		{name: "summary", input: "summary", want: config.FocusSummary},
		{name: "invalid", input: "invalid", want: config.FocusUnset},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config.Focus

			yaml.Unmarshal([]byte(tt.input), &got)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func TestUnmarshallYAMLCompatibility(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  config.Compatibility
	}{
		{name: "empty", input: "", want: config.CompatibilityUnset},
		{name: "default", input: "default", want: config.CompatibilityDefault},
		{name: "ttyd", input: "ttyd", want: config.CompatibilityTtyd},
		{name: "kitty", input: "kitty", want: config.CompatibilityKitty},
		{name: "invalid", input: "invalid", want: config.CompatibilityUnset},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config.Compatibility

			yaml.Unmarshal([]byte(tt.input), &got)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func TestUnmarshallYAMLColour(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  config.Colour
	}{
		{name: "empty", input: "", want: config.ColourUnset},
		{name: "adaptive", input: "adaptive", want: config.ColourAdaptive},
		{name: "dark", input: "dark", want: config.ColourDark},
		{name: "light", input: "light", want: config.ColourLight},
		{name: "invalid", input: "invalid", want: config.ColourUnset},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config.Colour

			yaml.Unmarshal([]byte(tt.input), &got)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func TestMarshallYAMLFocus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.Focus
		want  string
	}{
		{name: "empty", input: config.FocusUnset, want: "\"\"\n"},
		{name: "author", input: config.FocusAuthor, want: "author\n"},
		{name: "emoji", input: config.FocusEmoji, want: "emoji\n"},
		{name: "summary", input: config.FocusSummary, want: "summary\n"},
		{name: "invalid", input: config.FocusUnset, want: "\"\"\n"},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, _ := yaml.Marshal(&tt.input)
			assert.Equal(t, tt.want, string(got), tt.name)
		})
	}
}

func TestMarshallYAMLCompatibility(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.Compatibility
		want  string
	}{
		{name: "empty", input: config.CompatibilityUnset, want: "\"\"\n"},
		{name: "default", input: config.CompatibilityDefault, want: "default\n"},
		{name: "ttyd", input: config.CompatibilityTtyd, want: "ttyd\n"},
		{name: "kitty", input: config.CompatibilityKitty, want: "kitty\n"},
		{name: "invalid", input: config.CompatibilityUnset, want: "\"\"\n"},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, _ := yaml.Marshal(&tt.input)
			assert.Equal(t, tt.want, string(got), tt.name)
		})
	}
}

func TestMarshallYAMLColour(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.Colour
		want  string
	}{
		{name: "empty", input: config.ColourUnset, want: "\"\"\n"},
		{name: "adaptive", input: config.ColourAdaptive, want: "adaptive\n"},
		{name: "dark", input: config.ColourDark, want: "dark\n"},
		{name: "light", input: config.ColourLight, want: "light\n"},
		{name: "invalid", input: config.ColourUnset, want: "\"\"\n"},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, _ := yaml.Marshal(&tt.input)
			assert.Equal(t, tt.want, string(got), tt.name)
		})
	}
}

func TestIndexColour(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.Colour
		want  int
	}{
		{name: "unset", input: config.Colour(config.ColourUnset), want: 1},
		{name: "adaptive", input: config.Colour(config.ColourAdaptive), want: 1},
		{name: "dark", input: config.Colour(config.ColourDark), want: 2},
		{name: "light", input: config.Colour(config.ColourLight), want: 3},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.input.Index())
		})
	}
}

func TestIndexCompatibility(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.Compatibility
		want  int
	}{
		{name: "unset", input: config.Compatibility(config.CompatibilityUnset), want: 1},
		{name: "default", input: config.Compatibility(config.CompatibilityDefault), want: 1},
		{name: "ttyd", input: config.Compatibility(config.CompatibilityTtyd), want: 2},
		{name: "kitty", input: config.Compatibility(config.CompatibilityKitty), want: 3},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.input.Index())
		})
	}
}

func TestIndexFocus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.Focus
		want  int
	}{
		{name: "unset", input: config.Focus(config.FocusUnset), want: 2},
		{name: "author", input: config.Focus(config.FocusAuthor), want: 1},
		{name: "emoji", input: config.Focus(config.FocusEmoji), want: 2},
		{name: "summary", input: config.Focus(config.FocusSummary), want: 3},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.input.Index())
		})
	}
}
