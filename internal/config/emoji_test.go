package config_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestUnmarshallYAMLEmojiSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  config.EmojiSet
	}{
		{name: "empty", input: "", want: config.EmojiSetUnset},
		{name: "gitmoji", input: "gitmoji", want: config.EmojiSetGitmoji},
		{name: "devmoji", input: "devmoji", want: config.EmojiSetDevmoji},
		{name: "emojilog", input: "emojilog", want: config.EmojiSetEmojiLog},
		{name: "invalid", input: "invalid", want: config.EmojiSetUnset},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config.EmojiSet

			yaml.Unmarshal([]byte(tt.input), &got)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func TestUnmarshallYAMLEmojiSelector(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  config.EmojiSelector
	}{
		{name: "empty", input: "", want: config.EmojiSelectorUnset},
		{name: "below", input: "below", want: config.EmojiSelectorBelow},
		{name: "above", input: "above", want: config.EmojiSelectorAbove},
		{name: "invalid", input: "invalid", want: config.EmojiSelectorUnset},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config.EmojiSelector

			yaml.Unmarshal([]byte(tt.input), &got)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func TestUnmarshallYAMLEmojiType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  config.EmojiType
	}{
		{name: "empty", input: "", want: config.EmojiTypeUnset},
		{name: "shortcode", input: "shortcode", want: config.EmojiTypeShortcode},
		{name: "character", input: "character", want: config.EmojiTypeCharacter},
		{name: "invalid", input: "invalid", want: config.EmojiTypeUnset},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config.EmojiType

			yaml.Unmarshal([]byte(tt.input), &got)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func TestMarshallYAMLEmojiSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.EmojiSet
		want  string
	}{
		{name: "empty", input: config.EmojiSetUnset, want: "\"\"\n"},
		{name: "gitmoji", input: config.EmojiSetGitmoji, want: "gitmoji\n"},
		{name: "devmoji", input: config.EmojiSetDevmoji, want: "devmoji\n"},
		{name: "emojilog", input: config.EmojiSetEmojiLog, want: "emojilog\n"},
		{name: "invalid", input: config.EmojiSetUnset, want: "\"\"\n"},
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

func TestMarshallYAMLEmojiSelector(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.EmojiSelector
		want  string
	}{
		{name: "empty", input: config.EmojiSelectorUnset, want: "\"\"\n"},
		{name: "below", input: config.EmojiSelectorBelow, want: "below\n"},
		{name: "abovei", input: config.EmojiSelectorAbove, want: "above\n"},
		{name: "invalid", input: config.EmojiSelectorUnset, want: "\"\"\n"},
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

func TestMarshallYAMLEmojiType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.EmojiType
		want  string
	}{
		{name: "empty", input: config.EmojiTypeUnset, want: "\"\"\n"},
		{name: "shortcode", input: config.EmojiTypeShortcode, want: "shortcode\n"},
		{name: "character", input: config.EmojiTypeCharacter, want: "character\n"},
		{name: "invalid", input: config.EmojiTypeUnset, want: "\"\"\n"},
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

func TestIndexEmojiSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.EmojiSet
		want  int
	}{
		{name: "unset", input: config.EmojiSet(config.EmojiSetUnset), want: 2},
		{name: "committed", input: config.EmojiSet(config.EmojiSetCommitted), want: 1},
		{name: "gitmoji", input: config.EmojiSet(config.EmojiSetGitmoji), want: 2},
		{name: "devmoji", input: config.EmojiSet(config.EmojiSetDevmoji), want: 3},
		{name: "emojilog", input: config.EmojiSet(config.EmojiSetEmojiLog), want: 4},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.input.Index())
		})
	}
}

func TestIndexEmojiType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.EmojiType
		want  int
	}{
		{name: "unset", input: config.EmojiType(config.EmojiTypeUnset), want: 1},
		{name: "shortcode", input: config.EmojiType(config.EmojiTypeShortcode), want: 1},
		{name: "character", input: config.EmojiType(config.EmojiTypeCharacter), want: 2},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.input.Index())
		})
	}
}

func TestIndexEmojiSelector(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input config.EmojiSelector
		want  int
	}{
		{name: "unset", input: config.EmojiSelector(config.EmojiSelectorUnset), want: 1},
		{name: "below", input: config.EmojiSelector(config.EmojiSelectorBelow), want: 1},
		{name: "above", input: config.EmojiSelector(config.EmojiSelectorAbove), want: 2},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.input.Index())
		})
	}
}
