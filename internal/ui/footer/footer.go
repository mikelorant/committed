package footer

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type Model struct {
	Author  commit.Author
	Signoff bool
}

func New(cfg commit.Config) Model {
	return Model{
		Author: cfg.Authors[0],
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.signoff())
}

func (m *Model) ToggleSignoff() {
	m.Signoff = !m.Signoff
}

func (m Model) Value() string {
	if !m.Signoff {
		return ""
	}

	return fmt.Sprintf("Signed-off-by: %s <%s>", m.Author.Name, m.Author.Email)
}

func (m Model) signoff() string {
	str := fmt.Sprintf("Signed-off-by: %s <%s>", m.Author.Name, m.Author.Email)

	return lipgloss.NewStyle().
		Width(74).
		Height(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.HiddenBorder(), false, true).
		Padding(0, 1, 0, 1).
		Render(str)
}
