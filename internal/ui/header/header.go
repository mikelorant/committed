package header

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/fuzzy"
	"github.com/mikelorant/committed/internal/ui/filterlist"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Expand        bool
	DefaultHeight int
	ExpandHeight  int
	Placeholder   string
	Emoji         emoji.Emoji
	Emojis        []emoji.Emoji

	focus     bool
	component component
	styles    Styles
	height    int

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

	filterHeight     = 9
	filterPromptText = "Choose an emoji:"
)

func New(cfg commit.Config) Model {
	m := Model{
		DefaultHeight: defaultHeight,
		ExpandHeight:  expandHeight,
		Emojis:        cfg.Emojis,
		styles:        defaultStyles(),
		summaryInput:  summaryInput(cfg.Placeholders.Summary),
		filterList: filterlist.New(
			castToListItems(cfg.Emojis),
			filterPromptText,
			filterHeight,
		),
	}

	if cfg.Amend && cfg.Repository.Head.Hash != "" {
		e := commit.MessageToEmoji(cfg.Repository.Head.Message)
		if e.Valid {
			m.Emoji = e.Emoji
		}
		m.summaryInput.SetValue(commit.MessageToSummary(cfg.Repository.Head.Message))
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.filterList.Init()
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.component == emojiComponent {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.filterList.Focused() {
					m.Emoji = m.filterList.SelectedItem().(listItem).emoji
					return m, nil
				}
			case "delete":
				m.Emoji = emoji.Emoji{}
			}
		}
	}

	//nolint:gocritic
	switch msg.(type) {
	case theme.Msg:
		m.styles = defaultStyles()
		styleSummaryInput(&m.summaryInput)
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
		m.filterList, cmd = filterlist.ToModel(m.filterList.Update(msg))
		return m, cmd

	case !m.focus && m.summaryInput.Focused():
		m.summaryInput.Blur()
		return m, nil
	case !m.focus && m.filterList.Focused():
		m.filterList.Blur()
		return m, nil

	case m.focus && m.component == emojiComponent:
		ranks := fuzzy.Rank(m.filterList.Filter(), castToFuzzyItems(m.Emojis))

		items := make([]list.Item, len(ranks))
		for i, rank := range ranks {
			items[i] = castToListItems(m.Emojis)[rank]
		}
		m.filterList.SetItems(items)
	}

	m.summaryInput, cmd = m.summaryInput.Update(msg)
	cmds = append(cmds, cmd)

	m.filterList, cmd = filterlist.ToModel(m.filterList.Update(msg))
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
	subject := lipgloss.JoinHorizontal(lipgloss.Top, m.emoji(), m.summary(), m.counter())

	if !m.Expand {
		return lipgloss.NewStyle().Height(m.height).Render(subject)
	}

	fl := m.styles.filterListBoundary.Render(m.filterList.View())

	expand := lipgloss.JoinVertical(lipgloss.Top, subject, fl)

	return lipgloss.NewStyle().Height(m.height).Render(expand)
}

func (m Model) emoji() string {
	return m.styles.emojiBoundary.Render(m.Emoji.Character)
}

func (m Model) summary() string {
	return m.styles.summaryBoundary.Render(m.summaryInput.View())
}

func (m Model) counter() string {
	i := len(m.summaryInput.Value())
	if m.Emoji.Character != "" {
		i += 3
	}

	c := counterStyle(i).Render(fmt.Sprintf("%d", i))
	d := m.styles.counterDivider
	t := m.styles.counterLimit.Render(fmt.Sprintf("%d", subjectLimit))

	return m.styles.counterBoundary.Render(fmt.Sprintf("%v%v%v", c, d, t))
}

func (m Model) emojiConnector() string {
	c := connector(connecterTopRight, connectorHorizonal, connectorRightBottom, 35)

	return m.styles.emojiConnector.Render(c)
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func summaryInput(str string) textinput.Model {
	ti := textinput.New()
	ti.Prompt = ""
	ti.Placeholder = str
	ti.CharLimit = 72
	ti.Width = 59

	styleSummaryInput(&ti)

	return ti
}

func styleSummaryInput(si *textinput.Model) {
	s := defaultStyles()

	si.PromptStyle = s.summaryInputPromptStyle
	si.TextStyle = s.summaryInputTextStyle
	si.PlaceholderStyle = s.summaryInputPlaceholderStyle
	si.Cursor.Style = s.summaryInputCursorStyle
}
