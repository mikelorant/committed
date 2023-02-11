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
	Refs   Refs
}

type Refs struct {
	Locals  []string
	Remotes []string
	Tags    []string
}

type BranchOptions struct {
	brancher     Brancher
	localBranch  string
	remoteBranch string
	headRef      *plumbing.Reference
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

	ro := BranchOptions{
		brancher:     r.Brancher,
		localBranch:  l,
		remoteBranch: rm,
		headRef:      h,
	}

	refs, err := headRefs(ro)
	if err != nil {
		return Branch{}, fmt.Errorf("unable to get head references: %w", err)
	}

	return Branch{
		Local:  l,
		Remote: rm,
		Refs:   refs,
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

func headRefs(ro BranchOptions) (Refs, error) {
	var rr Refs

	rs, err := ro.brancher.References()
	if err != nil {
		return rr, fmt.Errorf("unable to get references: %w", err)
	}

	if err := rs.ForEach(getRefFunc(ro, &rr)); err != nil {
		return rr, fmt.Errorf("unable to get references: %w", err)
	}

	return rr, nil
}

func getRefFunc(ro BranchOptions, rr *Refs) func(*plumbing.Reference) error {
	return func(ref *plumbing.Reference) error {
		refName := ref.Name().Short()

		switch {
		case isLocalRemote(ref, ro):
			return nil
		case isBranch(ref, ro):
			rr.Locals = append(rr.Locals, refName)
		case isRemote(ref, ro):
			rr.Remotes = append(rr.Remotes, refName)
		case isTag(ref, ro):
			rr.Tags = append(rr.Tags, refName)
		default:
			return nil
		}

		return nil
	}
}

func isLocalRemote(ref *plumbing.Reference, ro BranchOptions) bool {
	refName := ref.Name().Short()

	return refName == ro.localBranch || refName == ro.remoteBranch
}

func isBranch(ref *plumbing.Reference, ro BranchOptions) bool {
	refHash := ref.Strings()[1]
	headHash := ro.headRef.Strings()[1]

	return ref.Name().IsBranch() && (refHash == headHash)
}

func isRemote(ref *plumbing.Reference, ro BranchOptions) bool {
	refHash := ref.Strings()[1]
	headHash := ro.headRef.Strings()[1]

	return ref.Name().IsRemote() && (refHash == headHash)
}

func isTag(ref *plumbing.Reference, ro BranchOptions) bool {
	refHash := ref.Strings()[1]
	headHash := ro.headRef.Strings()[1]

	// Returns a tag with the given hash
	tag, err := ro.brancher.TagObject(plumbing.NewHash(refHash))
	if err != nil {
		return false
	}

	// Returns the commit pointed to by the tag
	commit, err := tag.Commit()
	if err != nil {
		return false
	}

	// Hash of the commit object.
	if commit.Hash.String() != headHash {
		return false
	}

	return true
}
