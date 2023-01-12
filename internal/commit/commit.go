package commit

import (
	_ "embed"
	"fmt"

	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

type Commit struct {
	Config  Config
	Options Options
	Applier func(c repository.Commit, opts ...repository.CommitOptions) error
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

type Request struct {
	Emoji   string
	Summary string
	Body    string
	Footer  string
	Author  repository.User
}

type Options struct {
	DryRun bool
	Amend  bool
}

type Placeholders struct {
	Hash    string
	Summary string
	Body    string
}

//go:embed message.txt
var message string

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
		Options: opts,
		Applier: repository.Apply,
	}, nil
}

func (c *Commit) Apply(req *Request) error {
	if req == nil {
		return nil
	}

	com := repository.Commit{
		Author:  UserToAuthor(req.Author),
		Subject: EmojiSummaryToSubject(req.Emoji, req.Summary),
		Body:    req.Body,
		Footer:  req.Footer,
	}

	opts := []repository.CommitOptions{
		repository.WithAmend(c.Options.Amend),
		repository.WithDryRun(c.Options.DryRun),
	}

	if err := c.Applier(com, opts...); err != nil {
		return fmt.Errorf("unable to apply commit: %w", err)
	}

	return nil
}
