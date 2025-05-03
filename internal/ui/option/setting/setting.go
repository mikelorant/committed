package setting

import (
	"strings"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/colour"

	tea "github.com/charmbracelet/bubbletea"
)

type Paner interface {
	Type() Type
	Focus()
	Blur()
	Render(Styles) string
}

type Model struct {
	Width    int
	Height   int
	Selected Paner

	paneSet map[string][]Paner
	panes   []Paner
	styles  Styles
	state   *commit.State
}

type Type int

const (
	TypeUnset = iota
	TypeNoop
	TypeRadio
	TypeToggle
)

const (
	defaultWidth  = 20
	defaultHeight = 20
)

func New(state *commit.State) Model {
	return Model{
		Width:   defaultWidth,
		Height:  defaultHeight,
		paneSet: make(map[string][]Paner),
		styles:  defaultStyles(state.Theme),
		state:   state,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg.(type) {
	case colour.Msg:
		m.styles = defaultStyles(m.state.Theme)
	}

	return m, nil
}

func (m Model) View() string {
	return m.RenderPanes()
}

func (m Model) RenderPanes() string {
	var str []string

	for _, p := range m.panes {
		str = append(str, p.Render(m.styles))
	}

	return strings.Join(str, "\n\n")
}

func (m *Model) AddPaneSet(name string, ps []Paner) {
	m.paneSet[name] = ps
	m.panes = m.paneSet[name]
	m.Selected = m.panes[0]
}

func (m *Model) SelectPane(title string) {
	for _, p := range m.panes {
		switch pane := p.(type) {
		case *Radio:
			if pane.Title != title {
				continue
			}
		case *Toggle:
			if pane.Title != title {
				continue
			}
		}

		m.Selected.Blur()
		m.Selected = p
		m.Selected.Focus()

		return
	}
}

//nolint:ireturn
func (m *Model) ActivePane() Paner {
	return m.Selected
}

func (m *Model) SwapPaneSet(name string) {
	m.panes = m.paneSet[name]
}

func (m Model) GetPaneSets() map[string][]Paner {
	return m.paneSet
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
