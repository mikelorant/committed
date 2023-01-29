package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type Worktree struct {
	Status git.Status
}

func (r *Repository) Worktree() (Worktree, error) {
	var wt Worktree

	w, err := r.Worktreer.Worktree()
	if err != nil {
		return Worktree{}, fmt.Errorf("unable to get worktree: %w", err)
	}

	s, err := w.Status()
	if err != nil {
		return Worktree{}, fmt.Errorf("unable to get status of worktree: %w", err)
	}
	wt.Status = s

	return wt, nil
}
