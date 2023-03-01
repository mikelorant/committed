package option

import (
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/option/section"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	SectionWidth  int
	SectionHeight int

	Panel Panel

	state   *commit.State
	section section.Model
	styles  Styles
}

type Panel int

const (
	PanelSection Panel = iota
)

const (
	defaultSectionWidth  = 40
	defaultSectionHeight = 20
)

func New(state *commit.State) Model {
	m := Model{
		SectionWidth:  defaultSectionWidth,
		SectionHeight: defaultSectionHeight,
		section:       section.New(state),
		styles:        defaultStyles(state.Theme),
		state:         state,
	}

	return m
}

//nolint:ireturn
func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.Panel {
	case PanelSection:
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "down":
				m.section.Next()
			case "up":
				m.section.Previous()
			}
		}
	}

	cmds := make([]tea.Cmd, 1)
	m.section, cmds[0] = section.ToModel(m.section.Update(msg))

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	m.section.Width = m.SectionWidth

	var boundarySection lipgloss.Style

	switch m.Panel {
	case PanelSection:
		boundarySection = m.styles.sectionBoundaryFocus
	}

	return boundarySection.
		Width(m.SectionWidth).
		Height(m.SectionHeight).
		Render(m.section.View())
}

func (m *Model) SetSettings(set []section.Setting) {
	m.section.Settings = set
}

func (m *Model) Category() string {
	return m.section.SelectedCategory()
}

func (m *Model) Setting() string {
	return m.section.SelectedSetting()
}

func (m *Model) SectionIndex(c int, s int) {
	m.section.CatIndex = c
	m.section.SetIndex = s
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
