package commit

import (
	"fmt"
	"io"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

type Commit struct {
	Options  Options
	Emojier  Emojier
	Configer Configer
	Opener   Opener
	Repoer   Repoer
	Creator  Creator
}

type (
	Applier func(repository.Commit, ...func(c *repository.Commit)) error
	Emojier func(...func(*emoji.Set)) *emoji.Set
	Opener  func(string) (io.Reader, error)
	Creator func(string) (io.WriteCloser, error)
)

type Repoer interface {
	Open() error
	Describe() (repository.Description, error)
	Apply(repository.Commit, ...func(c *repository.Commit)) error
}

type Configer interface {
	Load(io.Reader) (config.Config, error)
	Save(io.WriteCloser, config.Config) error
}

type Options struct {
	ConfigFile string
	DryRun     bool
	Amend      bool
}

type Request struct {
	Apply   bool
	Emoji   string
	Summary string
	Body    string
	Footer  string
	Author  repository.User
	Amend   bool
}

func New() Commit {
	return Commit{
		Emojier:  emoji.New,
		Repoer:   repository.New(),
		Configer: new(config.Config),
		Opener:   FileOpen(),
		Creator:  FileCreate(),
	}
}

func (c *Commit) Configure(opts Options) (*State, error) {
	c.Options = opts

	repo, err := getRepo(c.Repoer)
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	cfg, err := getConfig(c.Opener, c.Configer, opts.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("unable to get config: %w", err)
	}

	if !FileExists(opts.ConfigFile) {
		if err := setConfig(c.Creator, c.Configer, opts.ConfigFile, cfg); err != nil {
			return nil, fmt.Errorf("unable to set config: %w", err)
		}
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

	if !req.Apply {
		return nil
	}

	opts := []func(c *repository.Commit){
		repository.WithAmend(req.Amend),
		repository.WithDryRun(c.Options.DryRun),
	}

	if err := c.Repoer.Apply(com, opts...); err != nil {
		return fmt.Errorf("unable to apply commit: %w", err)
	}

	return nil
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

func getConfig(open Opener, configer Configer, file string) (config.Config, error) {
	r, err := open(file)
	if err != nil {
		return config.Config{}, fmt.Errorf("unable to open config file: %v: %w", file, err)
	}

	cfg, err := configer.Load(r)
	if err != nil {
		return config.Config{}, fmt.Errorf("unable to load config file: %w", err)
	}

	return cfg, nil
}

func setConfig(create Creator, configer Configer, file string, cfg config.Config) error {
	w, err := create(file)
	if err != nil {
		return fmt.Errorf("unable to create config: %w", err)
	}

	if err := configer.Save(w, cfg); err != nil {
		return fmt.Errorf("unable to save config: %w", err)
	}

	return nil
}

func getEmojis(emojier Emojier, cfg config.Config) *emoji.Set {
	prof := EmojiConfigToEmojiProfile(cfg.View.EmojiSet)
	fn := emoji.WithEmojiSet(prof)

	return emojier(fn)
}
