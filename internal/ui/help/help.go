package help

import (
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	focus    bool
	state    *commit.State
	styles   Styles
	viewport viewport.Model
}

const (
	defaultWidth  = 72
	defaultHeight = 23
)

func New(state *commit.State) Model {
	return Model{
		state:    state,
		styles:   defaultStyles(state.Theme),
		viewport: newViewport(defaultWidth, defaultHeight, state),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	//nolint:gocritic
	switch msg.(type) {
	case colour.Msg:
		m.styles = defaultStyles(m.state.Theme)
		styleViewport(&m.viewport, m.state)
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

func (m *Model) SetContent(str string) {
	m.viewport.SetContent(str)
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func newViewport(w, h int, state *commit.State) viewport.Model {
	vp := viewport.New(w, h)
	vp.SetContent(state.Placeholders.Help)

	styleViewport(&vp, state)

	return vp
}

func styleViewport(vp *viewport.Model, state *commit.State) {
	vp.Style = defaultStyles(state.Theme).viewport
}
