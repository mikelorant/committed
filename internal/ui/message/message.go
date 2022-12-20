package message

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/emoji"
)

type Model struct {
	emoji   emoji.Emoji
	summary string
	body    string
	footer  string
}

type Config struct {
	Emoji   emoji.Emoji
	Summary string
	Body    string
	Footer  string
}

func New(cfg Config) Model {
	return Model{
		emoji:   cfg.Emoji,
		summary: cfg.Summary,
		body:    cfg.Body,
		footer:  cfg.Footer,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	message := m.summary
	if m.emoji.ShortCode != "" {
		message = fmt.Sprintf("%s %s", m.emoji.ShortCode, m.summary)
	}

	if m.body != "" {
		b := lipgloss.NewStyle().
			MarginTop(1).
			Render(m.body)

		message = lipgloss.JoinVertical(lipgloss.Top, message, b)
	}

	if m.footer != "" {
		f := lipgloss.NewStyle().
			MarginTop(1).
			Render(m.footer)

		message = lipgloss.JoinVertical(lipgloss.Top, message, f)
	}

	return lipgloss.NewStyle().
		MarginLeft(4).
		MarginBottom(1).
		Render(message)
}
