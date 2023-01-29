package config_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type badBuffer struct {
	buf bytes.Buffer
	err error
}

func (b *badBuffer) Write(p []byte) (int, error) {
	if b.err != nil {
		return 0, b.err
	}

	return b.buf.Write(p)
}

func (b *badBuffer) Read(p []byte) (int, error) {
	return b.buf.Read(p)
}

func (b *badBuffer) Close() error {
	return nil
}

var errMock = errors.New("error")

func TestLoad(t *testing.T) {
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
			name:   "highlight_empty",
			data:   "view: {highlightActive:}",
			config: config.Config{View: config.View{HighlightActive: false}},
		},
		{
			name:   "highlight_true",
			data:   "view: {highlightActive: true}",
			config: config.Config{View: config.View{HighlightActive: true}},
		},
		{
			name:   "highlight_false",
			data:   "view: {highlightActive: false}",
			config: config.Config{View: config.View{HighlightActive: false}},
		},
		{
			name: "highlight_invalid",
			data: "view: {highlightActive: invalid}",
			err:  new(yaml.TypeError),
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

func TestSave(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		config func(*config.Config)
		data   string
		err    error
	}{
		{
			name: "empty",
			data: "{}",
		},
		{
			name:   "view_focus_unset",
			config: func(c *config.Config) { c.View.Focus = config.FocusUnset },
			data:   "{}",
		},
		{
			name:   "view_focus_author",
			config: func(c *config.Config) { c.View.Focus = config.FocusAuthor },
			data:   "view: {focus: author}",
		},
		{
			name:   "view_focus_emoji",
			config: func(c *config.Config) { c.View.Focus = config.FocusEmoji },
			data:   "view: {focus: emoji}",
		},
		{
			name:   "view_focus_summary",
			config: func(c *config.Config) { c.View.Focus = config.FocusSummary },
			data:   "view: {focus: summary}",
		},
		{
			name:   "view_compatibility_unset",
			config: func(c *config.Config) { c.View.Compatibility = config.CompatibilityUnset },
			data:   "{}",
		},
		{
			name:   "view_compatibility_default",
			config: func(c *config.Config) { c.View.Compatibility = config.CompatibilityDefault },
			data:   "view: {compatibility: default}",
		},
		{
			name:   "view_compatibility_ttyd",
			config: func(c *config.Config) { c.View.Compatibility = config.CompatibilityTtyd },
			data:   "view: {compatibility: ttyd}",
		},
		{
			name:   "view_compatibility_kitty",
			config: func(c *config.Config) { c.View.Compatibility = config.CompatibilityKitty },
			data:   "view: {compatibility: kitty}",
		},
		{
			name:   "view_colour_unset",
			config: func(c *config.Config) { c.View.Colour = config.ColourUnset },
			data:   "{}",
		},
		{
			name:   "view_colour_adaptive",
			config: func(c *config.Config) { c.View.Colour = config.ColourAdaptive },
			data:   "view: {colour: adaptive}",
		},
		{
			name:   "view_colour_dark",
			config: func(c *config.Config) { c.View.Colour = config.ColourDark },
			data:   "view: {colour: dark}",
		},
		{
			name:   "view_colour_light",
			config: func(c *config.Config) { c.View.Colour = config.ColourLight },
			data:   "view: {colour: light}",
		},
		{
			name:   "view_emojiset_unset",
			config: func(c *config.Config) { c.View.EmojiSet = config.EmojiSetUnset },
			data:   "{}",
		},
		{
			name:   "view_emojiSet_gitmoji",
			config: func(c *config.Config) { c.View.EmojiSet = config.EmojiSetGitmoji },
			data:   "view: {emojiSet: gitmoji}",
		},
		{
			name:   "view_emojiset_devmoji",
			config: func(c *config.Config) { c.View.EmojiSet = config.EmojiSetDevmoji },
			data:   "view: {emojiSet: devmoji}",
		},
		{
			name:   "view_emojiset_emojilog",
			config: func(c *config.Config) { c.View.EmojiSet = config.EmojiSetEmojiLog },
			data:   "view: {emojiSet: emojilog}",
		},
		{
			name:   "view_emojiselector_unset",
			config: func(c *config.Config) { c.View.EmojiSelector = config.EmojiSelectorUnset },
			data:   "{}",
		},
		{
			name:   "view_emojiselector_below",
			config: func(c *config.Config) { c.View.EmojiSelector = config.EmojiSelectorBelow },
			data:   "view: {emojiSelector: below}",
		},
		{
			name:   "view_emojiselector_above",
			config: func(c *config.Config) { c.View.EmojiSelector = config.EmojiSelectorAbove },
			data:   "view: {emojiSelector: above}",
		},
		{
			name:   "view_emojitype_unset",
			config: func(c *config.Config) { c.Commit.EmojiType = config.EmojiTypeUnset },
			data:   "{}",
		},
		{
			name:   "view_emojitype_below",
			config: func(c *config.Config) { c.Commit.EmojiType = config.EmojiTypeShortcode },
			data:   "commit: {emojiType: shortcode}",
		},
		{
			name:   "view_emojitype_above",
			config: func(c *config.Config) { c.Commit.EmojiType = config.EmojiTypeCharacter },
			data:   "commit: {emojiType: character}",
		},
		{
			name: "authors_one",
			config: func(c *config.Config) {
				c.Authors = []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				}
			},
			data: "authors: [{name: John Doe, email: john.doe@example.com}]",
		},
		{
			name: "authors_multiple",
			config: func(c *config.Config) {
				c.Authors = []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "John Doe", Email: "jdoe@example.org"},
				}
			},
			data: "authors: [{name: John Doe, email: john.doe@example.com}, {name: John Doe, email: jdoe@example.org}]",
		},
		{
			name:   "theme_empty",
			config: func(c *config.Config) { c.View.Theme = "" },
			data:   "{}",
		},
		{
			name:   "theme_test",
			config: func(c *config.Config) { c.View.Theme = "test" },
			data:   "view: {theme: test}",
		},
		{
			name:   "signoff_false",
			config: func(c *config.Config) { c.Commit.Signoff = false },
			data:   "{}",
		},
		{
			name:   "signoff_true",
			config: func(c *config.Config) { c.Commit.Signoff = true },
			data:   "commit: {signoff: true}",
		},
		{
			name:   "highlightactive_false",
			config: func(c *config.Config) { c.View.HighlightActive = false },
			data:   "{}",
		},
		{
			name:   "highlightactive_true",
			config: func(c *config.Config) { c.View.HighlightActive = true },
			data:   "view: {highlightActive: true}",
		},
		{
			name: "error",
			err:  errMock,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var cfg config.Config

			if tt.config != nil {
				tt.config(&cfg)
			}

			buf := badBuffer{
				err: tt.err,
			}

			err := new(config.Config).Save(&buf, cfg)
			if tt.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "unable to encode config: yaml: write error: error")
				return
			}
			assert.NoError(t, err)

			data, _ := io.ReadAll(&buf)
			assert.Equal(t, tt.data, strings.TrimSpace(string(data)))
		})
	}
}
