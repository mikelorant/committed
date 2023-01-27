package message

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	emoji   string
	summary string
	body    string
	footer  string
	styles  Styles
}

type State struct {
	Emoji   string
	Summary string
	Body    string
	Footer  string
	Theme   theme.Theme
}

func New(state State) Model {
	return Model{
		emoji:   state.Emoji,
		summary: state.Summary,
		body:    state.Body,
		footer:  state.Footer,
		styles:  defaultStyles(state.Theme),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	message := m.styles.summary.Render(m.summary)
	if m.emoji != "" {
		s := fmt.Sprintf("%s %s", m.emoji, m.summary)
		message = m.styles.summary.Render(s)
	}

	if m.body != "" {
		b := m.styles.body.Render(m.body)
		message = lipgloss.JoinVertical(lipgloss.Top, message, b)
	}

	if m.footer != "" {
		f := m.styles.footer.Render(m.footer)
		message = lipgloss.JoinVertical(lipgloss.Top, message, f)
	}

	return m.styles.message.Render(message)
}
