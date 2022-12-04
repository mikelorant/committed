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

type UserTest struct {
	t             *testing.T
	fixtures      *fixtures.Fixture
	gitRepository *git.Repository
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		name      string
		userName  string
		userEmail string
	}{
		{
			name:      "both_name_email",
			userName:  "John Doe",
			userEmail: "john.doe@example.com",
		},
		{
			name:     "only_name",
			userName: "John Doe",
		},
		{
			name:      "only_email",
			userEmail: "john.doe@example.com",
		},
	}

	u := UserTest{
		t:        t,
		fixtures: fixtures.Basic().One(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u.repository()
			u.config(tt.userName, tt.userEmail)

			user, err := repository.NewUser(u.gitRepository)
			if err != nil {
				t.Errorf("unable to initialise user")
			}

			assert.Equal(t, tt.userName, user.Name)
			assert.Equal(t, tt.userEmail, user.Email)
		})
	}
}

func (u *UserTest) repository() {
	u.t.Helper()

	dotgit := u.fixtures.DotGit()
	st := filesystem.NewStorage(dotgit, cache.NewObjectLRUDefault())
	wt := memfs.New()

	repo, err := git.Open(st, wt)
	if err != nil {
		u.t.Errorf("unable to open repository")
	}
	u.gitRepository = repo
}

func (u *UserTest) config(name, email string) {
	u.t.Helper()

	cfg, err := u.gitRepository.Config()
	if err != nil {
		u.t.Errorf("unable to get repository config")
	}

	cfg.User.Name = name
	cfg.User.Email = email

	u.gitRepository.SetConfig(cfg)
	if err != nil {
		u.t.Errorf("unable to set repository config")
	}
}
