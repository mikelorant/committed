package ui_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/theme/themetest"
	"github.com/mikelorant/committed/internal/ui"
	"github.com/mikelorant/committed/internal/ui/option/setting"

	"github.com/stretchr/testify/assert"
)

func TestToConfig(t *testing.T) {
	t.Parallel()

	type args struct {
		cfg      config.Config
		paneSets func(map[string][]setting.Paner)
		theme    theme.Theme
	}

	type want struct {
		cfg func(*config.Config)
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
			name: "focus_unset",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][0] = &setting.Radio{Title: "Focus", Index: toInt(config.FocusUnset)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Focus = config.FocusUnset },
			},
		},
		{
			name: "focus_author",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][0] = &setting.Radio{Title: "Focus", Index: toInt(config.FocusAuthor)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Focus = config.FocusAuthor },
			},
		},
		{
			name: "focus_emoji",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][0] = &setting.Radio{Title: "Focus", Index: toInt(config.FocusEmoji)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Focus = config.FocusEmoji },
			},
		},
		{
			name: "focus_summary",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][0] = &setting.Radio{Title: "Focus", Index: toInt(config.FocusSummary)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Focus = config.FocusSummary },
			},
		},
		{
			name: "emoji_selector_unset",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][1] = &setting.Radio{Title: "Emoji Selector", Index: toInt(config.EmojiSelectorUnset)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSelector = config.EmojiSelectorUnset },
			},
		},
		{
			name: "emoji_selector_below",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][1] = &setting.Radio{Title: "Emoji Selector", Index: toInt(config.EmojiSelectorBelow)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSelector = config.EmojiSelectorBelow },
			},
		},
		{
			name: "emoji_selector_above",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][1] = &setting.Radio{Title: "Emoji Selector", Index: toInt(config.EmojiSelectorAbove)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSelector = config.EmojiSelectorAbove },
			},
		},
		{
			name: "emoji_set_unset",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][2] = &setting.Radio{Title: "Emoji Set", Index: toInt(config.EmojiSetUnset)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSet = config.EmojiSetUnset },
			},
		},
		{
			name: "emoji_set_committed",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][2] = &setting.Radio{Title: "Emoji Set", Index: toInt(config.EmojiSetCommitted)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSet = config.EmojiSetCommitted },
			},
		},
		{
			name: "emoji_set_gitmoji",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][2] = &setting.Radio{Title: "Emoji Set", Index: toInt(config.EmojiSetGitmoji)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSet = config.EmojiSetGitmoji },
			},
		},
		{
			name: "emoji_set_devmoji",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][2] = &setting.Radio{Title: "Emoji Set", Index: toInt(config.EmojiSetDevmoji)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSet = config.EmojiSetDevmoji },
			},
		},
		{
			name: "emoji_set_emojilog",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][2] = &setting.Radio{Title: "Emoji Set", Index: toInt(config.EmojiSetEmojiLog)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.EmojiSet = config.EmojiSetEmojiLog },
			},
		},
		{
			name: "ignore_global_author",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["General"][3] = &setting.Toggle{Title: "Emoji Set", Enable: true}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.IgnoreGlobalAuthor = true },
			},
		},
		{
			name: "colour_unset",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][0] = &setting.Radio{Title: "Colour", Index: toInt(config.ColourUnset)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Colour = config.ColourUnset },
			},
		},
		{
			name: "colour_adaptive",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][0] = &setting.Radio{Title: "Colour", Index: toInt(config.ColourAdaptive)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Colour = config.ColourAdaptive },
			},
		},
		{
			name: "colour_dark",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][0] = &setting.Radio{Title: "Colour", Index: toInt(config.ColourDark)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Colour = config.ColourDark },
			},
		},
		{
			name: "colour_light",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][0] = &setting.Radio{Title: "Colour", Index: toInt(config.ColourLight)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Colour = config.ColourLight },
			},
		},
		{
			name: "compatibility_unset",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][1] = &setting.Radio{Title: "Compatibility", Index: toInt(config.CompatibilityUnset)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Compatibility = config.CompatibilityUnset },
			},
		},
		{
			name: "compatibility_default",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][1] = &setting.Radio{Title: "Compatibility", Index: toInt(config.CompatibilityDefault)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Compatibility = config.CompatibilityDefault },
			},
		},
		{
			name: "compatibility_ttyd",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][1] = &setting.Radio{Title: "Compatibility", Index: toInt(config.CompatibilityTtyd)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Compatibility = config.CompatibilityTtyd },
			},
		},
		{
			name: "compatibility_kitty",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][1] = &setting.Radio{Title: "Compatibility", Index: toInt(config.CompatibilityKitty)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Compatibility = config.CompatibilityKitty },
			},
		},
		{
			name: "highlight_active",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Visual"][2] = &setting.Toggle{Title: "Highlight Active", Enable: true}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.HighlightActive = true },
			},
		},
		{
			name: "emoji_type_unset",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Commit"][0] = &setting.Radio{Title: "Emoji Type", Index: toInt(config.EmojiTypeUnset)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.Commit.EmojiType = config.EmojiTypeUnset },
			},
		},
		{
			name: "emoji_type_shortcode",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Commit"][0] = &setting.Radio{Title: "Emoji Type", Index: toInt(config.EmojiTypeShortcode)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.Commit.EmojiType = config.EmojiTypeShortcode },
			},
		},
		{
			name: "emoji_type_character",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Commit"][0] = &setting.Radio{Title: "Emoji Type", Index: toInt(config.EmojiTypeCharacter)}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.Commit.EmojiType = config.EmojiTypeCharacter },
			},
		},
		{
			name: "signoff",
			args: args{
				paneSets: func(ps map[string][]setting.Paner) {
					ps["Commit"][1] = &setting.Toggle{Title: "Sign-off", Enable: true}
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.Commit.Signoff = true },
			},
		},
		{
			name: "theme",
			args: args{
				theme: theme.New(theme.Tint{
					Default:  themetest.NewStubTints(5)[0],
					Defaults: themetest.NewStubTints(5),
				}),
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.View.Theme = "id0" },
			},
		},
		{
			name: "authors",
			args: args{
				cfg: config.Config{
					Authors: testAuthors(),
				},
			},
			want: want{
				cfg: func(cfg *config.Config) { cfg.Authors = testAuthors() },
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ps := testPaneSets()
			if tt.args.paneSets != nil {
				tt.args.paneSets(ps)
			}

			if tt.args.cfg.Authors == nil {
				tt.args.cfg.Authors = testAuthors()
			}

			want := wantConfig()
			if tt.want.cfg != nil {
				tt.want.cfg(&want)
			}

			if want.Authors == nil {
				want.Authors = testAuthors()
			}

			cfg := ui.ToConfig(tt.args.cfg, ps, tt.args.theme)

			assert.Equal(t, want, cfg)
		})
	}
}

func testPaneSets() map[string][]setting.Paner {
	return map[string][]setting.Paner{
		"General": {
			&setting.Radio{Title: "Focus"},
			&setting.Radio{Title: "EmojiSelector"},
			&setting.Radio{Title: "EmojiSet"},
			&setting.Toggle{Title: "IgnoreGlobalAuthor"},
		},
		"Visual": {
			&setting.Radio{Title: "Colour"},
			&setting.Radio{Title: "Compatibility"},
			&setting.Toggle{Title: "HighlightActive"},
		},
		"Commit": {
			&setting.Radio{Title: "EmojiType"},
			&setting.Toggle{Title: "Signoff"},
		},
	}
}

func testAuthors() []repository.User {
	return []repository.User{
		{Name: "John Doe", Email: "john.doe@example.com"},
	}
}

func wantConfig() config.Config {
	return config.Config{
		View: config.View{
			Colour:        config.ColourAdaptive,
			Compatibility: config.CompatibilityDefault,
			EmojiSelector: config.EmojiSelectorBelow,
			EmojiSet:      config.EmojiSetCommitted,
			Focus:         config.FocusAuthor,
		},
		Commit: config.Commit{
			EmojiType: config.EmojiTypeShortcode,
		},
	}
}

func toInt[T ~int](n T) int {
	return int(n) - 1
}
