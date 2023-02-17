package message

import (
	"bufio"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/theme"
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

	body := removeComments(m.body)
	if body != "" {
		b := m.styles.body.Render(body)
		message = lipgloss.JoinVertical(lipgloss.Top, message, b)
	}

	if m.footer != "" {
		f := m.styles.footer.Render(m.footer)
		message = lipgloss.JoinVertical(lipgloss.Top, message, f)
	}

	return m.styles.message.Render(message)
}

func removeComments(str string) string {
	var sb strings.Builder

	r := strings.NewReader(str)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		txt := scanner.Text()

		if strings.HasPrefix(txt, "#") {
			continue
		}

		fmt.Fprintln(&sb, txt)
	}

	return strings.TrimSpace(sb.String())
}
