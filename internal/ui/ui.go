package ui

import (
	_ "embed"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/repository"
)

type model struct {
	commit       string
	name         string
	email        string
	emoji        string
	summary      string
	body         string
	localBranch  string
	remoteBranch string
	branchRefs   []string
	remotes      []string
	err          error
}

//go:embed message.txt
var message string

const (
	mockCommit  string = "1234567890abcdef1234567890abcdef1234567890"
	mockEmoji   string = "🐛"
	mockSummary string = "Capitalized, short (50 chars or less) summary"
)

func New() error {
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		fh, err := tea.LogToFile(logfilePath, "committed")
		if err != nil {
			return fmt.Errorf("unable to log to file: %w", err)
		}
		defer fh.Close()
	}

	im, err := initialModel()
	if err != nil {
		return fmt.Errorf("unable to build initial model: %w", err)
	}

	p := tea.NewProgram(im)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("unable to run program: %w", err)
	}

	return nil
}

func (m model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	return commit(m)
}

func initialModel() (model, error) {
	r, err := repository.New()
	if err != nil {
		return model{}, fmt.Errorf("unable to get repository: %w", err)
	}

	return model{
		commit:       mockCommit,
		name:         r.User.Name,
		email:        r.User.Email,
		emoji:        mockEmoji,
		summary:      mockSummary,
		body:         message,
		localBranch:  r.Branch.Local,
		remoteBranch: r.Branch.Remote,
		branchRefs:   r.Branch.Refs,
		remotes:      r.Remote.Remotes,
	}, nil
}
