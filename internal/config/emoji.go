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
	EmojiSetCommitted
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

func (e EmojiSet) MarshalYAML() (interface{}, error) {
	return []string{
		"",
		"committed",
		"gitmoji",
		"devmoji",
		"emojilog",
	}[e], nil
}

func (e EmojiSet) Default() int {
	return 2
}

func (e EmojiSet) Index() int {
	if e == EmojiSetUnset {
		return e.Default()
	}

	return int(e)
}

func (e *EmojiType) UnmarshalYAML(value *yaml.Node) error {
	*e = ParseEmojiType(value.Value)

	return nil
}

func (e EmojiType) MarshalYAML() (interface{}, error) {
	return []string{
		"",
		"shortcode",
		"character",
	}[e], nil
}

func (e EmojiType) Default() int {
	return 1
}

func (e EmojiType) Index() int {
	if e == EmojiTypeUnset {
		return e.Default()
	}

	return int(e)
}

func (e *EmojiSelector) UnmarshalYAML(value *yaml.Node) error {
	*e = ParseEmojiSelector(value.Value)

	return nil
}

func (e EmojiSelector) MarshalYAML() (interface{}, error) {
	return []string{
		"",
		"below",
		"above",
	}[e], nil
}

func (e EmojiSelector) Default() int {
	return 1
}

func (e EmojiSelector) Index() int {
	if e == EmojiSelectorUnset {
		return e.Default()
	}

	return int(e)
}

func ParseEmojiSet(str string) EmojiSet {
	emojiSet := map[string]EmojiSet{
		"":          EmojiSetUnset,
		"committed": EmojiSetCommitted,
		"gitmoji":   EmojiSetGitmoji,
		"devmoji":   EmojiSetDevmoji,
		"emojilog":  EmojiSetEmojiLog,
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
