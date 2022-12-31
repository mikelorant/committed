package repository

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
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

type Repository struct {
	gitRepository *git.Repository

	Branch Branch

	Configer     Configer
	GlobalConfig func(scope config.Scope) (*config.Config, error)
	Remoter      Remoter
	Header       Header
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

	branch, err := NewBranch(repo)
	if err != nil {
		return nil, fmt.Errorf("unable to initialise branch: %w", err)
	}

	return &Repository{
		gitRepository: repo,
		Configer:      repo,
		GlobalConfig:  config.LoadConfig,
		Remoter:       repo,
		Header:        repo,
		Branch:        branch,
	}, nil
}

func (e repositoryNotFoundError) Error() string {
	return errRepositoryNotFound.Error()
}

func NotFoundError() error {
	return repositoryNotFoundError{}
}
