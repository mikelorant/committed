package emoji

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
)

type Emoji struct {
	Name        string `json:"name"`
	Character   string `json:"emoji"`
	Description string `json:"description"`
	Characters  int    `json:"characters"`
	Codepoint   string `json:"codepoint"`
	Hex         string `json:"hex"`
	ShortCode   string `json:"shortcode"`
	Variant     bool   `json:"variant"`
	ZWJ         bool   `json:"zwj"`
}

//go:embed emoji.yaml
var emojiYAML string

func New() ([]Emoji, error) {
	var e []Emoji

	r := strings.NewReader(emojiYAML)

	if err := yaml.NewDecoder(r).Decode(&e); err != nil {
		return nil, fmt.Errorf("unable to decode emojis: %w", err)
	}

	return e, nil
}
