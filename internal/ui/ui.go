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
		state: headerView,
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
	var cmd tea.Cmd
	var cmds []tea.Cmd

	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "alt+1":
			if m.state == infoView {
				return m, nil
			}
			m.state = infoView
		case "alt+2":
			fallthrough
		case "alt+3":
			if m.state == headerView {
				return m, nil
			}
			m.state = headerView
		case "alt+4":
			if m.state == bodyView {
				return m, nil
			}
			m.state = bodyView
		case "enter":
			if m.state == headerView {
				m.state = bodyView
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.models.info.focus = false
	m.models.header.focus = false
	m.models.body.focus = false
	m.models.footer.focus = false
	m.models.status.focus = false

	switch m.state {
	case infoView:
		m.models.info.focus = true
	case headerView:
		m.models.header.focus = true
	case bodyView:
		m.models.body.focus = true
	case footerView:
		m.models.footer.focus = true
	case statusView:
		m.models.status.focus = true
	}

	m.models.info, cmd = m.models.info.Update(msg)
	cmds = append(cmds, cmd)

	m.models.header, cmd = m.models.header.Update(msg)
	cmds = append(cmds, cmd)

	m.models.body, cmd = m.models.body.Update(msg)
	cmds = append(cmds, cmd)

	m.models.footer, cmd = m.models.footer.Update(msg)
	cmds = append(cmds, cmd)

	m.models.status, cmd = m.models.status.Update(msg)
	cmds = append(cmds, cmd)

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
