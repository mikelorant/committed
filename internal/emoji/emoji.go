package emoji

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

type Emoji struct {
	EmojiID          int
	EmojiName        string `json:"name"`
	EmojiCharacter   string `json:"emoji"`
	EmojiDescription string `json:"description"`
	EmojiShortCode   string `json:"code"`
	EmojiEntity      string `json:"entity"`
	EmojiWide        bool
}

//go:embed emoji.json
var emojiJSON string

func New() ([]Emoji, error) {
	var e []Emoji

	r := strings.NewReader(emojiJSON)

	if err := json.NewDecoder(r).Decode(&e); err != nil {
		return nil, fmt.Errorf("unable to decode emojis: %w", err)
	}

	for i := range e {
		e[i].EmojiID = i
	}

	return e, nil
}

func (e Emoji) Title() string {
	var space string
	if runewidth.StringWidth(e.Character()) == 1 {
		space = " "
	}

	return fmt.Sprintf("%s%s - %s", e.EmojiCharacter, space, e.EmojiDescription)
}

func (e Emoji) Description() string { return e.EmojiDescription }
func (e Emoji) FilterValue() string { return e.EmojiName }
func (e Emoji) Name() string        { return e.EmojiName }
func (e Emoji) Character() string   { return e.EmojiCharacter }
func (e Emoji) ShortCode() string   { return e.EmojiShortCode }
func (e Emoji) Entity() string      { return e.EmojiEntity }
func (e Emoji) IsWide() bool        { return e.EmojiWide }
