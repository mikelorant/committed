package repository

import (
	"fmt"
)

func (r *Repository) Remotes() ([]string, error) {
	var remotes []string

	rs, err := r.Remoter.Remotes()
	if err != nil {
		return remotes, fmt.Errorf("unable to list remotes: %w", err)
	}

	for _, r := range rs {
		remotes = append(remotes, r.Config().Name)
	}

	return remotes, nil
}
