package repository

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

type Commit struct {
	Author  string
	Subject string
	Body    string
	Footer  string
	Amend   bool
	DryRun  bool

	cmd []string
}

type CommitOptions func(c *Commit)

const commitCommand = "git"

func Apply(c Commit, opts ...CommitOptions) error {
	for _, o := range opts {
		o(&c)
	}

	c.build()

	var buf bytes.Buffer

	if err := c.exec(&buf); err != nil {
		return fmt.Errorf("unable to apply commit: %w", err)
	}

	output(os.Stdout, &buf)

	return nil
}

func WithAmend(b bool) func(c *Commit) {
	return func(c *Commit) {
		c.Amend = b
	}
}

func WithDryRun(b bool) func(c *Commit) {
	return func(c *Commit) {
		c.DryRun = b
	}
}

func (c *Commit) build() {
	var cmd []string

	cmd = append(cmd, "commit")
	cmd = append(cmd, "--author", c.Author)
	cmd = append(cmd, "--message", c.Subject)

	if c.Body != "" {
		cmd = append(cmd, "--message", c.Body)
	}

	if c.Footer != "" {
		cmd = append(cmd, "--message", c.Footer)
	}

	if c.DryRun {
		cmd = append(cmd, "--dry-run")
	}

	if c.Amend {
		cmd = append(cmd, "--amend")
	}

	c.cmd = cmd
}

func (c *Commit) exec(buf *bytes.Buffer) error {
	cmd := exec.Command(commitCommand, c.cmd...)
	fh, err := pty.Start(cmd)
	if err != nil {
		return fmt.Errorf("unable to exec commit command: %w", err)
	}
	defer fh.Close()

	if _, err = io.Copy(buf, fh); err != nil {
		var pathError *fs.PathError
		if !errors.As(err, &pathError) {
			return fmt.Errorf("unable to copy commit output: %w", err)
		}
		if pathError.Path != "/dev/ptmx" {
			return fmt.Errorf("unable to copy commit output: %w", err)
		}
	}

	return nil
}

func output(w io.Writer, r io.Reader) error {
	fmt.Fprintln(w)
	if _, err := io.Copy(w, r); err != nil {
		return fmt.Errorf("unable to copy output: %w", err)
	}

	return nil
}
