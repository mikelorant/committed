package commit

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/snapshot"
)

type Commit struct {
	Options     Options
	Emojier     Emojier
	Configer    Configer
	Snapshotter Snapshotter
	Opener      Opener
	ReadFiler   ReadFiler
	Repoer      Repoer
	Creator     Creator
	Saver       Saver
}

type (
	Applier   func(repository.Commit, ...func(c *repository.Commit)) error
	Creator   func(string) (io.WriteCloser, error)
	Emojier   func(...func(*emoji.Set)) *emoji.Set
	Opener    func(string) (io.Reader, error)
	ReadFiler func(string) ([]byte, error)
	Saver     func(io.WriteCloser, snapshot.Snapshot) error
)

type Repoer interface {
	Open() error
	Describe() (repository.Description, error)
	Apply(repository.Commit) error
	IgnoreGlobalConfig()
}

type Configer interface {
	Load(io.Reader) (config.Config, error)
	Save(io.WriteCloser, config.Config) error
}

type Snapshotter interface {
	Load(io.Reader) (snapshot.Snapshot, error)
	Save(io.WriteCloser, snapshot.Snapshot) error
}

type Options struct {
	ConfigFile   string
	SnapshotFile string
	DryRun       bool
	Amend        bool
	Mode         Mode
	File         FileOptions
}

type FileOptions struct {
	MessageFile string
	Source      string
	SHA         string
}

type Request struct {
	Apply       bool
	Emoji       string
	Summary     string
	Body        string
	RawBody     string
	Footer      string
	Author      repository.User
	Amend       bool
	DryRun      bool
	File        bool
	MessageFile string
}

type Mode int

const (
	ModeUnset Mode = iota
	ModeCommit
	ModeEditor
	ModeHook
)

func New() Commit {
	return Commit{
		Emojier:     emoji.New,
		Repoer:      repository.New(),
		Configer:    new(config.Config),
		Snapshotter: new(snapshot.Snapshot),
		Opener:      FileOpen(),
		ReadFiler:   os.ReadFile,
		Creator:     FileCreate(),
	}
}

func (c *Commit) Configure(opts Options) (*State, error) {
	cfg, err := getConfig(c.Opener, c.Configer, opts.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("unable to get config: %w", err)
	}

	if cfg.View.IgnoreGlobalAuthor {
		c.Repoer.IgnoreGlobalConfig()
	}

	repo, err := getRepo(c.Repoer)
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	if !FileExists(opts.ConfigFile) {
		if err := setConfig(c.Creator, c.Configer, opts.ConfigFile, cfg); err != nil {
			return nil, fmt.Errorf("unable to set config: %w", err)
		}
	}

	snap, err := getSnapshot(c.Opener, c.Snapshotter, opts.SnapshotFile)
	if err != nil {
		return nil, fmt.Errorf("unable to get snapshot: %w", err)
	}

	var file File
	if opts.Mode > ModeCommit {
		file, err = readFile(c.ReadFiler, opts)
		if err != nil {
			return nil, fmt.Errorf("unable to read message file: %w", err)
		}
	}

	if file.Amend {
		opts.Amend = true
	}

	c.Options = opts

	return &State{
		Placeholders: placeholders(),
		Emojis:       getEmojis(c.Emojier, cfg),
		Repository:   repo,
		Config:       cfg,
		Snapshot:     snap,
		Options:      opts,
		File:         file,
	}, nil
}

func (c *Commit) Apply(req *Request) error {
	if req == nil {
		return nil
	}

	com := repository.Commit{
		Author:      UserToAuthor(req.Author),
		Subject:     EmojiSummaryToSubject(req.Emoji, req.Summary),
		Body:        req.Body,
		Footer:      req.Footer,
		Amend:       req.Amend,
		DryRun:      req.DryRun,
		File:        req.File,
		MessageFile: req.MessageFile,
	}

	snap := snapshot.Snapshot{
		Emoji:   req.Emoji,
		Summary: req.Summary,
		Body:    req.RawBody,
		Footer:  req.Footer,
		Author:  req.Author,
		Amend:   req.Amend,
	}

	if !req.Apply {
		if err := setSnapshot(c.Creator, c.Snapshotter, c.Options.SnapshotFile, snap); err != nil {
			return fmt.Errorf("unable to set snapshot: %w", err)
		}

		return nil
	}

	if err := c.Repoer.Apply(com); err != nil {
		var exitErr *exec.ExitError

		if !errors.As(err, &exitErr) {
			return fmt.Errorf("unable to apply commit: %w", err)
		}

		snap.Restore = true

		if err := setSnapshot(c.Creator, c.Snapshotter, c.Options.SnapshotFile, snap); err != nil {
			return fmt.Errorf("unable to set snapshot: %w", err)
		}
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

func getSnapshot(open Opener, snapshotter Snapshotter, file string) (snapshot.Snapshot, error) {
	r, err := open(file)
	if err != nil {
		return snapshot.Snapshot{}, fmt.Errorf("unable to open snapshot: %v: %w", file, err)
	}

	snap, err := snapshotter.Load(r)
	if err != nil {
		return snapshot.Snapshot{}, fmt.Errorf("unable to load snapshot: %w", err)
	}

	return snap, nil
}

func setSnapshot(create Creator, snapshotter Snapshotter, file string, snap snapshot.Snapshot) error {
	w, err := create(file)
	if err != nil {
		return fmt.Errorf("unable to create snapshot: %w", err)
	}

	if err := snapshotter.Save(w, snap); err != nil {
		return fmt.Errorf("unable to save snapshot: %w", err)
	}

	return nil
}

func getEmojis(emojier Emojier, cfg config.Config) *emoji.Set {
	prof := EmojiConfigToEmojiProfile(cfg.View.EmojiSet)
	fn := emoji.WithEmojiSet(prof)

	return emojier(fn)
}

func readFile(readFile ReadFiler, opts Options) (File, error) {
	data, err := readFile(opts.File.MessageFile)
	if err != nil {
		return File{}, fmt.Errorf("unable to read file: %w", err)
	}

	msg := string(data)

	f := File{
		Message: msg,
	}

	if isAmend(msg, opts) {
		f.Amend = true
	}

	return f, nil
}

func isAmend(msg string, opts Options) bool {
	if opts.Mode == ModeHook && opts.File.SHA == "HEAD" {
		return true
	}

	r := strings.NewReader(msg)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "#") && strings.TrimSpace(scanner.Text()) != "" {
			return true
		}
	}

	return false
}
