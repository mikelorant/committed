package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type SubjectModel struct {
	config SubjectConfig
	focus  bool
}

type SubjectConfig struct {
	emoji   string
	summary string
}

const (
	subjectLimit int = 50
)

func NewSubject(cfg commit.Config) SubjectModel {
	c := SubjectConfig{
		emoji:   cfg.Emoji,
		summary: cfg.Summary,
	}

	return SubjectModel{
		config: c,
	}
}

func (m SubjectModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m SubjectModel) Update(msg tea.Msg) (SubjectModel, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m SubjectModel) View() string {
	return m.render()
}

func (m SubjectModel) render() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.subjectRow())
}

func (m SubjectModel) subjectRow() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.emoji(),
		m.summary(),
		m.counter(),
	)
}

func (m SubjectModel) emoji() string {
	return lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.config.emoji)
}

func (m SubjectModel) summary() string {
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

func (m SubjectModel) counter() string {
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
