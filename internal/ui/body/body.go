package body

import (
	"strings"
	"unicode"

	"github.com/acarl005/stripansi"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Height int
	Width  int

	focus    bool
	styles   Styles
	textArea textarea.Model
}

const (
	tabSize = 4

	defaultWidth = 72
)

func New(cfg commit.Config, h int) Model {
	m := Model{
		Height:   h,
		styles:   defaultStyles(),
		textArea: newTextArea(cfg.Placeholders.Body, defaultWidth),
	}

	if cfg.Amend && cfg.Repository.Head.Hash != "" {
		m.textArea.SetValue(commit.MessageToBody(cfg.Repository.Head.Message))
	}

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
			case "tab":
				if m.textArea.Focused() {
					m.textArea.InsertString(strings.Repeat(" ", tabSize))
				}
			}
		}
	}

	//nolint:gocritic
	switch msg.(type) {
	case theme.Msg:
		m.styles = defaultStyles()
		styleTextArea(&m.textArea)
		switch m.textArea.Focused() {
		case true:
			cmd = m.textArea.Focus()
			cmds = append(cmds, cmd)
		case false:
			m.textArea.Blur()
		}
	}

	m.textArea.SetHeight(m.Height)

	switch {
	case m.focus && !m.textArea.Focused():
		cmd = m.textArea.Focus()
		return m, cmd
	case !m.focus && m.textArea.Focused():
		m.textArea.Blur()
	}

	m.textArea, cmd = m.textArea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.styles.boundary.Height(m.Height).Render(m.textArea.View())
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

func (m Model) Value() string {
	if m.textArea.Value() == "" {
		return ""
	}

	m.textArea.Cursor.SetMode(cursor.CursorHide)
	res := strings.TrimRightFunc(stripansi.Strip(m.textArea.View()), unicode.IsSpace)
	m.textArea.Cursor.SetMode(cursor.CursorBlink)

	return res
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func newTextArea(ph string, w int) textarea.Model {
	ta := textarea.New()
	ta.Placeholder = ph
	ta.Prompt = ""
	ta.ShowLineNumbers = false
	ta.CharLimit = 0
	ta.SetWidth(w)

	styleTextArea(&ta)

	return ta
}

func styleTextArea(ta *textarea.Model) {
	s := defaultStyles()

	ta.FocusedStyle.CursorLine = s.textAreaFocusedText
	ta.FocusedStyle.Placeholder = s.textAreaPlaceholder
	ta.FocusedStyle.Prompt = s.textAreaPrompt
	ta.FocusedStyle.Text = s.textAreaFocusedText

	ta.BlurredStyle.CursorLine = s.textAreaBlurredText
	ta.BlurredStyle.Placeholder = s.textAreaPlaceholder
	ta.BlurredStyle.Prompt = s.textAreaPrompt
	ta.BlurredStyle.Text = s.textAreaBlurredText

	ta.Cursor.Style = s.textAreaCursorStyle
}
