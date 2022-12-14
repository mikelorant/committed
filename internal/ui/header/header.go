package header

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/fuzzy"
	"github.com/mikelorant/committed/internal/ui/filterlist"
)

type Model struct {
	Expand        bool
	DefaultHeight int
	ExpandHeight  int
	Placeholder   string
	Emoji         emoji.Emoji

	focus     bool
	component component
	height    int
	emojis    []emoji.Emoji

	summaryInput textinput.Model
	filterList   filterlist.Model
}

type component int

const (
	emojiComponent component = iota
	summaryComponent

	subjectLimit = 50

	defaultHeight = 3
	expandHeight  = 16
	defaultWidth  = 72

	emptyCounter   = 0
	minimumCounter = 5
	warningCounter = 40
	maximumCounter = 50

	filterPromptText = "Choose an emoji:"
)

func New(cfg commit.Config) Model {
	e, err := emoji.New()
	if err != nil {
		log.Fatal("Unable to use emojis.")
	}

	return Model{
		DefaultHeight: defaultHeight,
		ExpandHeight:  expandHeight,
		emojis:        e,
		summaryInput:  summaryInput(cfg.Summary),
		filterList: filterlist.New(
			castToListItems(e),
			filterPromptText,
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.filterList.Init()
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.component == emojiComponent {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.Emoji = m.filterList.SelectedItem().(listItem).emoji
				return m, nil
			case "delete":
				m.Emoji = emoji.Emoji{}
			}
		}
	}

	m.height = m.DefaultHeight
	if m.Expand {
		m.height = m.ExpandHeight
	}

	switch {
	case m.focus && m.component == summaryComponent && !m.summaryInput.Focused():
		m.filterList.Blur()
		cmd = m.summaryInput.Focus()
		return m, cmd
	case m.focus && m.component == emojiComponent && !m.filterList.Focused():
		m.summaryInput.Blur()
		m.filterList.Focus()
		m.filterList, cmd = m.filterList.Update(msg)
		return m, cmd

	case !m.focus && m.summaryInput.Focused():
		m.summaryInput.Blur()
		return m, nil
	case !m.focus && m.filterList.Focused():
		m.filterList.Blur()
		return m, nil

	case m.focus && m.component == emojiComponent:
		ranks := fuzzy.Rank(m.filterList.Filter(), castToFuzzyItems(m.emojis))

		items := make([]list.Item, len(ranks))
		for i, rank := range ranks {
			items[i] = castToListItems(m.emojis)[rank]
		}
		m.filterList.SetItems(items)
	}

	m.summaryInput, cmd = m.summaryInput.Update(msg)
	cmds = append(cmds, cmd)

	m.filterList, cmd = m.filterList.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		Render(m.headerRow())
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

func (m *Model) SelectEmoji() {
	m.component = emojiComponent
}

func (m *Model) SelectSummary() {
	m.component = summaryComponent
}

func (m Model) Summary() string {
	return m.summaryInput.Value()
}

func (m Model) headerRow() string {
	subject := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.emoji(),
		m.summary(),
		m.counter(),
	)

	if !m.Expand {
		return lipgloss.NewStyle().
			Height(m.height).
			Render(subject)
	}

	expand := lipgloss.JoinVertical(
		lipgloss.Top,
		subject,
		m.emojiConnector(),
		m.filterList.View(),
	)

	return lipgloss.NewStyle().
		Height(m.height).
		Render(expand)
}

func (m Model) emoji() string {
	return lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.Emoji.Character)
}

func (m Model) summary() string {
	return lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.summaryInput.View())
}

func (m Model) counter() string {
	i := len(m.summaryInput.Value())
	if m.Emoji.Character != "" {
		i += 3
	}

	clr := counterColour(i)
	bold := false
	if i > maximumCounter {
		bold = true
	}

	c := lipgloss.NewStyle().
		Foreground(lipgloss.Color(clr)).
		Bold(bold).
		Render(fmt.Sprintf("%d", i))

	t := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render(fmt.Sprintf("%d", subjectLimit))

	return lipgloss.NewStyle().
		Width(5).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center).
		Render(fmt.Sprintf("%s/%s", c, t))
}

func (m Model) emojiConnector() string {
	c := connector(connecterTopRight, connectorHorizonal, connectorRightBottom, 35)

	return lipgloss.NewStyle().
		MarginLeft(6).
		Render(c)
}

func summaryInput(str string) textinput.Model {
	ti := textinput.New()
	ti.Prompt = ""
	ti.Placeholder = str
	ti.CharLimit = 72
	ti.Width = 59

	return ti
}

func counterColour(i int) string {
	var clr string
	switch {
	case i > emptyCounter && i < minimumCounter:
		clr = yellow
	case i > warningCounter && i <= maximumCounter:
		clr = yellow
	case i > maximumCounter:
		clr = brightRed
	default:
		clr = green
	}

	return clr
}
