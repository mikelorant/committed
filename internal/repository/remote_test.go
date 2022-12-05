package repository_test

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	fixtures "github.com/go-git/go-git-fixtures/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type RemoteTest struct {
	t             *testing.T
	fixtures      *fixtures.Fixture
	gitRepository *git.Repository
}

func TestNewRemote(t *testing.T) {
	tests := []struct {
		name    string
		remotes []string
	}{
		{
			name:    "origin",
			remotes: []string{"origin"},
		},
	}

	r := RemoteTest{
		t:        t,
		fixtures: fixtures.Basic().One(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.repository()

			remote, err := repository.NewRemote(r.gitRepository)
			if err != nil {
				t.Errorf("unable to initialise remote")
			}

			assert.Equal(t, tt.remotes, remote.Remotes)
		})
	}
}

func (r *RemoteTest) repository() {
	r.t.Helper()

	dotgit := r.fixtures.DotGit()
	st := filesystem.NewStorage(dotgit, cache.NewObjectLRUDefault())
	wt := memfs.New()

	repo, err := git.Open(st, wt)
	if err != nil {
		r.t.Errorf("unable to open repository")
	}
	r.gitRepository = repo
}
