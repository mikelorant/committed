package repository

import (
	"fmt"
	"os"
)

type Commit struct {
	Author  string
	Subject string
	Body    string
	Footer  string
	Amend   bool
	DryRun  bool
}

const command = "git"

func (r *Repository) Apply(c Commit) error {
	if err := r.Runner(os.Stdout, command, build(c)); err != nil {
		return fmt.Errorf("unable to run command: %w", err)
	}

	return nil
}

func build(c Commit) []string {
	var args []string

	args = append(args, "commit")
	args = append(args, "--author", c.Author)
	args = append(args, "--message", c.Subject)

	if c.Body != "" {
		args = append(args, "--message", c.Body)
	}

	if c.Footer != "" {
		args = append(args, "--message", c.Footer)
	}

	if c.DryRun {
		args = append(args, "--dry-run")
	}

	if c.Amend {
		args = append(args, "--amend")
	}

	return args
}
