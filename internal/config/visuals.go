package config

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type (
	Compatibility int
	Theme         int
	Colour        int
)

const (
	FocusUnset Focus = iota
	FocusAuthor
	FocusEmoji
	FocusSummary
)

const (
	CompatibilityUnset Compatibility = iota
	CompatibilityUnicode14
	CompatibilityUnicode9
)

const (
	ColourUnset Colour = iota
	ColourAdaptive
	ColourDark
	ColourLight
)

type Focus int

func (f *Focus) UnmarshalYAML(value *yaml.Node) error {
	*f = ParseFocus(value.Value)

	return nil
}

func (f Focus) MarshalYAML() (interface{}, error) {
	return []string{
		"",
		"author",
		"emoji",
		"summary",
	}[f], nil
}

func (f Focus) Default() int {
	return 2
}

func (f Focus) Index() int {
	if f == FocusUnset {
		return f.Default()
	}

	return int(f)
}

func (c *Compatibility) UnmarshalYAML(value *yaml.Node) error {
	*c = ParseCompatibility(value.Value)

	return nil
}

func (c Compatibility) MarshalYAML() (interface{}, error) {
	return []string{
		"",
		"unicode14",
		"unicode9",
	}[c], nil
}

func (c Compatibility) Default() int {
	return 1
}

func (c Compatibility) Index() int {
	if c == CompatibilityUnset {
		return c.Default()
	}

	return int(c)
}

func (c *Colour) UnmarshalYAML(value *yaml.Node) error {
	*c = ParseColour(value.Value)

	return nil
}

func (c Colour) MarshalYAML() (interface{}, error) {
	return []string{
		"",
		"adaptive",
		"dark",
		"light",
	}[c], nil
}

func (c Colour) Default() int {
	return 1
}

func (c Colour) Index() int {
	if c == ColourUnset {
		return c.Default()
	}

	return int(c)
}

func ParseFocus(str string) Focus {
	focus := map[string]Focus{
		"":        FocusUnset,
		"author":  FocusAuthor,
		"emoji":   FocusEmoji,
		"summary": FocusSummary,
	}

	return focus[strings.ToLower(str)]
}

func ParseCompatibility(str string) Compatibility {
	compatibility := map[string]Compatibility{
		"":          CompatibilityUnset,
		"unicode14": CompatibilityUnicode14,
		"unicode9":  CompatibilityUnicode9,
	}

	return compatibility[strings.ToLower(str)]
}

func ParseColour(str string) Colour {
	colour := map[string]Colour{
		"":         ColourUnset,
		"adaptive": ColourAdaptive,
		"dark":     ColourDark,
		"light":    ColourLight,
	}

	return colour[strings.ToLower(str)]
}
