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

var firstDevmojiEmoji = emoji.Emoji{
	Name:        "feat",
	Character:   "✨",
	Description: "fix: a bug fix",
	Characters:  1,
	Codepoint:   "2728",
	Hex:         "F0 9F 93 9D",
	Shortcode:   ":sparkles:",
}

var firstEmojiLogEmoji = emoji.Emoji{
	Name:        "new",
	Character:   "📦",
	Description: "Use when you add something entirely.",
	Characters:  1,
	Codepoint:   "1f4e6",
	Hex:         "F0 9F 93 A6",
	Shortcode:   ":package:",
}

func TestNew(t *testing.T) {
	t.Parallel()

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
				len:   19,
				name:  "devmoji",
				emoji: firstDevmojiEmoji,
			},
		},
		{
			name:    "emojilog",
			options: emoji.WithEmojiSet(emoji.EmojiLogProfile),
			want: want{
				len:   7,
				name:  "emojilog",
				emoji: firstEmojiLogEmoji,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := emoji.New(tt.options)

			assert.Equal(t, tt.want.len, len(e.Emojis))
			if len(e.Emojis) > 0 {
				assert.Equal(t, tt.want.name, e.Name)
				assert.Equal(t, tt.want.emoji, e.Emojis[0])
			}
		})
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := emoji.Has(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHasEmoji(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := emoji.HasCharacter(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHasShortcode(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := emoji.HasShortcode(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
