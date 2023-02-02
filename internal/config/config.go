package config

import (
	"errors"
	"fmt"
	"io"

	"github.com/mikelorant/committed/internal/repository"
	"gopkg.in/yaml.v3"
)

type Config struct {
	View    View              `yaml:"view,omitempty"`
	Commit  Commit            `yaml:"commit,omitempty"`
	Authors []repository.User `yaml:"authors,omitempty"`
}

type View struct {
	Focus         Focus         `yaml:"focus,omitempty"`
	EmojiSet      EmojiSet      `yaml:"emojiSet,omitempty"`
	EmojiSelector EmojiSelector `yaml:"emojiSelector,omitempty"`
	Compatibility Compatibility `yaml:"compatibility,omitempty"`
	Theme         string        `yaml:"theme,omitempty"`
	Colour        Colour        `yaml:"colour,omitempty"`
}

type Commit struct {
	EmojiType EmojiType `yaml:"emojiType,omitempty"`
	Signoff   bool      `yaml:"signoff,omitempty"`
}

func (c *Config) Load(fh io.Reader) (Config, error) {
	var cfg Config

	err := yaml.NewDecoder(fh).Decode(&cfg)
	switch {
	case err == nil:
	case errors.Is(err, io.EOF):
	default:
		return cfg, fmt.Errorf("unable to decode config: %w", err)
	}

	return cfg, nil
}

func (c *Config) Save(fh io.WriteCloser, cfg Config) error {
	err := yaml.NewEncoder(fh).Encode(&cfg)
	if err != nil {
		return fmt.Errorf("unable to encode config: %w", err)
	}
	defer fh.Close()

	return nil
}
