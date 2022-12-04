package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type User struct {
	Name  string
	Email string
}

func NewUser(r *git.Repository) (User, error) {
	cfg, err := r.Config()
	if err != nil {
		return User{}, fmt.Errorf("unable to get repository config: %w", err)
	}
	if (cfg.User.Name != "") || (cfg.User.Email != "") {
		return user(cfg), nil
	}

	cfg, err = r.ConfigScoped(config.GlobalScope)
	if err != nil {
		return User{}, fmt.Errorf("unable to get global config: %w", err)
	}

	return user(cfg), nil
}

func user(c *config.Config) User {
	return User{
		Name:  c.User.Name,
		Email: c.User.Email,
	}
}
