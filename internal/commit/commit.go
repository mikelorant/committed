package commit

import (
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

type Commit struct {
	Config  Config
	Emoji   string
	Summary string
	Body    string
	Footer  string
	Author  repository.User

	options Options
	cmd     []string
}

type Config struct {
	Placeholders Placeholders
	Repository   repository.Description
	Emojis       []emoji.Emoji
	Emoji        emoji.NullEmoji
	Summary      string
	Body         string
	Amend        bool
}

type Options struct {
	Apply bool
	Amend bool
}

type Placeholders struct {
	Hash    string
	Summary string
	Body    string
}

//go:embed message.txt
var message string

var exitError *exec.ExitError

const (
	mockHash string = "1234567890abcdef1234567890abcdef1234567890"
	summary  string = "Capitalized, short (50 chars or less) summary"
)

func New(opts Options) (*Commit, error) {
	r, err := repository.New()
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	e := emoji.New()

	d, err := r.Describe()
	if err != nil {
		return nil, fmt.Errorf("unable to describe repository: %w", err)
	}

	placeholders := Placeholders{
		Hash:    mockHash,
		Summary: summary,
		Body:    message,
	}

	cfg := Config{
		Placeholders: placeholders,
		Emojis:       e.Emojis,
		Repository:   d,
		Amend:        opts.Amend,
	}

	return &Commit{
		Config:  cfg,
		options: opts,
	}, nil
}

func (c *Commit) Apply() error {
	com := repository.Commit{
		Author:  c.UserToAuthor(),
		Subject: c.EmojiSummaryToSubject(),
		Body:    c.Body,
		Footer:  c.Footer,
	}

	opts := []repository.CommitOptions{
		repository.WithAmend(c.options.Amend),
		repository.WithDryRun(!c.options.Apply),
	}

	if err := repository.Apply(com, opts...); err != nil {
		return fmt.Errorf("unable to apply commit: %w", err)
	}

	return nil
}
