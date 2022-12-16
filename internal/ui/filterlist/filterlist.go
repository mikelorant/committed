package filterlist

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	PromptText string
	Height     int

	focus        bool
	list         list.Model
	textInput    textinput.Model
	selectedItem list.Item
	items        []list.Item
	filter       string
}

const listPrompt = "❯"

func New(items []list.Item, prompt string, h int) Model {
	return Model{
		PromptText: prompt,
		Height:     h,
		list:       newList(items, h),
		textInput:  newTextInput(prompt),
		items:      items,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
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
	e := lipgloss.JoinVertical(
		lipgloss.Top,
		m.textInput.View(),
		m.list.View(),
	)

	p := verticalPaginator(m.list.Paginator.Page, m.list.Paginator.TotalPages)
	ep := lipgloss.JoinHorizontal(lipgloss.Top, e, p)

	return lipgloss.NewStyle().
		Width(74).
		Height(m.Height - 1).
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder()).
		Render(ep)
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

func newList(is []list.Item, h int) list.Model {
	// Item prompt is set as a left border character.
	b := lipgloss.Border{
		Left: listPrompt,
	}

	// Assign border style to the selected item.
	s := list.NewDefaultItemStyles()
	s.SelectedTitle = lipgloss.NewStyle().
		Border(b, false, false, false, true).
		BorderForeground(lipgloss.Color(cyan)).
		Foreground(lipgloss.Color(cyan)).
		Padding(0, 0, 0, 1)

	// Delegate is the list of items.
	// Only the title is used and description is disabled.
	// Spacing between items set to 0.
	d := list.NewDefaultDelegate()
	d.SetSpacing(0)
	d.ShowDescription = false
	d.Styles = s

	l := list.New(is, d, 70, h)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.SetShowFilter(false)

	return l
}

func newTextInput(pt string) textinput.Model {
	promptMark := lipgloss.NewStyle().
		Foreground(lipgloss.Color(green)).
		MarginRight(1).
		Render("?")

	promptText := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Bold(true).
		MarginRight(1).
		Render(pt)

	prompt := lipgloss.JoinHorizontal(
		lipgloss.Left,
		promptMark,
		promptText,
	)

	ti := textinput.New()
	ti.Prompt = prompt
	ti.Placeholder = ""
	ti.CharLimit = 20
	ti.Width = 52

	return ti
}
