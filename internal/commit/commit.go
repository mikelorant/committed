package commit

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

type Commit struct {
	Options Options
	Applier Applier
	Emojier Emojier
	Loader  Loader
	Opener  Opener
	Repoer  Repoer
}

type (
	Applier func(repository.Commit, ...func(c *repository.Commit)) error
	Emojier func(...func(*emoji.Set)) *emoji.Set
	Loader  func(io.Reader) (config.Config, error)
	Opener  func(string) (io.Reader, error)
)

type Repoer interface {
	Open() error
	Describe() (repository.Description, error)
}

type Options struct {
	ConfigFile string
	DryRun     bool
	Amend      bool
}

type Request struct {
	Emoji   string
	Summary string
	Body    string
	Footer  string
	Author  repository.User
	Amend   bool
}

func New() Commit {
	return Commit{
		Applier: repository.Apply,
		Emojier: emoji.New,
		Loader:  config.Load,
		Opener:  FileOpen(),
		Repoer:  repository.New(),
	}
}

func (c *Commit) Configure(opts Options) (*State, error) {
	c.Options = opts

	repo, err := getRepo(c.Repoer)
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	cfg, err := getConfig(c.Opener, c.Loader, opts.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("unable to get config: %w", err)
	}

	return &State{
		Placeholders: placeholders(),
		Emojis:       getEmojis(c.Emojier, cfg),
		Repository:   repo,
		Config:       cfg,
		Options:      opts,
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

	opts := []func(c *repository.Commit){
		repository.WithAmend(req.Amend),
		repository.WithDryRun(c.Options.DryRun),
	}

	if err := c.Applier(com, opts...); err != nil {
		return fmt.Errorf("unable to apply commit: %w", err)
	}

	return nil
}

func FileOpen() func(string) (io.Reader, error) {
	return func(file string) (io.Reader, error) {
		var pathError *fs.PathError

		fh, err := os.Open(os.ExpandEnv(file))
		switch {
		case err == nil:
		case errors.As(err, &pathError):
			return strings.NewReader(""), nil
		default:
			return nil, err
		}

		return fh, nil
	}
}

func getRepo(repo Repoer) (repository.Description, error) {
	if err := repo.Open(); err != nil {
		return repository.Description{}, fmt.Errorf("unable to open repository: %w", err)
	}

	desc, err := repo.Describe()
	if err != nil {
		return repository.Description{}, fmt.Errorf("unable to describe repository: %w", err)
	}

	return desc, nil
}

func getConfig(open Opener, load Loader, file string) (config.Config, error) {
	r, err := open(file)
	if err != nil {
		return config.Config{}, fmt.Errorf("unable to open config file: %v: %w", file, err)
	}

	cfg, err := load(r)
	if err != nil {
		return config.Config{}, fmt.Errorf("unable to load config file: %w", err)
	}

	return cfg, nil
}

func getEmojis(emojier Emojier, cfg config.Config) *emoji.Set {
	prof := EmojiConfigToEmojiProfile(cfg.View.EmojiSet)
	fn := emoji.WithEmojiSet(prof)

	return emojier(fn)
}
