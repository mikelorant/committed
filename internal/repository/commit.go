package repository

import (
	"fmt"
	"io"
	"os"
)

type Commit struct {
	Author      string
	Subject     string
	Body        string
	Footer      string
	Amend       bool
	DryRun      bool
	Hook        bool
	MessageFile string
}

const command = "git"

func (r *Repository) Apply(c Commit) error {
	if c.Hook {
		return r.hook(c)
	}

	if err := r.Runner(os.Stdout, command, build(c)); err != nil {
		return fmt.Errorf("unable to run command: %w", err)
	}

	return nil
}

func (r *Repository) hook(c Commit) error {
	fh, err := r.OpenFiler(c.MessageFile, os.O_RDWR|os.O_TRUNC, 0o0755)
	if err != nil {
		return fmt.Errorf("unble to open file: %w", err)
	}

	err = write(c, fh)
	if err != nil {
		return fmt.Errorf("unable to write file: %w", err)
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

func write(c Commit, w io.WriteCloser) error {
	var err error

	fmt.Fprintln(w, c.Subject)
	fmt.Fprintln(w, "")

	if c.Body != "" {
		fmt.Fprintln(w, c.Body)
		fmt.Fprintln(w, "")
	}

	if c.Footer != "" {
		fmt.Fprintln(w, c.Footer)
	}

	if err = w.Close(); err != nil {
		return fmt.Errorf("unable to close file: %w", err)
	}

	return nil
}
