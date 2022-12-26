package repository

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
)

type Commit struct {
	Hash    string
	Author  User
	When    time.Time
	Message string
}

func NewHeadCommit(r *git.Repository) (Commit, error) {
	h, err := r.Head()

	switch {
	case err == nil:
	case err.Error() == errReferenceNotFound.Error():
		return Commit{}, nil
	default:
		return Commit{}, fmt.Errorf("unable to get head reference: %w", err)
	}

	o, err := r.CommitObject(h.Hash())
	if err != nil {
		return Commit{}, fmt.Errorf("unable to get head commit: %w", err)
	}

	return Commit{
		Hash: o.Hash.String(),
		Author: User{
			Name:  o.Author.Name,
			Email: o.Author.Email,
		},
		When:    o.Author.When,
		Message: o.Message,
	}, nil
}
