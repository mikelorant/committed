package option

import (
	"github.com/mikelorant/committed/internal/commit"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	styles Styles
	state  *commit.State
}

type Panel int

const (
	PanelUnset Panel = iota
)

func New(state *commit.State) Model {
	m := Model{
		styles: defaultStyles(state.Theme),
		state:  state,
	}

	return m
}

//nolint:ireturn
func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
