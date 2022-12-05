package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type Remote struct {
	Remotes []string
}

func NewRemote(g *git.Repository) (Remote, error) {
	rs, err := g.Remotes()
	if err != nil {
		return Remote{}, fmt.Errorf("unable to list remotes: %w", err)
	}

	var remotes []string
	for _, r := range rs {
		remotes = append(remotes, r.Config().Name)
	}

	return Remote{
		Remotes: remotes,
	}, nil
}
