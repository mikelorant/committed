package repository_test

import (
	"errors"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	fixtures "github.com/go-git/go-git-fixtures/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MockRepositoryWorktree struct {
	fixture *fixtures.Fixture
	err     error
}

func (m MockRepositoryWorktree) Worktree() (*git.Worktree, error) {
	dotgit := m.fixture.DotGit()
	st := filesystem.NewStorage(dotgit, cache.NewObjectLRUDefault())
	wt := memfs.New()

	repo, err := git.Open(st, wt)
	if err != nil {
		return &git.Worktree{}, err
	}

	if m.err != nil {
		return &git.Worktree{}, errMockWorktree
	}

	return repo.Worktree()
}

var errMockWorktree = errors.New("error")

func TestWorktree(t *testing.T) {
	type args struct {
		fixture *fixtures.Fixture
		err     error
	}

	type want struct {
		count int
		err   error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "dirty",
			args: args{
				fixture: fixtures.Basic().One(),
			},
			want: want{
				count: 9,
			},
		},
		{
			name: "empty",
			args: args{
				fixture: fixtures.ByTag("empty")[0],
			},
		},
		{
			name: "error",
			args: args{
				fixture: fixtures.Basic().One(),
				err:     errMockWorktree,
			},
			want: want{
				err: errMockWorktree,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r repository.Repository

			r.Worktreer = MockRepositoryWorktree{
				fixture: tt.args.fixture,
				err:     tt.args.err,
			}

			wt, err := r.Worktree()
			if tt.want.err != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, errMockWorktree)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, wt.Status, tt.want.count)
		})
	}
}
