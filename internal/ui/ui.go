package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type MainModel struct {
	header  HeaderModel
	subject SubjectModel
	body    BodyModel
	footer  FooterModel
	status  StatusModel
	config  commit.Config
	err     error
}

func New(cfg commit.Config) error {
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		fh, err := tea.LogToFile(logfilePath, "committed")
		if err != nil {
			return fmt.Errorf("unable to log to file: %w", err)
		}
		defer fh.Close()
	}

	im := MainModel{
		config:  cfg,
		header:  NewHeader(cfg),
		subject: NewSubject(cfg),
		body:    NewBody(cfg),
		footer:  NewFooter(cfg),
		status:  NewStatus(cfg),
	}

	p := tea.NewProgram(im)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("unable to run program: %w", err)
	}

	return nil
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m MainModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		m.header.render(),
		m.subject.render(),
		m.body.render(),
		m.footer.render(),
		m.status.render(),
	)
}
