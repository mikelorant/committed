package ui

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/emoji"
)

type HeaderModel struct {
	config      HeaderConfig
	state       State
	height      int
	textInput   textinput.Model
	emojiList   list.Model
	emojiFilter textinput.Model
}

type HeaderConfig struct {
	emoji  emoji.Emoji
	emojis []emoji.Emoji

	summary string
}

type connectorStyle string

const (
	subjectLimit int = 50

	headerDefault = 3
	headerExpand  = 16

	emojiListPrompt = "❯"

	connecterTopRight    connectorStyle = "└"
	connectopTopLeft     connectorStyle = "┘"
	connectorBottomLeft  connectorStyle = "┐"
	connectorBottomRight connectorStyle = "┌"
	connectorLeftTop     connectorStyle = "└"
	connectorLeftBottom  connectorStyle = "┌"
	connectorRightBottom connectorStyle = "┐"
	connectorRightTop    connectorStyle = "┘"
	connectorHorizonal   connectorStyle = "─"
	connectorVertical    connectorStyle = "│"

	paginatorDot       = "○"
	paginatorActiveDot = "●"
)

func NewHeader(cfg commit.Config) HeaderModel {
	c := HeaderConfig{
		summary: cfg.Summary,
	}

	e, err := emoji.New()
	if err != nil {
		log.Fatal("Unable to use emojis.")
	}
	c.emojis = e

	return HeaderModel{
		config:      c,
		height:      headerDefault,
		textInput:   textInput(c.summary),
		emojiList:   emojiList(castToItems(c.emojis)),
		emojiFilter: emojiFilter(),
	}
}

func (m HeaderModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m HeaderModel) Update(msg tea.Msg) (HeaderModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch m.state.display {
	case expandedDisplay:
		m.height = headerExpand
	default:
		m.height = headerDefault
	}

	if m.state.component == emojiComponent {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "down":
				m.emojiList.CursorDown()
				return m, nil
			case "up":
				m.emojiList.CursorUp()
				return m, nil
			case "pgdown":
				if m.emojiList.Paginator.OnLastPage() {
					break
				}
				m.emojiList.NextPage()
				m.emojiList, cmd = m.emojiList.Update(msg)
				return m, cmd
			case "pgup":
				m.emojiList.PrevPage()
				m.emojiList, cmd = m.emojiList.Update(msg)
				return m, cmd
			case "enter":
				m.config.emoji = m.emojiList.SelectedItem().(emoji.Emoji)
				return m, nil
			case "esc":
				m.emojiFilter.Reset()
				m.emojiList.ResetSelected()
			case "delete":
				m.config.emoji = emoji.Emoji{}
			}
		}
	}

	switch {
	case m.state.component == summaryComponent && !m.textInput.Focused():
		cmd = m.textInput.Focus()
		return m, cmd
	case m.state.component != summaryComponent && m.textInput.Focused():
		m.textInput.Blur()
		return m, nil
	case m.state.component == emojiComponent && !m.emojiFilter.Focused():
		cmd = m.emojiFilter.Focus()
		return m, cmd
	case m.state.component != emojiComponent && m.emojiFilter.Focused():
		m.emojiFilter.Blur()
		return m, nil
	case m.state.component == emojiComponent:
		items := fuzzyFilter(m.emojiFilter.Value(), castToItems(m.config.emojis))
		if len(items) > 0 {
			m.emojiList.SetItems(items)
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	m.emojiFilter, cmd = m.emojiFilter.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m HeaderModel) View() string {
	return lipgloss.NewStyle().
		Render(m.headerRow())
}

func (m HeaderModel) EmojiShortCode() string {
	return m.config.emoji.ShortCode()
}

func (m HeaderModel) Summary() string {
	return m.textInput.Value()
}

func (m HeaderModel) headerRow() string {
	subject := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.emoji(),
		m.summary(),
		m.counter(),
	)

	if m.state.display == defaultDisplay {
		return lipgloss.NewStyle().
			Height(m.height).
			Render(subject)
	}

	expand := lipgloss.JoinVertical(
		lipgloss.Top,
		subject,
		m.emojiChooserConnector(),
		m.emojiChooser(),
	)

	return lipgloss.NewStyle().
		Height(m.height).
		Render(expand)
}

func (m HeaderModel) emoji() string {
	return lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.config.emoji.Character())
}

func (m HeaderModel) summary() string {
	return lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.textInput.View())
}

func (m HeaderModel) counter() string {
	i := len(m.textInput.Value())
	if m.config.emoji.Character() != "" {
		i += 2
	}

	clr := white
	bold := false
	switch {
	case i > 0 && i < 5:
		clr = yellow
	case i >= 5 && i <= 40:
		clr = green
	case i > 40 && i <= 50:
		clr = yellow
	case i > 50:
		clr = brightRed
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

func (m HeaderModel) emojiChooser() string {
	e := lipgloss.JoinVertical(
		lipgloss.Top,
		m.emojiFilter.View(),
		m.emojiList.View(),
	)

	p := verticalPaginator(m.emojiList.Paginator.Page, m.emojiList.Paginator.TotalPages)
	ep := lipgloss.JoinHorizontal(lipgloss.Top, e, p)

	return lipgloss.NewStyle().
		Width(74).
		Height(10).
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder()).
		Render(ep)
}

func (m HeaderModel) emojiChooserConnector() string {
	c := connector(connecterTopRight, connectorHorizonal, connectorRightBottom, 35)

	return lipgloss.NewStyle().
		MarginLeft(6).
		Render(c)
}

func textInput(str string) textinput.Model {
	ti := textinput.New()
	ti.Prompt = ""
	ti.Placeholder = str
	ti.CharLimit = 72
	ti.Width = 59

	return ti
}

func emojiList(ei []list.Item) list.Model {
	// Item prompt is set as a left border character.
	b := lipgloss.Border{
		Left: emojiListPrompt,
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

	l := list.New(ei, d, 70, 9)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.SetShowFilter(false)

	return l
}

func emojiFilter() textinput.Model {
	promptMark := lipgloss.NewStyle().
		Foreground(lipgloss.Color(green)).
		MarginRight(1).
		Render("?")

	promptText := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Bold(true).
		MarginRight(1).
		Render("Choose an emoji:")

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

func connector(start, axes, end connectorStyle, len int) string {
	return fmt.Sprintf("%v%v%v", start, strings.Repeat(string(axes), len), end)
}

func verticalPaginator(pos, total int) string {
	dots := make([]string, total)
	for i := range dots {
		dots[i] = paginatorDot
	}

	dots = append(dots[:pos], dots[pos:]...)
	dots[pos] = lipgloss.NewStyle().
		Foreground(lipgloss.Color(cyan)).
		Render(paginatorActiveDot)

	return strings.Join(dots, "\n")
}

func fuzzyFilter(term string, items []list.Item) []list.Item {
	if len(term) < 3 {
		return items
	}

	var lines []string
	for i, v := range items {
		name := fmt.Sprintf("%v:%s", i, v.(emoji.Emoji).Name())
		shortcode := fmt.Sprintf("%v:%s", i, v.(emoji.Emoji).ShortCode())
		description := fmt.Sprintf("%v:%s", i, v.(emoji.Emoji).Description())
		lines = append(lines, name, shortcode, description)
	}

	ranks := fuzzy.RankFindFold(term, lines)
	sort.Sort(ranks)

	var index []int
	for _, v := range ranks {
		pos := strings.Split(v.Target, ":")[0]
		i, _ := strconv.Atoi(pos)
		if !contains(index, i) {
			index = append(index, i)
		}
	}

	var matches []list.Item
	for _, i := range index {
		matches = append(matches, items[i])
	}

	return matches
}

func castToItems[T list.Item](items []T) []list.Item {
	var result []list.Item
	for _, item := range items {
		result = append(result, item)
	}
	return result
}

func contains[T comparable](vs []T, val T) bool {
	for _, v := range vs {
		if v == val {
			return true
		}
	}
	return false
}
