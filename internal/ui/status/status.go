package status

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Next     string
	Previous string

	Shortcuts Shortcuts
	styles    Styles
}

func New() Model {
	return Model{
		Shortcuts: newShortcuts(),
		styles:    defaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg.(type) {
	case theme.Msg:
		m.styles = defaultStyles()
	}

	m.Shortcuts.Shortcuts = defaultShortcuts()
	m.next()
	m.previous()
	m.Shortcuts.view = m.Shortcuts.render()

	return m, nil
}

func (m Model) View() string {
	return m.Shortcuts.view
}

func (m *Model) next() {
	if m.Next == "" {
		return
	}

	next := Shortcut{
		Modifier: NoModifier,
		Key:      "tab",
		Label:    m.Next,
	}

	m.Shortcuts.Shortcuts = append(m.Shortcuts.Shortcuts, next)
}

func (m *Model) previous() {
	if m.Previous == "" {
		return
	}

	previous := Shortcut{
		Modifier: ShiftModifier,
		Key:      "tab",
		Label:    m.Previous,
	}

	m.Shortcuts.Shortcuts = append(m.Shortcuts.Shortcuts, previous)
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
