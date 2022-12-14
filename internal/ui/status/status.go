package status

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Next     string
	Previous string

	shortcuts shortcuts
	styles    Styles
}

func New() Model {
	return Model{
		shortcuts: newShortcuts(),
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

	m.shortcuts.shortcuts = defaultShortcuts()
	m.next()
	m.previous()
	m.shortcuts.view = m.shortcuts.render()

	return m, nil
}

func (m Model) View() string {
	return m.shortcuts.view
}

func (m *Model) next() {
	if m.Next == "" {
		return
	}

	next := Shortcut{
		Modifier: noModifier,
		Key:      "tab",
		Label:    m.Next,
	}

	m.shortcuts.shortcuts = append(m.shortcuts.shortcuts, next)
}

func (m *Model) previous() {
	if m.Previous == "" {
		return
	}

	previous := Shortcut{
		Modifier: shiftModifier,
		Key:      "tab",
		Label:    m.Previous,
	}

	m.shortcuts.shortcuts = append(m.shortcuts.shortcuts, previous)
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
