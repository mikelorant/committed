package commit

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/creack/pty"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
	"gopkg.in/alessio/shellescape.v1"
)

type Commit struct {
	Config  Config
	Name    string
	Email   string
	Emoji   string
	Summary string
	Body    string
	Footer  string
	cmd     []string
}

type Config struct {
	Hash         string
	Authors      []Author
	Summary      string
	Body         string
	LocalBranch  string
	RemoteBranch string
	BranchRefs   []string
	Remotes      []string
	Emojis       []emoji.Emoji
}

type Author struct {
	Name  string
	Email string
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

func New() (*Commit, error) {
	r, err := repository.New()
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	e, err := emoji.New()
	if err != nil {
		return nil, fmt.Errorf("unable to get emojis: %w", err)
	}

	defaultAuthor := Author{
		Name:  r.User.Name,
		Email: r.User.Email,
	}

	var authors []Author
	authors = append(authors, defaultAuthor)

	cfg := Config{
		Hash:         mockHash,
		Authors:      authors,
		Summary:      summary,
		Body:         message,
		LocalBranch:  r.Branch.Local,
		RemoteBranch: r.Branch.Remote,
		BranchRefs:   r.Branch.Refs,
		Remotes:      r.Remote.Remotes,
		Emojis:       e,
	}

	return &Commit{
		Config: cfg,
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

	if c.Name != "" && c.Email != "" {
		author := fmt.Sprintf("%s <%s>", c.Name, c.Email)
		cmd = append(cmd, "--author", shellescape.Quote(author))
	}

	var subject string
	if c.Emoji != "" {
		subject = fmt.Sprintf("%s %s", c.Emoji, c.Summary)
	} else {
		subject = c.Summary
	}
	cmd = append(cmd, "--message", shellescape.Quote(subject))

	if c.Body != "" {
		cmd = append(cmd, "--message", shellescape.Quote(c.Body))
	}

	cmd = append(cmd, commitOptions...)

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
	_, err = io.Copy(&buf, fh)
	if err != nil {
		return fmt.Errorf("unable to copy commit output: %w", err)
	}

	out := buf.String()

	fmt.Println()
	fmt.Println(strings.TrimSpace(string(out)))

	return nil
}
