package emoji_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/emoji"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	emojiCount := 72

	firstEmoji := emoji.Emoji{
		Name:        "art",
		Character:   "🎨",
		Description: "Improve structure / format of the code.",
		Characters:  1,
		Codepoint:   "1f3a8",
		Hex:         "F0 9F 8E A8",
		ShortCode:   ":art:",
	}

	e, err := emoji.New()
	assert.NoError(t, err)
	assert.Equal(t, emojiCount, len(e))
	assert.Equal(t, firstEmoji, e[0])
}

func TestHasEmoji(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "emoji_standard", input: "🎨", want: true},
		{name: "emoji_variant", input: "⚡️", want: true},
		{name: "emoji_wide", input: "⬇️", want: true},
		{name: "emoji_zwj", input: "🧑‍💻", want: true},
		{name: "emoji_multiple", input: "🎨🔥🐛", want: true},
		{name: "shortcode", input: ":art:", want: false},
		{name: "empty", input: "", want: false},
		{name: "ascii_symbol", input: "@", want: false},
		{name: "ascii_word", input: "emoji", want: false},
		{name: "ascii_shape", input: "●", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := emoji.HasEmoji(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHasShortcode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "shortcode_standard", input: ":art:", want: true},
		{name: "shortcode_multiple", input: ":art::bug:", want: false},
		{name: "shortcode_delimiter_only", input: ":::::", want: false},
		{name: "shortcode_short", input: ":a:", want: true},
		{name: "shortcode_empty", input: "::", want: false},
		{name: "shortcode_spaces", input: ":art: text", want: false},
		{name: "emoji", input: "🎨", want: false},
		{name: "empty", input: "", want: false},
		{name: "ascii_word", input: "emoji", want: false},
		{name: "ascii_symbol", input: "@", want: false},
		{name: "ascii_shape", input: "●", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := emoji.HasShortcode(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
