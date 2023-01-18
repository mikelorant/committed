package repository

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

var ErrLocalBranchNotFound = errors.New("local branch not found")

type Branch struct {
	Local  string
	Remote string
	Refs   []string
}

type RefsOptions struct {
	brancher     Brancher
	localBranch  string
	remoteBranch string
	headRef      *plumbing.Reference
}

type RefsResult struct {
	refs []string
}

func (r *Repository) Branch() (Branch, error) {
	c, err := r.Brancher.Config()
	if err != nil {
		return Branch{}, fmt.Errorf("unable to get repository config: %w", err)
	}

	h, err := r.Brancher.Head()
	switch {
	case err == nil:
	case err.Error() == plumbing.ErrReferenceNotFound.Error():
		return Branch{}, nil
	default:
		return Branch{}, fmt.Errorf("unable to get head reference: %w", err)
	}

	l, err := local(h)
	if err != nil {
		return Branch{}, fmt.Errorf("unable to get local branch: %w", err)
	}

	rm := remote(l, c)

	ro := RefsOptions{
		brancher:     r.Brancher,
		localBranch:  l,
		remoteBranch: rm,
		headRef:      h,
	}

	hrefs, err := headRefs(ro)
	if err != nil {
		return Branch{}, fmt.Errorf("unable to get head references: %w", err)
	}

	return Branch{
		Local:  l,
		Remote: rm,
		Refs:   hrefs,
	}, nil
}

func local(ref *plumbing.Reference) (string, error) {
	r := ref.Name().Short()

	if r == "" {
		return "", ErrLocalBranchNotFound
	}

	return r, nil
}

func remote(ref string, cfg *config.Config) string {
	bs := cfg.Branches

	if _, ok := bs[ref]; !ok {
		return ""
	}

	if bs[ref].Remote == "" {
		return ""
	}

	return fmt.Sprintf("%s/%s", bs[ref].Remote, bs[ref].Name)
}

func headRefs(ro RefsOptions) ([]string, error) {
	var rr RefsResult

	rs, err := ro.brancher.References()
	if err != nil {
		return nil, fmt.Errorf("unable to get references: %w", err)
	}

	if err := rs.ForEach(getRefFunc(ro, &rr)); err != nil {
		return nil, fmt.Errorf("unable to get references: %w", err)
	}

	return rr.refs, nil
}

func getRefFunc(ro RefsOptions, rr *RefsResult) func(*plumbing.Reference) error {
	return func(ref *plumbing.Reference) error {
		if ref.Type() != plumbing.HashReference {
			return nil
		}

		if ref.Hash() != ro.headRef.Hash() {
			return nil
		}

		name := ref.Name().Short()
		if name == ro.localBranch || name == ro.remoteBranch {
			return nil
		}

		rr.refs = append(rr.refs, name)

		return nil
	}
}
