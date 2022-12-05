package repository_test

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	fixtures "github.com/go-git/go-git-fixtures/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type BranchTest struct {
	t             *testing.T
	fixtures      *fixtures.Fixture
	gitRepository *git.Repository
}

func TestNewBranch(t *testing.T) {
	type args struct {
		local string
	}

	type want struct {
		local  string
		remote string
		refs   []string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "master",
			args: args{
				local: "master",
			},
			want: want{
				local:  "master",
				remote: "origin/master",
				refs:   []string{"v1.0.0"},
			},
		},
		{
			name: "branch",
			args: args{
				local: "branch",
			},
			want: want{
				local:  "branch",
				remote: "origin/branch",
				refs:   nil,
			},
		},
		{
			name: "invalid",
			args: args{
				local: "invalid",
			},
			want: want{
				local:  "master",
				remote: "origin/master",
				refs:   []string{"v1.0.0"},
			},
		},
	}

	b := BranchTest{
		t:        t,
		fixtures: fixtures.Basic().One(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b.repository()
			b.branch(tt.args.local)

			branch, err := repository.NewBranch(b.gitRepository)
			if err != nil {
				t.Errorf("unable to initialise branch: %v", tt.args.local)
			}

			assert.Equal(t, tt.want.local, branch.Local)
			assert.Equal(t, tt.want.remote, branch.Remote)
			assert.Equal(t, tt.want.refs, branch.Refs)
		})
	}
}

func (b *BranchTest) repository() {
	b.t.Helper()

	dotgit := b.fixtures.DotGit()
	st := filesystem.NewStorage(dotgit, cache.NewObjectLRUDefault())
	wt := memfs.New()

	repo, err := git.Open(st, wt)
	if err != nil {
		b.t.Errorf("unable to open repository")
	}
	b.gitRepository = repo
}

func (b *BranchTest) branch(br string) {
	b.t.Helper()

	wt, err := b.gitRepository.Worktree()
	if err != nil {
		b.t.Errorf("unable to get worktree")
	}

	co := &git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(br),
	}

	if wt.Checkout(co); err != nil {
		b.t.Errorf("unable to checkout branch: %v", b)
	}
}
