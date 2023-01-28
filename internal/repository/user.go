package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5/config"
)

type User struct {
	Name    string
	Email   string
	Default bool
}

func (r *Repository) Users() ([]User, error) {
	var users []User

	cfg, err := r.Configer.Config()
	if err != nil {
		return users, fmt.Errorf("unable to get repository config: %w", err)
	}
	if (cfg.User.Name != "") || (cfg.User.Email != "") {
		users = append(users, user(cfg))
	}

	cfg, err = r.GlobalConfig(config.GlobalScope)
	if err != nil {
		return users, fmt.Errorf("unable to get global config: %w", err)
	}
	if (cfg.User.Name != "") || (cfg.User.Email != "") {
		users = append(users, user(cfg))
	}

	return users, nil
}

func user(c *config.Config) User {
	return User{
		Name:  c.User.Name,
		Email: c.User.Email,
	}
}
