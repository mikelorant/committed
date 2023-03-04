package option

import (
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/colour"
	"github.com/mikelorant/committed/internal/ui/option/help"
	"github.com/mikelorant/committed/internal/ui/option/section"
	"github.com/mikelorant/committed/internal/ui/option/setting"
	"github.com/mikelorant/committed/internal/ui/option/theme"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	SectionWidth  int
	SettingWidth  int
	HelpWidth     int
	ThemeWidth    int
	SectionHeight int
	SettingHeight int
	HelpHeight    int
	ThemeHeight   int

	Panel Panel

	state   *commit.State
	focus   bool
	section section.Model
	setting setting.Model
	help    help.Model
	theme   theme.Model
	styles  Styles
}

type Panel int

const (
	PanelSection Panel = iota
	PanelSetting
	PanelHelp
	PanelTheme
)

const (
	defaultSectionWidth  = 40
	defaultSettingWidth  = 40
	defaultHelpWidth     = 40
	defaultThemeWidth    = 40
	defaultSectionHeight = 20
	defaultSettingHeight = 14
	defaultHelpHeight    = 3
	defaultThemeHeight   = 20
)

func New(state *commit.State) Model {
	m := Model{
		SectionWidth:  defaultSectionWidth,
		SettingWidth:  defaultSettingWidth,
		HelpWidth:     defaultHelpWidth,
		ThemeWidth:    defaultThemeWidth,
		SectionHeight: defaultSectionHeight,
		SettingHeight: defaultSettingHeight,
		HelpHeight:    defaultHelpHeight,
		ThemeHeight:   defaultThemeHeight,
		section:       section.New(state),
		setting:       setting.New(state),
		theme:         theme.New(state),
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
			case "right", "enter", "tab":
				if m.Category() == "Theme" {
					m.Panel = PanelTheme
					break
				}
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
			case "left", "enter", "tab":
				m.Panel = PanelSection
			case " ":
				switch p := m.setting.Selected.(type) {
				case *setting.Toggle:
					p.Enable = !p.Enable
				}
			}
		}

	case PanelTheme:
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.theme, _ = theme.ToModel(m.theme.Update(msg))
				fallthrough
			case "left", "tab":
				m.theme.Blur()
				m.Panel = PanelSection
			}
		}
	}

	//nolint:gocritic
	switch msg.(type) {
	case colour.Msg:
		m.styles = defaultStyles(m.state.Theme)
	}

	if m.Panel == PanelTheme && !m.theme.Focused() {
		m.theme.Focus()
		msg = nil
	}

	m.setting.SwapPaneSet(m.section.SelectedCategory())

	cmds := make([]tea.Cmd, 3)
	m.section, cmds[0] = section.ToModel(m.section.Update(msg))
	m.setting, cmds[1] = setting.ToModel(m.setting.Update(msg))
	m.theme, cmds[2] = theme.ToModel(m.theme.Update(msg))

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	m.section.Width = m.SectionWidth
	m.setting.Width = m.SettingWidth
	m.help.Width = m.HelpWidth
	m.theme.Width = m.ThemeWidth
	m.theme.Height = m.ThemeHeight

	var (
		boundarySection lipgloss.Style
		boundarySetting lipgloss.Style
		boundaryTheme   lipgloss.Style
	)

	if !m.state.Config.View.HighlightActive {
		m.styles.sectionBoundary = m.styles.sectionBoundaryFocus
		m.styles.settingBoundary = m.styles.settingBoundaryFocus
		m.styles.helpBoundary = m.styles.helpBoundaryFocus
	}

	switch m.Panel {
	case PanelSection:
		boundarySection = m.styles.sectionBoundaryFocus
		boundarySetting = m.styles.settingBoundary
		boundaryTheme = m.styles.themeBoundary
	case PanelSetting:
		boundarySection = m.styles.sectionBoundary
		boundarySetting = m.styles.settingBoundaryFocus
		boundaryTheme = m.styles.themeBoundary
	case PanelTheme:
		boundarySection = m.styles.sectionBoundary
		boundarySetting = m.styles.settingBoundary
		boundaryTheme = m.styles.themeBoundary
	}

	section := boundarySection.
		Width(m.SectionWidth).
		Height(m.SectionHeight).
		Render(m.section.View())

	setting := boundarySetting.
		Width(m.SettingWidth).
		Height(m.SettingHeight).
		Render(m.setting.View())

	help := m.styles.helpBoundary.
		Width(m.HelpWidth).
		Height(m.HelpHeight).
		Render(m.help.View())

	theme := boundaryTheme.
		Width(m.ThemeWidth).
		Height(m.ThemeHeight).
		Render(m.theme.View())

	if m.Category() == "Theme" {
		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			section,
			theme,
		)
	}

	settingHelp := lipgloss.JoinVertical(
		lipgloss.Top,
		setting,
		help,
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, section, settingHelp)
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
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

func (m *Model) SetHelp(content string) {
	m.help.Content = content
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
