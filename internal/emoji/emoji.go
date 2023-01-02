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

func HasEmoji(str string) bool {
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
