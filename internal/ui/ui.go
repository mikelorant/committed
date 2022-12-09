package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type MainModel struct {
	state  component
	models Models
	config commit.Config
	err    error
}

type Models struct {
	info   InfoModel
	header HeaderModel
	body   BodyModel
	footer FooterModel
	status StatusModel
}

type State struct {
	component component
	display   display
}

type (
	component int
	display   int
)

const (
	emptyComponent component = iota
	authorComponent
	emojiComponent
	summaryComponent
	bodyComponent
)

const (
	defaultDisplay display = iota
	compactDisplay
	expandedDisplay
)

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
		state: summaryComponent,
		models: Models{
			info:   NewInfo(cfg),
			header: NewHeader(cfg),
			body:   NewBody(cfg),
			footer: NewFooter(cfg),
			status: NewStatus(cfg),
		},
		config: cfg,
	}

	p := tea.NewProgram(im)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("unable to run program: %w", err)
	}

	return nil
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(
		m.models.info.Init(),
		m.models.header.Init(),
		m.models.body.Init(),
		m.models.footer.Init(),
		m.models.status.Init(),
	)
}

//nolint:ireturn
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "alt+1":
			if m.state == authorComponent {
				return m, nil
			}
			m.state = authorComponent
		case "alt+2":
			if m.state == emojiComponent {
				return m, nil
			}
			m.state = emojiComponent
		case "alt+3":
			if m.state == summaryComponent {
				return m, nil
			}
			m.state = summaryComponent
		case "alt+4":
			if m.state == bodyComponent {
				return m, nil
			}
			m.state = bodyComponent
		case "enter":
			if m.state == summaryComponent {
				m.state = bodyComponent
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.models.info.state = State{}
	m.models.header.state = State{}
	m.models.body.state = State{}
	m.models.footer.state = State{}
	m.models.status.state = State{}

	switch m.state {
	case authorComponent:
		m.models.header.state.component = authorComponent
	case emojiComponent:
		m.models.header.state.component = emojiComponent
		m.models.header.state.display = expandedDisplay
		m.models.body.state.display = compactDisplay
	case summaryComponent:
		m.models.header.state.component = summaryComponent
	case bodyComponent:
		m.models.body.state.component = bodyComponent
	}

	cmds := make([]tea.Cmd, 5)
	m.models.info, cmds[0] = m.models.info.Update(msg)
	m.models.header, cmds[1] = m.models.header.Update(msg)
	m.models.body, cmds[2] = m.models.body.Update(msg)
	m.models.footer, cmds[3] = m.models.footer.Update(msg)
	m.models.status, cmds[4] = m.models.status.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		m.models.info.View(),
		m.models.header.View(),
		m.models.body.View(),
		m.models.footer.View(),
		m.models.status.View(),
	)
}
