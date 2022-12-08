package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type HeaderModel struct {
	config    HeaderConfig
	focus     bool
	textInput textinput.Model
}

type HeaderConfig struct {
	emoji   string
	summary string
}

const (
	subjectLimit int = 50
)

func NewHeader(cfg commit.Config) HeaderModel {
	c := HeaderConfig{
		emoji:   cfg.Emoji,
		summary: cfg.Summary,
	}

	return HeaderModel{
		config:    c,
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

	switch {
	case m.focus && !m.textInput.Focused():
		cmd = m.textInput.Focus()
		return m, cmd
	case !m.focus && m.textInput.Focused():
		m.textInput.Blur()
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m HeaderModel) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.headerRow())
}

func (m HeaderModel) headerRow() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.emoji(),
		m.summary(),
		m.counter(),
	)
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

func textInput(str string) textinput.Model {
	ti := textinput.New()
	ti.Prompt = ""
	ti.Placeholder = str
	ti.CharLimit = 72
	ti.Width = 59

	return ti
}
