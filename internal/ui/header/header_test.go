package header_test

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/ui/header"
	"github.com/mikelorant/committed/internal/ui/uitest"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	type args struct {
		env    map[string]string
		config func(c *commit.Config)
		model  func(m header.Model) header.Model
	}

	type want struct {
		model func(m header.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
		},
		{
			name: "placeholder",
			args: args{
				config: func(c *commit.Config) {
					c.Placeholders.Summary = "placeholder"
				},
			},
			want: want{
				model: func(m header.Model) {
					assert.Equal(t, "", m.Summary())
				},
			},
		},
		{
			name: "amend_summary",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Head.Hash = "1"
					c.Repository.Head.Message = "summary"
					c.Amend = true
				},
			},
			want: want{
				model: func(m header.Model) {
					assert.Equal(t, "summary", m.Summary())
				},
			},
		},
		{
			name: "amend_emoji_summary",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Head.Hash = "1"
					c.Repository.Head.Message = ":bug: summary"
					c.Amend = true
				},
			},
			want: want{
				model: func(m header.Model) {
					assert.Equal(t, ":bug:", m.Emoji.Shortcode)
					assert.Equal(t, "summary", m.Summary())
				},
			},
		},
		{
			name: "amend_emoji",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Head.Hash = "1"
					c.Repository.Head.Message = ":bug:"
					c.Amend = true
				},
			},
		},
		{
			name: "amend_summary_multiline",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Head.Hash = "1"
					c.Repository.Head.Message = "summary\n\nbody"
					c.Amend = true
				},
			},
		},
		{
			name: "focus",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					return m
				},
			},
			want: want{
				model: func(m header.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
		{
			name: "blur_emoji",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m, _ = header.ToModel(m.Update(nil))
					m.Blur()
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m header.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "blur_summary",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m.Blur()
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m header.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "expand",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "expand_emojis",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
					c.Repository.Head.Message = "summary\n\nbody"
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "expand_emojis_env",
			args: args{
				env: map[string]string{
					"COMMITTED_TERMINAL": "ttyd",
				},
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
					c.Repository.Head.Message = "summary\n\nbody"
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "filter_emoji",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, "bug"), nil)
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "filter_emoji_no_match",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, "test test test"), nil)
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "select_emoji",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":art:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_down",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					for i := 0; i < 2; i++ {
						m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					}
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":fire:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_down_up",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					for i := 0; i < 2; i++ {
						m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					}
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyUp}))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":zap:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_page_down_page_up",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					for i := 0; i < 2; i++ {
						m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					}
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgUp}))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":tada:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_page_down_last_page",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					for i := 0; i < 7; i++ {
						m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					}
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":monocle_face:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_page_down_last_page_page_up",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					for i := 0; i < 7; i++ {
						m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					}
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgUp}))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":mag:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_page_down_last_page_exceeded",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					for i := 0; i < 10; i++ {
						m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					}
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":monocle_face:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_page_down_last_page_exceeded_page_up",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					for i := 0; i < 10; i++ {
						m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					}
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgUp}))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":mag:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_filter",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, "bug"), nil)
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":bug:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_filter_clear",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, "bug"), nil)
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEscape}))
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, ":art:", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "select_emoji_delete",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDelete}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, "", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "emoji_empty_delete",
			args: args{
				config: func(c *commit.Config) {
					c.Emojis = emoji.New().Emojis
				},
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectEmoji()
					m.Expand = true
					m, _ = header.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDelete}))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, "", m.Emoji.Shortcode)
				},
			},
		},
		{
			name: "summary_text",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, "test", m.Summary())
				},
			},
		},
		{
			name: "summary_emoji",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, "🎨"), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, "🎨", m.Summary())
				},
			},
		},
		{
			name: "summary_emoji_text",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, "🎨 text"), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Equal(t, "🎨 text", m.Summary())
				},
			},
		},
		{
			name: "summary_empty",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 0)
				},
			},
		},
		{
			name: "summary_short_boundary_low",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 1)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 1)
				},
			},
		},
		{
			name: "summary_short_boundary_high",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 4)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 4)
				},
			},
		},
		{
			name: "summary_normal_boundary_low",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 5)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 5)
				},
			},
		},
		{
			name: "summary_normal_boundary_high",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 40)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 40)
				},
			},
		},
		{
			name: "summary_warning_boundary_low",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 41)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 41)
				},
			},
		},
		{
			name: "summary_warning_boundary_high",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 50)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 50)
				},
			},
		},
		{
			name: "summary_maximum_boundary_low",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 51)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 51)
				},
			},
		},
		{
			name: "summary_maximum_boundary_high",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 72)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 72)
				},
			},
		},
		{
			name: "summary_exceed",
			args: args{
				model: func(m header.Model) header.Model {
					m.Focus()
					m.SelectSummary()
					m, _ = header.ToModel(m.Update(nil))
					m, _ = header.ToModel(uitest.SendString(m, strings.Repeat("*", 100)), nil)
					return m
				},
			},
			want: want{
				func(m header.Model) {
					assert.Len(t, m.Summary(), 72)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.env {
				t.Setenv(k, v)
			}

			c := testConfig()
			if tt.args.config != nil {
				tt.args.config(&c)
			}

			m := header.New(&c)

			if tt.args.model != nil {
				m = tt.args.model(m)
			}

			if tt.want.model != nil {
				tt.want.model(m)
			}

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}

func testConfig() commit.Config {
	return commit.Config{
		Placeholders: commit.Placeholders{
			Hash: "1",
		},
	}
}
