package commit

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os/exec"

	"github.com/creack/pty"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

type Commit struct {
	Config  Config
	Author  Author
	Emoji   string
	Summary string
	Body    string
	Footer  string

	options Options
	cmd     []string
}

type Config struct {
	Authors      []Author
	Placeholders Placeholders
	LocalBranch  string
	RemoteBranch string
	BranchRefs   []string
	Remotes      []string
	Emojis       []emoji.Emoji
}

type Options struct {
	Apply bool
}

type Author struct {
	Name  string
	Email string
}

type Placeholders struct {
	Hash    string
	Summary string
	Body    string
}

//go:embed message.txt
var message string

var (
	commitCommand = "git"
	commitOptions = []string{
		"--dry-run",
	}
)

var exitError *exec.ExitError

const (
	mockHash string = "1234567890abcdef1234567890abcdef1234567890"
	summary  string = "Capitalized, short (50 chars or less) summary"
)

func New(opts Options) (*Commit, error) {
	r, err := repository.New()
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	e, err := emoji.New()
	if err != nil {
		return nil, fmt.Errorf("unable to get emojis: %w", err)
	}

	var authors []Author
	for _, user := range r.Users {
		a := Author{
			Name:  user.Name,
			Email: user.Email,
		}

		authors = append(authors, a)
	}

	placeholders := Placeholders{
		Hash:    mockHash,
		Summary: summary,
		Body:    message,
	}

	cfg := Config{
		Authors:      authors,
		Placeholders: placeholders,
		LocalBranch:  r.Branch.Local,
		RemoteBranch: r.Branch.Remote,
		BranchRefs:   r.Branch.Refs,
		Remotes:      r.Remote.Remotes,
		Emojis:       e,
	}

	return &Commit{
		Config:  cfg,
		options: opts,
	}, nil
}

func (c *Commit) Create() error {
	c.build()
	if err := c.exec(); err != nil {
		return fmt.Errorf("unable to commit: %w", err)
	}

	return nil
}

func (c *Commit) build() {
	var cmd []string

	cmd = append(cmd, "commit")

	if c.Author.Name != "" && c.Author.Email != "" {
		author := fmt.Sprintf("%s <%s>", c.Author.Name, c.Author.Email)
		cmd = append(cmd, "--author", author)
	}

	var subject string
	if c.Emoji != "" {
		subject = fmt.Sprintf("%s %s", c.Emoji, c.Summary)
	} else {
		subject = c.Summary
	}
	cmd = append(cmd, "--message", subject)

	if c.Body != "" {
		cmd = append(cmd, "--message", c.Body)
	}

	if c.Footer != "" {
		cmd = append(cmd, "--message", c.Footer)
	}

	if !c.options.Apply {
		cmd = append(cmd, commitOptions...)
	}

	c.cmd = cmd
}

func (c *Commit) exec() error {
	cmd := exec.Command(commitCommand, c.cmd...)
	fh, err := pty.Start(cmd)
	if err != nil {
		return fmt.Errorf("unable to exec commit command: %w", err)
	}
	defer fh.Close()

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, fh); err != nil {
		var pathError *fs.PathError
		if !errors.As(err, &pathError) {
			return fmt.Errorf("unable to copy commit output: %w", err)
		}
		if pathError.Path != "/dev/ptmx" {
			return fmt.Errorf("unable to copy commit output: %w", err)
		}
	}

	out := buf.String()

	fmt.Println()
	fmt.Println(string(out))

	return nil
}
