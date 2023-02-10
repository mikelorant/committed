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
	CompatibilityDefault
	CompatibilityTtyd
	CompatibilityKitty
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

func (c *Compatibility) UnmarshalYAML(value *yaml.Node) error {
	*c = ParseCompatibility(value.Value)

	return nil
}

func (c Compatibility) MarshalYAML() (interface{}, error) {
	return []string{
		"",
		"default",
		"ttyd",
		"kitty",
	}[c], nil
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
		"":        CompatibilityUnset,
		"default": CompatibilityDefault,
		"ttyd":    CompatibilityTtyd,
		"kitty":   CompatibilityKitty,
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
