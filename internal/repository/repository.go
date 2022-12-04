package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type Repository struct {
	gitRepository *git.Repository
	User          User
}

const repositoryPath string = "."

func New() (*Repository, error) {
	openOpts := git.PlainOpenOptions{
		DetectDotGit: true,
	}

	repo, err := git.PlainOpenWithOptions(repositoryPath, &openOpts)
	if err != nil {
		return nil, fmt.Errorf("unable to open git repository: %v: %w", repositoryPath, err)
	}

	user, err := NewUser(repo)
	if err != nil {
		return nil, fmt.Errorf("unable to initialise user: %w", err)
	}

	return &Repository{
		gitRepository: repo,
		User:          user,
	}, nil
}
