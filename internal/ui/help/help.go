package help

import (
	_ "embed"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	focus    bool
	styles   Styles
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
		styles:   defaultStyles(),
		viewport: newViewport(defaultWidth, defaultHeight),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	//nolint:gocritic
	switch msg.(type) {
	case theme.Msg:
		m.styles = defaultStyles()
		styleViewport(&m.viewport)
	}

	if m.focus {
		m.viewport, cmd = m.viewport.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	return m.styles.boundary.Render(m.viewport.View())
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

	styleViewport(&vp)

	return vp
}

func styleViewport(vp *viewport.Model) {
	vp.Style = defaultStyles().viewport
}
