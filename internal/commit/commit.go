package commit

import (
	_ "embed"
	"fmt"

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
	Name         string
	Email        string
	Summary      string
	Body         string
	LocalBranch  string
	RemoteBranch string
	BranchRefs   []string
	Remotes      []string
}

//go:embed message.txt
var message string

var commitOptions = []string{
	"--dry-run",
}

const (
	mockHash    string = "1234567890abcdef1234567890abcdef1234567890"
	mockEmoji   string = "🐛"
	mockSummary string = "Capitalized, short (50 chars or less) summary"
)

func New() (*Commit, error) {
	r, err := repository.New()
	if err != nil {
		return nil, fmt.Errorf("unable to get repository: %w", err)
	}

	cfg := Config{
		Hash:         mockHash,
		Name:         r.User.Name,
		Email:        r.User.Email,
		Summary:      mockSummary,
		Body:         message,
		LocalBranch:  r.Branch.Local,
		RemoteBranch: r.Branch.Remote,
		BranchRefs:   r.Branch.Refs,
		Remotes:      r.Remote.Remotes,
	}

	return &Commit{
		Config: cfg,
	}, nil
}

func (c *Commit) Create() error {
	var cmd []string

	cmd = append(cmd, "git", "commit")

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

	return nil
}
