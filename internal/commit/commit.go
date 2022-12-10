package commit

import (
	_ "embed"
	"fmt"

	"github.com/mikelorant/committed/internal/repository"
)

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

const (
	mockHash    string = "1234567890abcdef1234567890abcdef1234567890"
	mockEmoji   string = "🐛"
	mockSummary string = "Capitalized, short (50 chars or less) summary"
)

func New() (Config, error) {
	r, err := repository.New()
	if err != nil {
		return Config{}, fmt.Errorf("unable to get repository: %w", err)
	}

	return Config{
		Hash:         mockHash,
		Name:         r.User.Name,
		Email:        r.User.Email,
		Summary:      mockSummary,
		Body:         message,
		LocalBranch:  r.Branch.Local,
		RemoteBranch: r.Branch.Remote,
		BranchRefs:   r.Branch.Refs,
		Remotes:      r.Remote.Remotes,
	}, nil
}
