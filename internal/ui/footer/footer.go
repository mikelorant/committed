package footer

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
)

type Model struct {
	Author  repository.User
	Signoff bool

	state *commit.State
}

func New(state *commit.State) Model {
	authors := concatSlice(state.Repository.Users, state.Config.Authors)

	if len(authors) == 0 {
		authors = []repository.User{{}}
	}

	return Model{
		Author: authors[0],
		state:  state,
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
	colour := m.state.Theme.Footer()

	return lipgloss.NewStyle().
		Width(74).
		Height(1).
		MarginLeft(4).
		MarginBottom(1).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.HiddenBorder(), false, true).
		Padding(0, 1, 0, 1).
		Foreground(colour.View).
		Render(m.signoff())
}

func (m *Model) ToggleSignoff() {
	m.Signoff = !m.Signoff
}

func (m Model) Value() string {
	if !m.Signoff {
		return ""
	}

	return m.signoff()
}

func (m Model) signoff() string {
	return fmt.Sprintf("Signed-off-by: %s <%s>", m.Author.Name, m.Author.Email)
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func concatSlice[T any](first []T, second []T) []T {
	n := len(first)
	return append(first[:n:n], second...)
}
