package message

import (
	"fmt"
	"strings"

	"github.com/mikelorant/committed/internal/theme"

	tea "github.com/charmbracelet/bubbletea"
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
	var str []string

	if m.summary != "" {
		switch {
		case m.emoji != "":
			str = append(str, fmt.Sprintf("%v %v", m.emoji, m.summary))
		default:
			str = append(str, m.summary)
		}
	}

	if m.body != "" {
		str = append(str, m.body)
	}

	if m.footer != "" {
		str = append(str, m.footer)
	}

	msg := strings.Join(str, "\n\n")

	return m.styles.message.Render(msg)
}
