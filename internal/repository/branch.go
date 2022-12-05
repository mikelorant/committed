package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

type Branch struct {
	repository *git.Repository
	config     *config.Config
	headRef    *plumbing.Reference
	Local      string
	Remote     string
}

func NewBranch(r *git.Repository) (Branch, error) {
	c, err := r.Config()
	if err != nil {
		return Branch{}, fmt.Errorf("unable to get repository config: %w", err)
	}

	h, err := r.Head()
	if err != nil {
		return Branch{}, fmt.Errorf("unable to get head reference: %w", err)
	}

	branch := Branch{
		repository: r,
		config:     c,
		headRef:    h,
	}

	local, err := branch.local()
	if err != nil {
		return branch, fmt.Errorf("unable to get local branch: %w", err)
	}
	branch.Local = local

	remote, err := branch.remote()
	if err != nil {
		return branch, fmt.Errorf("unable to get remote branch: %w", err)
	}
	branch.Remote = remote

	return branch, nil
}

func (b *Branch) local() (string, error) {
	return b.headRef.Name().Short(), nil
}

func (b *Branch) remote() (string, error) {
	l := b.Local
	bs := b.config.Branches

	if l == "" {
		return "", nil
	}

	if _, ok := bs[l]; !ok {
		return "", nil
	}

	return fmt.Sprintf("%s/%s", bs[l].Remote, bs[l].Name), nil
}
