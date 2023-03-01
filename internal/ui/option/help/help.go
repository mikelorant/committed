package help

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width   int
	Height  int
	Content string
	Styles  Styles
}

type Styles struct{}

const (
	defaultWidth  = 20
	defaultHeight = 5
)

func New() Model {
	return Model{
		Width:  defaultWidth,
		Height: defaultHeight,
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
	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		MaxWidth(m.Width).
		MaxHeight(m.Height).
		Render(m.renderHelp())
}

func (m *Model) SetContent(content string) {
	m.Content = content
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func (m Model) renderHelp() string {
	return m.Content
}
