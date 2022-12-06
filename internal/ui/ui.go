package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type MainModel struct {
	state  sessionState
	info   InfoModel
	header HeaderModel
	body   BodyModel
	footer FooterModel
	status StatusModel
	config commit.Config
	err    error
}

type sessionState int

const (
	infoView sessionState = iota
	headerView
	bodyView
	footerView
	statusView
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
		state:  infoView,
		info:   NewInfo(cfg),
		header: NewHeader(cfg),
		body:   NewBody(cfg),
		footer: NewFooter(cfg),
		status: NewStatus(cfg),
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
		m.info.Init(),
		m.header.Init(),
		m.body.Init(),
		m.footer.Init(),
		m.status.Init(),
	)
}

//nolint:ireturn
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	var cmd tea.Cmd
	var cmds []tea.Cmd

	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+a":
			m.state = infoView
		case "ctrl+b":
			m.state = bodyView
		case "ctrl+e":
			m.state = headerView
		case "ctrl+s":
			m.state = headerView
		}

		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		switch m.state {
		case infoView:
			m.info, cmd = m.info.Update(msg)
			cmds = append(cmds, cmd)
		case headerView:
			m.header, cmd = m.header.Update(msg)
			cmds = append(cmds, cmd)
		case bodyView:
			m.body, cmd = m.body.Update(msg)
			cmds = append(cmds, cmd)
		case footerView:
			m.footer, cmd = m.footer.Update(msg)
			cmds = append(cmds, cmd)
		case statusView:
			m.status, cmd = m.status.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	m.info.focus = m.state == infoView
	m.header.focus = m.state == headerView
	m.body.focus = m.state == bodyView
	m.footer.focus = m.state == footerView
	m.status.focus = m.state == statusView

	return lipgloss.JoinVertical(lipgloss.Top,
		m.info.View(),
		m.header.View(),
		m.body.View(),
		m.footer.View(),
		m.status.View(),
	)
}
