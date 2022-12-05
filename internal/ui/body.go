package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type BodyModel struct {
	body  string
	focus bool
}

func NewBody(cfg commit.Config) BodyModel {
	return BodyModel{
		body: cfg.Body,
	}
}

func (m BodyModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m BodyModel) Update(msg tea.Msg) (BodyModel, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m BodyModel) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(body(m.body, m.focus))
}

func body(str string, focus bool) string {
	return lipgloss.NewStyle().
		Width(74).
		Height(19).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		Faint(!focus).
		Render(strings.TrimSpace(str))
}
