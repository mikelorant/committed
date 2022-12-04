package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type Repository struct {
	gitRepository *git.Repository
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

	return &Repository{
		gitRepository: repo,
	}, nil
}
