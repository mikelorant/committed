package config_test

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		data   string
		config config.Config
		err    error
	}{
		{
			name:   "empty",
			config: config.Config{},
		},
		{
			name: "invalid",
			data: "invalid",
			err:  new(yaml.TypeError),
		},
		{
			name:   "focus_empty",
			data:   "view: {focus:}",
			config: config.Config{View: config.View{Focus: config.FocusUnset}},
		},
		{
			name:   "focus_author",
			data:   "view: {focus: author}",
			config: config.Config{View: config.View{Focus: config.FocusAuthor}},
		},
		{
			name:   "focus_emoji",
			data:   "view: {focus: emoji}",
			config: config.Config{View: config.View{Focus: config.FocusEmoji}},
		},
		{
			name:   "focus_summary",
			data:   "view: {focus: summary}",
			config: config.Config{View: config.View{Focus: config.FocusSummary}},
		},
		{
			name:   "focus_invalid",
			data:   "view: {focus: invalid}",
			config: config.Config{View: config.View{Focus: config.FocusUnset}},
		},
		{
			name:   "compatibility_empty",
			data:   "view: {compatibility:}",
			config: config.Config{View: config.View{Compatibility: config.CompatibilityUnset}},
		},
		{
			name:   "compatibility_default",
			data:   "view: {compatibility: default}",
			config: config.Config{View: config.View{Compatibility: config.CompatibilityDefault}},
		},
		{
			name:   "compatibility_ttyd",
			data:   "view: {compatibility: ttyd}",
			config: config.Config{View: config.View{Compatibility: config.CompatibilityTtyd}},
		},
		{
			name:   "compatibility_kitty",
			data:   "view: {compatibility: kitty}",
			config: config.Config{View: config.View{Compatibility: config.CompatibilityKitty}},
		},
		{
			name:   "compatibility_invalid",
			data:   "view: {compatibility: invalid}",
			config: config.Config{View: config.View{Compatibility: config.CompatibilityUnset}},
		},
		{
			name:   "colour_unset",
			data:   "view: {colour:}",
			config: config.Config{View: config.View{Colour: config.ColourUnset}},
		},
		{
			name:   "colour_adaptive",
			data:   "view: {colour: adaptive}",
			config: config.Config{View: config.View{Colour: config.ColourAdaptive}},
		},
		{
			name:   "colour_dark",
			data:   "view: {colour: dark}",
			config: config.Config{View: config.View{Colour: config.ColourDark}},
		},
		{
			name:   "colour_light",
			data:   "view: {colour: light}",
			config: config.Config{View: config.View{Colour: config.ColourLight}},
		},
		{
			name:   "colour_invalid",
			data:   "view: {colour: invalid}",
			config: config.Config{View: config.View{Colour: config.ColourUnset}},
		},
		{
			name:   "emojiset_empty",
			data:   "view: {emojiSet:}",
			config: config.Config{View: config.View{EmojiSet: config.EmojiSetUnset}},
		},
		{
			name:   "emojiset_empty",
			data:   "view: {emojiSet:}",
			config: config.Config{View: config.View{EmojiSet: config.EmojiSetUnset}},
		},
		{
			name:   "emojiset_gitmoji",
			data:   "view: {emojiSet: gitmoji}",
			config: config.Config{View: config.View{EmojiSet: config.EmojiSetGitmoji}},
		},
		{
			name:   "emojiset_devmoji",
			data:   "view: {emojiSet: devmoji}",
			config: config.Config{View: config.View{EmojiSet: config.EmojiSetDevmoji}},
		},
		{
			name:   "emojiset_emojilog",
			data:   "view: {emojiSet: emojilog}",
			config: config.Config{View: config.View{EmojiSet: config.EmojiSetEmojiLog}},
		},
		{
			name:   "emojiset_invalid",
			data:   "view: {emojiSet: invalid}",
			config: config.Config{View: config.View{EmojiSet: config.EmojiSetUnset}},
		},
		{
			name:   "emojiselector_empty",
			data:   "view: {emojiSelector:}",
			config: config.Config{View: config.View{EmojiSelector: config.EmojiSelectorUnset}},
		},
		{
			name:   "emojiselector_below",
			data:   "view: {emojiSelector: below}",
			config: config.Config{View: config.View{EmojiSelector: config.EmojiSelectorBelow}},
		},
		{
			name:   "emojiselector_above",
			data:   "view: {emojiSelector: above}",
			config: config.Config{View: config.View{EmojiSelector: config.EmojiSelectorAbove}},
		},
		{
			name:   "emojiselector_invalid",
			data:   "view: {emojiSelector: invalid}",
			config: config.Config{View: config.View{EmojiSelector: config.EmojiSelectorUnset}},
		},
		{
			name:   "emojitype_empty",
			data:   "commit: {emojiType:}",
			config: config.Config{Commit: config.Commit{EmojiType: config.EmojiTypeUnset}},
		},
		{
			name:   "emojitype_shortcode",
			data:   "commit: {emojiType: shortcode}",
			config: config.Config{Commit: config.Commit{EmojiType: config.EmojiTypeShortcode}},
		},
		{
			name:   "emojitype_character",
			data:   "commit: {emojiType: character}",
			config: config.Config{Commit: config.Commit{EmojiType: config.EmojiTypeCharacter}},
		},
		{
			name:   "emojitype_invalid",
			data:   "commit: {emojiType: invalid}",
			config: config.Config{Commit: config.Commit{EmojiType: config.EmojiTypeUnset}},
		},
		{
			name:   "signoff_empty",
			data:   "commit: {signoff:}",
			config: config.Config{Commit: config.Commit{Signoff: false}},
		},
		{
			name:   "signoff_false",
			data:   "commit: {signoff: false}",
			config: config.Config{Commit: config.Commit{Signoff: false}},
		},
		{
			name:   "signoff_true",
			data:   "commit: {signoff: true}",
			config: config.Config{Commit: config.Commit{Signoff: true}},
		},
		{
			name:   "signoff_invalid",
			data:   "commit: {signoff: invalid}",
			config: config.Config{Commit: config.Commit{Signoff: false}},
			err:    new(yaml.TypeError),
		},
		{
			name:   "theme_empty",
			data:   "view: {theme:}",
			config: config.Config{View: config.View{Theme: ""}},
		},
		{
			name:   "theme_valid",
			data:   "view: {theme: valid}",
			config: config.Config{View: config.View{Theme: "valid"}},
		},
		{
			name:   "authors_empty",
			data:   "authors:",
			config: config.Config{Authors: nil},
		},
		{
			name: "authors_one",
			data: "authors: [{name: John Doe, email: john.doe@example.com}]",
			config: config.Config{Authors: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
			}},
		},
		{
			name: "authors_one_name",
			data: "authors: [{name: John Doe}]",
			config: config.Config{Authors: []repository.User{
				{Name: "John Doe"},
			}},
		},
		{
			name: "authors_one_email",
			data: "authors: [{email: john.doe@example.com}]",
			config: config.Config{Authors: []repository.User{
				{Email: "john.doe@example.com"},
			}},
		},
		{
			name: "authors_multiple",
			data: "authors: [{name: John Doe, email: john.doe@example.com}, {name: John Doe, email: jdoe@example.org}]",
			config: config.Config{Authors: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
			}},
		},
		{
			name: "authors_invalid",
			data: "authors: [{name_email: John Doe = john.doe@example.com}]",
			config: config.Config{Authors: []repository.User{
				{Name: "", Email: ""},
			}},
		},
		{
			name: "default_author_one",
			data: "authors: [{name: John Doe, email: john.doe@example.com, default: true}]",
			config: config.Config{Authors: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com", Default: true},
			}},
		},
		{
			name: "default_author_second",
			data: `authors: [
				{name: John Doe, email: john.doe@example.com},
				{name: John Doe, email: jdoe@example.org, default: true}
				]`,
			config: config.Config{Authors: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org", Default: true},
			}},
		},
		{
			name: "default_author_multiple",
			data: `authors: [
				{name: John Doe, email: john.doe@example.com, default: true},
				{name: John Doe, email: jdoe@example.org, default: true}
			]`,
			config: config.Config{Authors: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com", Default: true},
				{Name: "John Doe", Email: "jdoe@example.org", Default: true},
			}},
		},
		{
			name: "default_author_none",
			data: `authors: [
				{name: John Doe, email: john.doe@example.com},
				{name: John Doe, email: jdoe@example.org}
			]`,
			config: config.Config{Authors: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
			}},
		},
		{
			name: "all",
			data: heredoc.Doc(`
				view:
				    focus: author
				    emojiSet: gitmoji
				    emojiSelector: below
				    compatibility: default
				    theme: builtin
				    colour: adaptive
				commit:
				    emojiType: shortcode
				    signoff: true
				authors:
				- name: John Doe
				  email: john.doe@example.com
			`),
			config: config.Config{
				View: config.View{
					Focus:         config.FocusAuthor,
					EmojiSet:      config.EmojiSetGitmoji,
					EmojiSelector: config.EmojiSelectorBelow,
					Compatibility: config.CompatibilityDefault,
					Theme:         "builtin",
					Colour:        config.ColourAdaptive,
				},
				Commit: config.Commit{
					EmojiType: config.EmojiTypeShortcode,
					Signoff:   true,
				},
				Authors: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
			},
		},
		{
			name: "all_invalid",
			data: heredoc.Doc(`
				view:
				    focus: author
				    emojiSet: gitmoji
				    emojiSelector: below
				    compatibility: none
				    theme: builtin
				    colour: blue
				commit:
				    emojiType: shortcode
				    signoff:
				authors:
				- name: John Doe
				  email:
			`),
			config: config.Config{
				View: config.View{
					Focus:         config.FocusAuthor,
					EmojiSet:      config.EmojiSetGitmoji,
					EmojiSelector: config.EmojiSelectorBelow,
					Compatibility: config.CompatibilityUnset,
					Theme:         "builtin",
					Colour:        config.ColourUnset,
				},
				Commit: config.Commit{
					EmojiType: config.EmojiTypeShortcode,
					Signoff:   false,
				},
				Authors: []repository.User{
					{Name: "John Doe", Email: ""},
				},
			},
		},
		{
			name: "all_incorrect_type",
			data: heredoc.Doc(`
				view:
				    focus: author
				    emojiSet: gitmoji
				    emojiSelector: below
				    compatibility: default
				    theme: builtin
				    colour: adaptive
				commit:
				    emojiType: shortcode
				    signoff: maybe
				authors:
				- name: John Doe
				  email: john.doe@example.com
			`),
			config: config.Config{
				View: config.View{
					Focus:         config.FocusAuthor,
					EmojiSet:      config.EmojiSetGitmoji,
					EmojiSelector: config.EmojiSelectorBelow,
					Compatibility: config.CompatibilityDefault,
					Theme:         "builtin",
					Colour:        config.ColourAdaptive,
				},
				Commit: config.Commit{
					EmojiType: config.EmojiTypeShortcode,
					Signoff:   false,
				},
				Authors: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
			},
			err: new(yaml.TypeError),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cfg, err := new(config.Config).Load(strings.NewReader(tt.data))
			if tt.err != nil {
				assert.NotNil(t, err)

				var typeError *yaml.TypeError
				assert.ErrorAs(t, err, &typeError)

				return
			}
			assert.Nil(t, err)

			assert.Equal(t, tt.config, cfg)
		})
	}
}
