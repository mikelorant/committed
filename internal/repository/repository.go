package repository

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type Configer interface {
	Config() (*config.Config, error)
}

type Remoter interface {
	Remotes() ([]*git.Remote, error)
}

type Header interface {
	Head() (*plumbing.Reference, error)
	CommitObject(h plumbing.Hash) (*object.Commit, error)
}

type Brancher interface {
	Configer
	Head() (*plumbing.Reference, error)
	References() (storer.ReferenceIter, error)
}

type Repository struct {
	Opener       func(string, *git.PlainOpenOptions) (*git.Repository, error)
	GlobalConfig func(config.Scope) (*config.Config, error)
	Configer     Configer
	Remoter      Remoter
	Header       Header
	Brancher     Brancher
}

type Description struct {
	Users   []User
	Remotes []string
	Head    Head
	Branch  Branch
}

const repositoryPath string = "."

func New() *Repository {
	return &Repository{
		GlobalConfig: config.LoadConfig,
		Opener:       git.PlainOpenWithOptions,
	}
}

func (r *Repository) Open() error {
	openOpts := git.PlainOpenOptions{
		DetectDotGit: true,
	}

	repo, err := r.Opener(repositoryPath, &openOpts)
	switch {
	case err == nil:
	case errors.Is(err, git.ErrRepositoryNotExists):
		return err
	default:
		return fmt.Errorf("unable to open git repository: %v: %w", repositoryPath, err)
	}

	r.Configer = repo
	r.Remoter = repo
	r.Header = repo
	r.Brancher = repo

	return nil
}

func (r *Repository) Describe() (Description, error) {
	us, err := r.Users()
	if err != nil {
		return Description{}, fmt.Errorf("unable to get users: %w", err)
	}

	rs, err := r.Remotes()
	if err != nil {
		return Description{}, fmt.Errorf("unable to get remotes: %w", err)
	}

	h, err := r.Head()
	if err != nil {
		return Description{}, fmt.Errorf("unable to get head commit: %w", err)
	}

	b, err := r.Branch()
	if err != nil {
		return Description{}, fmt.Errorf("unable to get branch: %w", err)
	}

	return Description{
		Users:   us,
		Remotes: rs,
		Head:    h,
		Branch:  b,
	}, nil
}
