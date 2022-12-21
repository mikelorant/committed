package help

import (
	_ "embed"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focus    bool
	viewport viewport.Model
}

//go:embed help.txt
var Content string

const (
	defaultWidth  = 72
	defaultHeight = 23
)

func New() Model {
	return Model{
		viewport: newViewport(defaultWidth, defaultHeight),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.focus {
		m.viewport, cmd = m.viewport.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		Width(74).
		MarginBottom(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		Render(m.viewport.View())
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m Model) Focused() bool {
	return m.focus
}

func newViewport(w, h int) viewport.Model {
	vp := viewport.New(w, h)
	vp.SetContent(Content)

	return vp
}
