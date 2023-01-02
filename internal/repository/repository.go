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
	Configer     Configer
	GlobalConfig func(scope config.Scope) (*config.Config, error)
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

type repositoryNotFoundError struct{}

var errRepositoryNotFound = errors.New("repository does not exist")

const repositoryPath string = "."

func New() (*Repository, error) {
	openOpts := git.PlainOpenOptions{
		DetectDotGit: true,
	}

	repo, err := git.PlainOpenWithOptions(repositoryPath, &openOpts)
	switch {
	case err == nil:
	case err.Error() == errRepositoryNotFound.Error():
		return nil, repositoryNotFoundError{}
	default:
		return nil, fmt.Errorf("unable to open git repository: %v: %w", repositoryPath, err)
	}

	return &Repository{
		Configer:     repo,
		GlobalConfig: config.LoadConfig,
		Remoter:      repo,
		Header:       repo,
		Brancher:     repo,
	}, nil
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

func (e repositoryNotFoundError) Error() string {
	return errRepositoryNotFound.Error()
}

func NotFoundError() error {
	return repositoryNotFoundError{}
}
