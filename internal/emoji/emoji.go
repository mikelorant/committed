package emoji

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/forPelevin/gomoji"
	"github.com/goccy/go-yaml"
)

type Emoji struct {
	Name        string `json:"name"`
	Character   string `json:"emoji"`
	Description string `json:"description"`
	Characters  int    `json:"characters"`
	Codepoint   string `json:"codepoint"`
	Hex         string `json:"hex"`
	Shortcode   string `json:"shortcode"`
	Variant     bool   `json:"variant"`
	ZWJ         bool   `json:"zwj"`
}

type NullEmoji struct {
	Valid bool
	Emoji Emoji
}

//go:embed gitmoji.yaml
var gitmoji string

func New() ([]Emoji, error) {
	var e []Emoji

	r := strings.NewReader(gitmoji)

	if err := yaml.NewDecoder(r).Decode(&e); err != nil {
		return nil, fmt.Errorf("unable to decode emojis: %w", err)
	}

	return e, nil
}

func Find(str string, es []Emoji) NullEmoji {
	switch {
	case HasCharacter(str):
		return FindByCharacter(str, es)
	case HasShortcode(str):
		return FindByShortcode(str, es)
	default:
		return NullEmoji{}
	}
}

func FindByCharacter(str string, es []Emoji) NullEmoji {
	for _, e := range es {
		if e.Character == str {
			return NullEmoji{
				Valid: true,
				Emoji: e,
			}
		}
	}

	return NullEmoji{}
}

func FindByShortcode(str string, es []Emoji) NullEmoji {
	for _, e := range es {
		if e.Shortcode == str {
			return NullEmoji{
				Valid: true,
				Emoji: e,
			}
		}
	}

	return NullEmoji{}
}

func Has(str string) bool {
	return HasCharacter(str) || HasShortcode(str)
}

func HasCharacter(str string) bool {
	return gomoji.ContainsEmoji(str)
}

func HasShortcode(str string) bool {
	if len(str) <= 2 {
		return false
	}

	if strings.Count(str, ":") > 2 {
		return false
	}

	if string(str[0]) == ":" && string(str[len(str)-1]) == ":" {
		return true
	}

	return false
}
