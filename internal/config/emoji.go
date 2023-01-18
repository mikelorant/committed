package config

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type (
	EmojiSet      int
	EmojiType     int
	EmojiSelector int
)

const (
	EmojiSetUnset EmojiSet = iota
	EmojiSetGitmoji
	EmojiSetDevmoji
	EmojiSetEmojiLog
)

const (
	EmojiSelectorUnset EmojiSelector = iota
	EmojiSelectorBelow
	EmojiSelectorAbove
)

const (
	EmojiTypeUnset EmojiType = iota
	EmojiTypeShortcode
	EmojiTypeCharacter
)

func (e *EmojiSet) UnmarshalYAML(value *yaml.Node) error {
	*e = ParseEmojiSet(value.Value)

	return nil
}

func (e *EmojiType) UnmarshalYAML(value *yaml.Node) error {
	*e = ParseEmojiType(value.Value)

	return nil
}

func (e *EmojiSelector) UnmarshalYAML(value *yaml.Node) error {
	*e = ParseEmojiSelector(value.Value)

	return nil
}

func ParseEmojiSet(str string) EmojiSet {
	emojiSet := map[string]EmojiSet{
		"":         EmojiSetUnset,
		"gitmoji":  EmojiSetGitmoji,
		"devmoji":  EmojiSetDevmoji,
		"emojilog": EmojiSetEmojiLog,
	}

	return emojiSet[strings.ToLower(str)]
}

func ParseEmojiType(str string) EmojiType {
	emojiSet := map[string]EmojiType{
		"":          EmojiTypeUnset,
		"shortcode": EmojiTypeShortcode,
		"character": EmojiTypeCharacter,
	}

	return emojiSet[strings.ToLower(str)]
}

func ParseEmojiSelector(str string) EmojiSelector {
	emojiSelector := map[string]EmojiSelector{
		"":      EmojiSelectorUnset,
		"above": EmojiSelectorAbove,
		"below": EmojiSelectorBelow,
	}

	return emojiSelector[strings.ToLower(str)]
}
