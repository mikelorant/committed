package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type HeaderModel struct {
	config HeaderConfig
	focus  bool
}

type HeaderConfig struct {
	emoji   string
	summary string
}

const (
	subjectLimit int = 50
)

func NewHeader(cfg commit.Config) HeaderModel {
	c := HeaderConfig{
		emoji:   cfg.Emoji,
		summary: cfg.Summary,
	}

	return HeaderModel{
		config: c,
	}
}

func (m HeaderModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m HeaderModel) Update(msg tea.Msg) (HeaderModel, tea.Cmd) {
	return m, nil
}

func (m HeaderModel) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.headerRow())
}

func (m HeaderModel) headerRow() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.emoji(),
		m.summary(),
		m.counter(),
	)
}

func (m HeaderModel) emoji() string {
	return lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.config.emoji)
}

func (m HeaderModel) summary() string {
	return lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		Faint(!m.focus).
		Render(m.config.summary)
}

func (m HeaderModel) counter() string {
	i := len(m.config.summary)
	if m.config.emoji != "" {
		i += 2
	}

	c := colour(fmt.Sprintf("%d", i), white)
	t := colour(fmt.Sprintf("%d", subjectLimit), white)

	return lipgloss.NewStyle().
		Width(5).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center).
		Render(fmt.Sprintf("%s/%s", c, t))
}
