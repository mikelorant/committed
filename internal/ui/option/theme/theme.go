package theme

import (
	"fmt"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/fuzzy"
	"github.com/mikelorant/committed/internal/ui/colour"
	"github.com/mikelorant/committed/internal/ui/filterlist"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width  int
	Height int
	theme  string

	filterList filterlist.Model
	styles     Styles
	focus      bool
	state      *commit.State
}

const (
	defaultFilterPromptText = "Select a theme:"
	defaultFilterHeight     = 18
	defaultFilterWidth      = 34
)

func New(state *commit.State) Model {
	m := Model{
		styles:     defaultStyles(state.Theme),
		state:      state,
		filterList: filterlist.New(state),
	}

	m.filterList.SetItems(castToListItems(state.Theme.List()))
	m.filterList.SetHeight(defaultFilterHeight)
	m.filterList.SetWidth(defaultFilterWidth)
	m.filterList.SetPromptText(defaultFilterPromptText)
	m.filterList.Border = false

	var index int
	for idx, id := range m.state.Theme.ListID() {
		if id == m.state.Theme.ID {
			m.theme = m.state.Theme.Registry.DisplayName()
			index = idx
			break
		}
	}

	m.filterList.Select(index)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	var refresh bool

	if m.focus {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.setTheme()
				fallthrough
			case "up", "down", "pgdown", "pgup":
				refresh = true
			}
		}
	}

	//nolint:gocritic
	switch msg.(type) {
	case colour.Msg:
		m.styles = defaultStyles(m.state.Theme)
	}

	switch {
	case !m.focus && m.filterList.Focused():
		m.filterList.Blur()
	case m.focus && !m.filterList.Focused():
		m.filterList.Focus()
		fallthrough
	case m.focus:
		ranks := fuzzy.Rank(m.filterList.Filter(), castToFuzzyItems(m.state.Theme.List()))

		items := make([]list.Item, len(ranks))
		for i, rank := range ranks {
			items[i] = castToListItems(m.state.Theme.List())[rank]
		}
		m.filterList.SetItems(items)
	}

	m.filterList, cmd = filterlist.ToModel(m.filterList.Update(msg))
	cmds = append(cmds, cmd)

	if refresh {
		m.state.Theme.Set(m.filterList.SelectedItem().(listItem).Description())
		cmds = append(cmds, colour.Update)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.themeTitle(),
		m.themeList(),
	)
}

func (m Model) Focused() bool {
	return m.focus
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m Model) themeTitle() string {
	label := m.styles.themeTitleLabel
	text := m.styles.themeTitleText.Render(m.theme)

	theme := fmt.Sprintf("%v %v", label, text)

	if m.focus {
		return m.styles.themeTitleFocus.Render(theme)
	}

	return m.styles.themeTitle.Render(theme)
}

func (m Model) themeList() string {
	theme := m.filterList.View()

	if m.focus {
		return m.styles.themeListBoundaryFocus.Render(theme)
	}

	return m.styles.themeListBoundary.Render(theme)
}

func (m *Model) setTheme() {
	m.theme = m.filterList.SelectedItem().(listItem).Title()
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
