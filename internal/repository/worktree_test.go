package repository_test

import (
	"errors"
	"testing"

	"github.com/mikelorant/committed/internal/repository"

	"github.com/go-git/go-billy/v5/memfs"
	fixtures "github.com/go-git/go-git-fixtures/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
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

func TestIsStaged(t *testing.T) {
	type file struct {
		name       string
		statusCode git.StatusCode
	}

	tests := []struct {
		name   string
		files  []file
		staged bool
	}{
		{
			name: "unmodified",
			files: []file{
				{name: "test", statusCode: git.Unmodified},
			},
			staged: false,
		},
		{
			name: "untracked",
			files: []file{
				{name: "test", statusCode: git.Untracked},
			},
			staged: false,
		},
		{
			name: "modified",
			files: []file{
				{name: "test", statusCode: git.Modified},
			},
			staged: true,
		},
		{
			name: "added",
			files: []file{
				{name: "test", statusCode: git.Added},
			},
			staged: true,
		},
		{
			name: "deleted",
			files: []file{
				{name: "test", statusCode: git.Deleted},
			},
			staged: true,
		},
		{
			name: "renamed",
			files: []file{
				{name: "test", statusCode: git.Renamed},
			},
			staged: true,
		},
		{
			name: "copied",
			files: []file{
				{name: "test", statusCode: git.Copied},
			},
			staged: true,
		},
		{
			name: "updated_but_unmerged",
			files: []file{
				{name: "test", statusCode: git.UpdatedButUnmerged},
			},
			staged: true,
		},
		{
			name: "multiple_staged",
			files: []file{
				{name: "modified", statusCode: git.Modified},
				{name: "added", statusCode: git.Added},
			},
			staged: true,
		},
		{
			name: "multiple_unstaged",
			files: []file{
				{name: "unmodified", statusCode: git.Unmodified},
				{name: "untracked", statusCode: git.Untracked},
			},
			staged: false,
		},
		{
			name: "multiple_mixed",
			files: []file{
				{name: "modified", statusCode: git.Modified},
				{name: "unmodified", statusCode: git.Unmodified},
			},
			staged: true,
		},
		{
			name:   "empty",
			files:  []file{},
			staged: false,
		},
		{
			name:   "nil",
			files:  nil,
			staged: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files := make(map[string]*git.FileStatus, len(tt.files))

			for _, v := range tt.files {
				files[v.name] = &git.FileStatus{
					Staging: v.statusCode,
				}
			}

			wt := repository.Worktree{
				Status: files,
			}

			assert.Equal(t, tt.staged, wt.IsStaged(), tt.name)
		})
	}
}
