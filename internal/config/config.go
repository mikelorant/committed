package config

import (
	"errors"
	"fmt"
	"io"

	"github.com/mikelorant/committed/internal/repository"
	"gopkg.in/yaml.v3"
)

type Config struct {
	View    View              `yaml:"view,omitempty,flow"`
	Commit  Commit            `yaml:"commit,omitempty,flow"`
	Authors []repository.User `yaml:"authors,omitempty,flow"`
}

type View struct {
	Focus           Focus         `yaml:"focus,omitempty,flow"`
	EmojiSet        EmojiSet      `yaml:"emojiSet,omitempty,flow"`
	EmojiSelector   EmojiSelector `yaml:"emojiSelector,omitempty,flow"`
	Compatibility   Compatibility `yaml:"compatibility,omitempty,flow"`
	Theme           string        `yaml:"theme,omitempty,flow"`
	Colour          Colour        `yaml:"colour,omitempty,flow"`
	HighlightActive bool          `yaml:"highlightActive,omitempty,flow"`
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
