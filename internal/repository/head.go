package repository

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
)

type Head struct {
	Hash    string
	Author  User
	When    time.Time
	Message string
}

func (r *Repository) Head() (Head, error) {
	h, err := r.Header.Head()

	switch {
	case err == nil:
	case err.Error() == plumbing.ErrReferenceNotFound.Error():
		return Head{}, nil
	default:
		return Head{}, fmt.Errorf("unable to get head reference: %w", err)
	}

	o, err := r.Header.CommitObject(h.Hash())
	if err != nil {
		return Head{}, fmt.Errorf("unable to get head commit: %w", err)
	}

	return Head{
		Hash: o.Hash.String(),
		Author: User{
			Name:  o.Author.Name,
			Email: o.Author.Email,
		},
		When:    o.Author.When,
		Message: o.Message,
	}, nil
}
