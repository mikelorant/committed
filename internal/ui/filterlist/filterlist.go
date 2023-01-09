package filterlist

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	PromptText string
	Height     int
	Width      int
	CharLimit  int

	focus        bool
	styles       Styles
	list         list.Model
	textInput    textinput.Model
	selectedItem list.Item
	items        []list.Item
	filter       string
}

const (
	defaultWidth     = 68
	defaultCharLimit = 20
)

func New(items []list.Item, prompt string, h int) Model {
	m := Model{
		PromptText: prompt,
		Height:     h,
		Width:      defaultWidth,
		CharLimit:  defaultCharLimit,
		styles:     defaultStyles(),
		items:      items,
	}

	m.list = m.newList(defaultWidth)
	m.textInput = m.newTextInput(defaultWidth, defaultCharLimit)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.focus {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "down":
				m.list.CursorDown()
				return m, nil
			case "up":
				m.list.CursorUp()
				return m, nil
			case "pgdown":
				if m.list.Paginator.OnLastPage() {
					break
				}
				m.list.NextPage()
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			case "pgup":
				m.list.PrevPage()
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			case "enter":
				m.selectedItem = m.list.SelectedItem()
				return m, nil
			case "esc":
				m.textInput.Reset()
				m.list.ResetSelected()
			}
		}
	}

	//nolint:gocritic
	switch msg.(type) {
	case theme.Msg:
		m.styles = defaultStyles()
		m.styleTextInput(&m.textInput)
		m.styleList(&m.list)
		m.styleListDelegate(&m.list)
	}

	if m.focus && !m.textInput.Focused() {
		cmd = m.textInput.Focus()
		return m, cmd
	}

	if !m.focus && m.textInput.Focused() {
		m.textInput.Blur()
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.list, cmd = m.list.Update(nil)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	e := lipgloss.JoinVertical(lipgloss.Top, m.textInput.View(), m.list.View())
	p := m.styles.paginatorBoundary.Render(m.stylePaginatorColumn())
	ep := lipgloss.JoinHorizontal(lipgloss.Top, e, p)

	return m.styles.boundary.Height(m.Height - 1).Render(ep)
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

func (m Model) Filter() string {
	return m.textInput.Value()
}

//nolint:ireturn
func (m Model) SelectedItem() list.Item {
	return m.list.SelectedItem()
}

func (m *Model) SetItems(i []list.Item) tea.Cmd {
	return m.list.SetItems(i)
}

func (m Model) newList(w int) list.Model {
	l := list.New(m.items, list.NewDefaultDelegate(), m.Width, m.Height)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.SetShowFilter(false)

	m.styleList(&l)
	m.styleListDelegate(&l)

	return l
}

func (m *Model) styleList(l *list.Model) {
	l.Styles.NoItems = m.styles.listNoItems
}

func (m *Model) styleListDelegate(l *list.Model) {
	s := list.NewDefaultItemStyles()
	s.NormalTitle = m.styles.listNormalTitle
	s.SelectedTitle = m.styles.listSelectedTitle

	// Delegate is the list of items.
	// Only the title is used and description is disabled.
	// Spacing between items set to 0.
	d := list.NewDefaultDelegate()
	d.SetSpacing(0)
	d.ShowDescription = false
	d.Styles = s

	l.SetDelegate(d)
}

func (m Model) stylePaginatorColumn() string {
	return verticalPaginator(m.list.Paginator.Page, m.list.Paginator.TotalPages)
}

func (m Model) newTextInput(w, c int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.CharLimit = m.CharLimit
	ti.Width = m.Width - lipgloss.Width(m.PromptText)

	m.styleTextInput(&ti)

	return ti
}

func (m Model) styleTextInput(ti *textinput.Model) {
	promptMark := m.styles.textInputPromptMark.Render("?")
	promptText := m.styles.textInputPromptText.Render(m.PromptText)

	ti.Prompt = lipgloss.JoinHorizontal(lipgloss.Left, promptMark, promptText)
	ti.PromptStyle = m.styles.textInputPromptStyle
	ti.TextStyle = m.styles.textInputTextStyle
	ti.PlaceholderStyle = m.styles.textInputPlaceholderStyle
	ti.Cursor.Style = m.styles.textInputCursorStyle
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
