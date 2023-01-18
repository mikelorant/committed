package config

import (
	"errors"
	"fmt"
	"io"

	"github.com/mikelorant/committed/internal/repository"
	"gopkg.in/yaml.v3"
)

type Config struct {
	View    View              `yaml:"view"`
	Commit  Commit            `yaml:"commit"`
	Authors []repository.User `yaml:"authors"`
}

type View struct {
	Focus         Focus         `yaml:"focus"`
	EmojiSet      EmojiSet      `yaml:"emojiSet"`
	EmojiSelector EmojiSelector `yaml:"emojiSelector"`
	Compatibility Compatibility `yaml:"compatibility"`
	Theme         string        `yaml:"theme"`
	Colour        Colour        `yaml:"colour"`
}

type Commit struct {
	EmojiType EmojiType `yaml:"emojiType"`
	Signoff   bool      `yaml:"signoff"`
}

func Load(fh io.Reader) (Config, error) {
	var c Config

	err := yaml.NewDecoder(fh).Decode(&c)
	switch {
	case err == nil:
	case errors.Is(err, io.EOF):
	default:
		return c, fmt.Errorf("unable to decode config: %w", err)
	}

	return c, nil
}
