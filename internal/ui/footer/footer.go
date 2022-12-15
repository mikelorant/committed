package footer

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type Model struct {
	Name  string
	Email string
}

func New(cfg commit.Config) Model {
	return Model{
		Name:  cfg.Authors[0].Name,
		Email: cfg.Authors[0].Email,
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

func (m Model) signoff() string {
	str := fmt.Sprintf("Signed-off-by: %s <%s>", m.Name, m.Email)

	return lipgloss.NewStyle().
		Width(74).
		Height(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.HiddenBorder(), false, true).
		Padding(0, 1, 0, 1).
		Render(str)
}
