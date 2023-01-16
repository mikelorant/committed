package shortcut

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Shortcuts Shortcuts
	styles    Styles
	view      string
}

type Shortcuts struct {
	KeyBindings []KeyBinding
	Modifiers   []Modifier
}

const (
	NoModifier = iota
	AltModifier
	ControlModifier
	ShiftModifier

	AlignLeft = iota
	AlignRight
)

func New(s Shortcuts) Model {
	return Model{
		Shortcuts: s,
		styles:    defaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg.(type) {
	case theme.Msg:
		m.styles = defaultStyles()
	}

	m.view = m.shortcutRow()

	return m, nil
}

func (m Model) View() string {
	return m.view
}

func (m Model) shortcutRow() string {
	l := newKeySet(AlignLeft, m.Shortcuts.Modifiers, m.Shortcuts.KeyBindings, true)
	r := newKeySet(AlignRight, m.Shortcuts.Modifiers, m.Shortcuts.KeyBindings, true)
	ml := newKeySet(AlignLeft, m.Shortcuts.Modifiers, modifierToKeyBinding(AlignLeft, m.Shortcuts.Modifiers), false)
	mr := newKeySet(AlignRight, m.Shortcuts.Modifiers, modifierToKeyBinding(AlignRight, m.Shortcuts.Modifiers), false)

	var left, right []string

	if len(l) != 0 {
		left = append(left, ml...)
	}
	left = append(left, l...)

	right = append(right, r...)
	if len(r) != 0 {
		right = append(right, mr...)
	}

	return m.joinShortcutRow(left, right)
}

func (m Model) joinShortcutRow(left, right []string) string {
	hleft := lipgloss.JoinHorizontal(lipgloss.Top, left...)
	hright := lipgloss.JoinHorizontal(lipgloss.Top, right...)

	bleft := m.styles.blockLeft.Render(hleft)
	bright := m.styles.blockRight.Render(hright)

	block := lipgloss.JoinHorizontal(lipgloss.Top, bleft, bright)

	return m.styles.boundary.Render(block)
}

func modifierToKeyBinding(a int, ms []Modifier) []KeyBinding {
	var ss []KeyBinding

	for _, v := range ms {
		var label string

		if v.Align != a {
			continue
		}

		if strings.TrimSpace(v.Label) != "" {
			label = defaultStyles().modifierPlus.String()
		}

		kb := KeyBinding{
			Modifier: v.Modifier,
			Key:      v.Label,
			Label:    label,
		}

		ss = append(ss, kb)
	}

	return ss
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
