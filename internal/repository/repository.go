package repository

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type Configer interface {
	Config() (*config.Config, error)
}

type Repository struct {
	gitRepository *git.Repository

	Branch     Branch
	Remote     Remote
	HeadCommit Commit

	Configer     Configer
	GlobalConfig func(scope config.Scope) (*config.Config, error)
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

	remote, err := NewRemote(repo)
	if err != nil {
		return nil, fmt.Errorf("unable to initialise remote: %w", err)
	}

	commit, err := NewHeadCommit(repo)
	if err != nil {
		return nil, fmt.Errorf("unable to get head commit: %w", err)
	}

	return &Repository{
		gitRepository: repo,
		Configer:      repo,
		GlobalConfig:  config.LoadConfig,
		Branch:        branch,
		Remote:        remote,
		HeadCommit:    commit,
	}, nil
}

func (e repositoryNotFoundError) Error() string {
	return errRepositoryNotFound.Error()
}

func NotFoundError() error {
	return repositoryNotFoundError{}
}
