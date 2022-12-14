package emoji_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/emoji"
	"github.com/stretchr/testify/assert"
)

var firstGitmojiEmoji = emoji.Emoji{
	Name:        "art",
	Character:   "🎨",
	Description: "Improve structure / format of the code.",
	Characters:  1,
	Codepoint:   "1f3a8",
	Hex:         "F0 9F 8E A8",
	Shortcode:   ":art:",
}

func TestNew(t *testing.T) {
	type want struct {
		len   int
		name  string
		emoji emoji.Emoji
	}

	tests := []struct {
		name    string
		options func(*emoji.Set)
		want    want
	}{
		{
			name: "default",
			want: want{
				len:   72,
				name:  "gitmoji",
				emoji: firstGitmojiEmoji,
			},
		},
		{
			name:    "gitmoji",
			options: emoji.WithEmojiSet(emoji.GitmojiProfile),
			want: want{
				len:   72,
				name:  "gitmoji",
				emoji: firstGitmojiEmoji,
			},
		},
		{
			name:    "devmoji",
			options: emoji.WithEmojiSet(emoji.DevmojiProfile),
			want: want{
				len:  0,
				name: "gitmoji",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := emoji.New(tt.options)

			assert.Equal(t, tt.want.len, len(e.Emojis))
			if len(e.Emojis) > 0 {
				assert.Equal(t, tt.want.emoji, e.Emojis[0])
			}
		})
	}
}

func TestFind(t *testing.T) {
	type want struct {
		valid     bool
		name      string
		character string
		shortcode string
	}

	tests := []struct {
		name  string
		input string
		want  want
	}{
		{
			name:  "emoji",
			input: "🎨",
			want: want{
				valid:     true,
				name:      "art",
				character: "🎨",
				shortcode: ":art:",
			},
		},
		{
			name:  "shortcode",
			input: ":art:",
			want: want{
				valid:     true,
				name:      "art",
				character: "🎨",
				shortcode: ":art:",
			},
		},
		{
			name:  "word",
			input: "something",
		},
		{
			name:  "empty",
			input: "",
		},
	}

	es := emoji.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := es.Find(tt.input)
			if !tt.want.valid {
				assert.False(t, got.Valid)
				return
			}
			assert.True(t, got.Valid)
			assert.Equal(t, tt.want.name, got.Emoji.Name)
			assert.Equal(t, tt.want.character, got.Emoji.Character)
			assert.Equal(t, tt.want.shortcode, got.Emoji.Shortcode)
		})
	}
}

func TestFindByCharacter(t *testing.T) {
	type want struct {
		valid bool
		name  string
	}

	tests := []struct {
		name  string
		input string
		want  want
	}{
		{
			name:  "emoji",
			input: "🎨",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:  "shortcode",
			input: ":art:",
			want: want{
				valid: false,
			},
		},
		{
			name:  "word",
			input: "something",
		},
		{
			name:  "empty",
			input: "",
		},
	}

	es := emoji.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := es.FindByCharacter(tt.input)
			if !tt.want.valid {
				assert.False(t, got.Valid)
				return
			}
			assert.True(t, got.Valid)
			assert.Equal(t, tt.want.name, got.Emoji.Name)
		})
	}
}

func TestFindByShortcode(t *testing.T) {
	type want struct {
		valid bool
		name  string
	}

	tests := []struct {
		name  string
		input string
		want  want
	}{
		{
			name:  "shortcode",
			input: ":art:",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:  "emoji",
			input: "🎨",
		},
		{
			name:  "word",
			input: "something",
		},
		{
			name:  "empty",
			input: "",
		},
	}

	es := emoji.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := es.FindByShortcode(tt.input)
			if !tt.want.valid {
				assert.False(t, got.Valid)
				return
			}
			assert.True(t, got.Valid)
			assert.Equal(t, tt.want.name, got.Emoji.Name)
		})
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "emoji", input: "🎨", want: true},
		{name: "shortcode", input: ":art:", want: true},
		{name: "empty", input: "", want: false},
		{name: "word", input: "emoji", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := emoji.Has(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
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
			got := emoji.HasCharacter(tt.input)
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
