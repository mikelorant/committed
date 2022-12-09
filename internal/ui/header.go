package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/emoji"
)

type HeaderModel struct {
	config    HeaderConfig
	state     State
	height    int
	textInput textinput.Model
}

type HeaderConfig struct {
	emoji   string
	summary string
}

type connectorStyle string

const (
	subjectLimit int = 50

	headerDefault = 3
	headerExpand  = 16

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
)

func NewHeader(cfg commit.Config) HeaderModel {
	c := HeaderConfig{
		emoji:   cfg.Emoji,
		summary: cfg.Summary,
	}

	return HeaderModel{
		config:    c,
		height:    headerDefault,
		textInput: textInput(c.summary),
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

	switch {
	case m.state.component == summaryComponent && !m.textInput.Focused():
		cmd = m.textInput.Focus()
		return m, cmd
	case m.state.component != summaryComponent && m.textInput.Focused():
		m.textInput.Blur()
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m HeaderModel) View() string {
	return lipgloss.NewStyle().
		Render(m.headerRow())
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
		Render(m.config.emoji)
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
	if m.config.emoji != "" {
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
	return lipgloss.NewStyle().
		Width(74).
		Height(10).
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder()).
		Render(emoji.Emoji())
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

func connector(start, axes, end connectorStyle, len int) string {
	return fmt.Sprintf("%v%v%v", start, strings.Repeat(string(axes), len), end)
}
