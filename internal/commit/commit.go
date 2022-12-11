package commit

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/mikelorant/committed/internal/repository"
)

type Commit struct {
	Config  Config
	Name    string
	Email   string
	Emoji   string
	Summary string
	Body    string
	Footer  string
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

//go:embed gitcommit.gotmpl
var gitCommand string

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
	tmpl, err := template.New("commit").Parse(gitCommand)
	if err != nil {
		return fmt.Errorf("unable to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return fmt.Errorf("unable to execute template: %w", err)
	}

	out := strings.ReplaceAll(buf.String(), "\n", " ")
	log.Println(out)

	return nil
}
