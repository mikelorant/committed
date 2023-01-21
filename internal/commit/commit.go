package commit

import (
	_ "embed"
	"fmt"

	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

type Commit struct {
	Options   Options
	Applier   func(repository.Commit, ...func(c *repository.Commit)) error
	Repoer    func() (*repository.Repository, error)
	Emojier   func(...func(*emoji.Set)) *emoji.Set
	Describer Describer

	config Config
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
var PlaceholderMessage string

const (
	PlaceholderHash    string = "1234567890abcdef1234567890abcdef1234567890"
	PlaceholderSummary string = "Capitalized, short (50 chars or less) summary"
)

type Describer interface {
	Describe() (repository.Description, error)
}

func New() Commit {
	return Commit{
		Applier: repository.Apply,
		Repoer:  repository.New,
		Emojier: emoji.New,
	}
}

func (c *Commit) Configure(opts Options) (*Config, error) {
	c.Options = opts

	r, err := c.Repoer()
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	if c.Describer == nil {
		c.Describer = r
	}

	e := c.Emojier()

	d, err := c.Describer.Describe()
	if err != nil {
		return nil, fmt.Errorf("unable to describe repository: %w", err)
	}

	placeholders := Placeholders{
		Hash:    PlaceholderHash,
		Summary: PlaceholderSummary,
		Body:    PlaceholderMessage,
	}

	c.config = Config{
		Placeholders: placeholders,
		Emojis:       e.Emojis,
		Repository:   d,
		Amend:        opts.Amend,
	}

	return &c.config, nil
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

	opts := []func(c *repository.Commit){
		repository.WithAmend(c.Options.Amend),
		repository.WithDryRun(c.Options.DryRun),
	}

	if err := c.Applier(com, opts...); err != nil {
		return fmt.Errorf("unable to apply commit: %w", err)
	}

	return nil
}
