package emoji

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

type Emoji struct {
	Name        string `json:"name"`
	Character   string `json:"emoji"`
	Description string `json:"description"`
	ShortCode   string `json:"code"`
	Entity      string `json:"entity"`
}

//go:embed emoji.json
var emojiJSON string

func New() ([]Emoji, error) {
	var e []Emoji

	r := strings.NewReader(emojiJSON)

	if err := json.NewDecoder(r).Decode(&e); err != nil {
		return nil, fmt.Errorf("unable to decode emojis: %w", err)
	}

	return e, nil
}
