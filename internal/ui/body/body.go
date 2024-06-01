package body

import (
	"strings"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/ansi"
)

type Model struct {
	Height int
	Width  int

	focus    bool
	state    *commit.State
	styles   Styles
	textArea textarea.Model
}

const (
	tabSize = 4

	defaultWidth = 72
)

func New(state *commit.State, h int) Model {
	m := Model{
		Height:   h,
		Width:    defaultWidth,
		state:    state,
		styles:   defaultStyles(state.Theme),
		textArea: newTextArea(state.Placeholders.Body, defaultWidth, state),
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
	case colour.Msg:
		m.styles = defaultStyles(m.state.Theme)
		styleTextArea(&m.textArea, m.state)
		switch m.textArea.Focused() {
		case true:
			cmd = m.textArea.Focus()
			cmds = append(cmds, cmd)
		case false:
			m.textArea.Blur()
		}
	}

	m.textArea.SetHeight(m.Height)
	m.textArea.SetWidth(m.Width)

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
	if m.focus || !m.state.Config.View.HighlightActive {
		return m.styles.focusBoundary.Height(m.Height).Render(m.textArea.View())
	}

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
	// Further details for the text reflow issue:
	// https://github.com/charmbracelet/bubbles/issues/333
	l := m.Width - 1

	str := ansi.Wordwrap(m.textArea.Value(), l, "")

	return strings.TrimSpace(str)
}

func (m Model) RawValue() string {
	return m.textArea.Value()
}

func (m *Model) SetValue(str string) {
	m.textArea.SetValue(str)
}

func (m *Model) Reset() {
	m.textArea.Reset()
}

func (m *Model) CursorStart() {
	m.textArea.CursorStart()
	for i := 1; i < m.textArea.LineCount(); i++ {
		m.textArea.CursorUp()
	}
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func newTextArea(ph string, w int, state *commit.State) textarea.Model {
	ta := textarea.New()
	ta.Placeholder = ph
	ta.Prompt = ""
	ta.ShowLineNumbers = false
	ta.CharLimit = 0
	ta.EndOfBufferCharacter = ' '
	ta.SetWidth(w)

	styleTextArea(&ta, state)

	return ta
}

func styleTextArea(ta *textarea.Model, state *commit.State) {
	s := defaultStyles(state.Theme)

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
