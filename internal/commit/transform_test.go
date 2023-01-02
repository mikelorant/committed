package commit_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestMessageToEmoji(t *testing.T) {
	type want struct {
		valid bool
		name  string
	}

	tests := []struct {
		name    string
		message string
		want    want
	}{
		{
			name:    "emoji_character_summary",
			message: "🎨 summary",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "emoji_shortcode_summary",
			message: ":art: summary",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "emoji_only",
			message: "🎨",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "emoji_multiline_message",
			message: "🎨 summary\n\nbody\n",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "unknown_emoji",
			message: "😀 summary",
		},
		{
			name:    "summary",
			message: "summary",
		},
		{
			name:    "empty",
			message: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := commit.Config{
				Repository: repository.Description{
					Head: repository.Head{
						Message: tt.message,
					},
				},
			}

			e := cfg.MessageToEmoji()
			if !tt.want.valid {
				assert.False(t, tt.want.valid)
				assert.Empty(t, e.Emoji.Name)
				return
			}
			assert.True(t, tt.want.valid)
			assert.Equal(t, tt.want.name, e.Emoji.Name)
		})
	}
}

func TestMessageToSummary(t *testing.T) {
	tests := []struct {
		name    string
		message string
		summary string
	}{
		{
			name:    "summary",
			message: "summary",
			summary: "summary",
		},
		{
			name:    "emoji_summary",
			message: "😀 summary",
			summary: "summary",
		},
		{
			name:    "emoji_only",
			message: "😀",
		},
		{
			name:    "short_summary",
			message: "a",
			summary: "a",
		},
		{
			name:    "long_summary",
			message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
		},
		{
			name:    "multiline",
			message: "summary\n\nbody",
			summary: "summary",
		},
		{
			name:    "multiline_newline",
			message: "summary\n",
			summary: "summary",
		},
		{
			name:    "multiline_no_newline",
			message: "summary\nbody\n",
		},
		{
			name: "empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := commit.Config{
				Repository: repository.Description{
					Head: repository.Head{
						Message: tt.message,
					},
				},
			}

			s := cfg.MessageToSummary()
			if tt.summary == "" {
				assert.Empty(t, s)
				return
			}
			assert.NotEmpty(t, s)
			assert.Equal(t, tt.summary, s)
		})
	}
}

func TestMessageToBody(t *testing.T) {
	tests := []struct {
		name    string
		message string
		body    string
	}{
		{
			name:    "summary_body",
			message: "summary\n\nbody",
			body:    "body",
		},
		{
			name:    "summary",
			message: "summary",
		},
		{
			name:    "multine_newline",
			message: "summary\n",
			body:    "",
		},
		{
			name:    "multine_no_newline",
			message: "summary\nbody",
			body:    "summary\nbody",
		},
		{
			name:    "empty",
			message: "",
		},
		{
			name:    "long_summary",
			message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
			body:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := commit.Config{
				Repository: repository.Description{
					Head: repository.Head{
						Message: tt.message,
					},
				},
			}

			b := cfg.MessageToBody()
			if tt.body == "" {
				assert.Empty(t, b)
				return
			}
			assert.NotEmpty(t, b)
			assert.Equal(t, tt.body, b)
		})
	}
}
