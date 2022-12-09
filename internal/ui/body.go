package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type BodyModel struct {
	config   BodyConfig
	state    State
	textArea textarea.Model
}

type BodyConfig struct {
	body string
}

const (
	tabSize = 4
)

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

	if m.state.component == bodyComponent {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "tab":
				m.textArea.InsertString(strings.Repeat(" ", tabSize))
			}
		}
	}

	switch {
	case m.state.component == bodyComponent && !m.textArea.Focused():
		cmd = m.textArea.Focus()
		return m, cmd
	case m.state.component != bodyComponent && m.textArea.Focused():
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
