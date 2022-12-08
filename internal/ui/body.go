package ui

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type BodyModel struct {
	config   BodyConfig
	focus    bool
	textArea textarea.Model
}

type BodyConfig struct {
	body string
}

func NewBody(cfg commit.Config) BodyModel {
	c := BodyConfig{
		body: cfg.Body,
	}

	return BodyModel{
		config:   c,
		textArea: textArea(c.body),
	}
}

func (m BodyModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m BodyModel) Update(msg tea.Msg) (BodyModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch {
	case m.focus && !m.textArea.Focused():
		cmd = m.textArea.Focus()
		return m, cmd
	case !m.focus && m.textArea.Focused():
		m.textArea.Blur()
		return m, nil
	}

	m.textArea, cmd = m.textArea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m BodyModel) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.body())
}

func (m *BodyModel) body() string {
	return lipgloss.NewStyle().
		Width(74).
		Height(19).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		Render(m.textArea.View())
}

func textArea(str string) textarea.Model {
	ta := textarea.New()
	ta.Placeholder = str
	ta.Prompt = ""
	ta.ShowLineNumbers = false
	ta.SetHeight(19)
	ta.SetWidth(72)

	return ta
}
