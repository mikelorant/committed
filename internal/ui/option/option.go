package option

import (
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/option/section"
	"github.com/mikelorant/committed/internal/ui/option/setting"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	SectionWidth  int
	SettingWidth  int
	SectionHeight int
	SettingHeight int

	Panel Panel

	state   *commit.State
	section section.Model
	setting setting.Model
	styles  Styles
}

type Panel int

const (
	PanelSection Panel = iota
	PanelSetting
)

const (
	defaultSectionWidth  = 40
	defaultSettingWidth  = 40
	defaultSectionHeight = 20
	defaultSettingHeight = 14
)

func New(state *commit.State) Model {
	m := Model{
		SectionWidth:  defaultSectionWidth,
		SettingWidth:  defaultSettingWidth,
		SectionHeight: defaultSectionHeight,
		SettingHeight: defaultSettingHeight,
		section:       section.New(state),
		setting:       setting.New(state),
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
			case "right", "enter":
				m.setting.SelectPane(m.section.SelectedSetting())
				m.Panel = PanelSetting
			}
		}

	case PanelSetting:
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "down":
				switch p := m.setting.Selected.(type) {
				case *setting.Radio:
					p.Next()
				}

			case "up":
				switch p := m.setting.Selected.(type) {
				case *setting.Radio:
					p.Previous()
				}
			case "left", "enter":
				m.Panel = PanelSection
			case " ":
				switch p := m.setting.Selected.(type) {
				case *setting.Toggle:
					p.Enable = !p.Enable
				}
			}
		}
	}

	m.setting.SwapPaneSet(m.section.SelectedCategory())
	cmds := make([]tea.Cmd, 2)
	m.section, cmds[0] = section.ToModel(m.section.Update(msg))
	m.setting, cmds[1] = setting.ToModel(m.setting.Update(msg))

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	m.section.Width = m.SectionWidth
	m.setting.Width = m.SettingWidth

	var (
		boundarySection lipgloss.Style
		boundarySetting lipgloss.Style
	)

	switch m.Panel {
	case PanelSection:
		boundarySection = m.styles.sectionBoundaryFocus
		boundarySetting = m.styles.settingBoundary
	case PanelSetting:
		boundarySection = m.styles.sectionBoundary
		boundarySetting = m.styles.settingBoundaryFocus
	}

	section := boundarySection.
		Width(m.SectionWidth).
		Height(m.SectionHeight).
		Render(m.section.View())

	setting := boundarySetting.
		Width(m.SettingWidth).
		Height(m.SettingHeight).
		Render(m.setting.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, section, setting)
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

func (m *Model) AddPaneSet(name string, ps []setting.Paner) {
	m.setting.AddPaneSet(name, ps)
}

func (m Model) GetPaneSets() map[string][]setting.Paner {
	return m.setting.GetPaneSets()
}

func (m *Model) SelectPane(title string) {
	m.setting.SelectPane(title)
}

//nolint:ireturn
func (m *Model) ActivePane() setting.Paner {
	return m.setting.ActivePane()
}

func (m *Model) SectionIndex(c int, s int) {
	m.section.CatIndex = c
	m.section.SetIndex = s
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
