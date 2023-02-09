package snapshot

import (
	"errors"
	"fmt"
	"io"

	"github.com/mikelorant/committed/internal/repository"
	"gopkg.in/yaml.v3"
)

type Snapshot struct {
	Emoji   string          `yaml:"emoji,omitempty"`
	Summary string          `yaml:"summary,omitempty"`
	Body    string          `yaml:"body,omitempty"`
	Footer  string          `yaml:"footer,omitempty"`
	Author  repository.User `yaml:"author,omitempty"`
	Amend   bool            `yaml:"amend,omitempty"`
	Restore bool            `yaml:"restore,omitempty"`
}

var (
	errReader = errors.New("empty reader")
	errWriter = errors.New("empty writer")
)

func (s *Snapshot) Load(fh io.Reader) (Snapshot, error) {
	var snap Snapshot

	if fh == nil {
		return snap, errReader
	}

	err := yaml.NewDecoder(fh).Decode(&snap)
	switch {
	case err == nil:
	case errors.Is(err, io.EOF):
	default:
		return snap, fmt.Errorf("unable to decode snapshot: %w", err)
	}

	return snap, nil
}

func (s *Snapshot) Save(fh io.WriteCloser, snap Snapshot) error {
	if fh == nil {
		return errWriter
	}

	err := yaml.NewEncoder(fh).Encode(&snap)
	if err != nil {
		return fmt.Errorf("unable to encode snapshot: %w", err)
	}
	defer fh.Close()

	return nil
}
