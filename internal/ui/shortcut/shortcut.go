package shortcut

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Shortcuts []Shortcut
	Modifiers []Modifier
	styles    Styles
	view      string
}

type Shortcut struct {
	Modifier int
	Key      string
	Label    string
}

type Modifier struct {
	Modifier int
	Label    string
	Align    int
}

const (
	NoModifier = iota
	AltModifier
	ControlModifier
	ShiftModifier

	AlignLeft = iota
	AlignRight
)

func NewShortcut(m []Modifier, s []Shortcut) Model {
	return Model{
		Modifiers: m,
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
	l := newShortcutSet(AlignLeft, m.Modifiers, m.Shortcuts, true)
	r := newShortcutSet(AlignRight, m.Modifiers, m.Shortcuts, true)
	ml := newShortcutSet(AlignLeft, m.Modifiers, modifierToShortcut(AlignLeft, m.Modifiers), false)
	mr := newShortcutSet(AlignRight, m.Modifiers, modifierToShortcut(AlignRight, m.Modifiers), false)

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

	bleft := m.styles.shortcutBlockLeft.Render(hleft)
	bright := m.styles.shortcutBlockRight.Render(hright)

	block := lipgloss.JoinHorizontal(lipgloss.Top, bleft, bright)

	return m.styles.shortcutBoundary.Render(block)
}

func modifierToShortcut(a int, ms []Modifier) []Shortcut {
	var ss []Shortcut

	for _, v := range ms {
		var label string

		if v.Align != a {
			continue
		}

		if strings.TrimSpace(v.Label) != "" {
			label = defaultStyles().shortcutPlus.String()
		}

		scs := Shortcut{
			Modifier: v.Modifier,
			Key:      v.Label,
			Label:    label,
		}

		ss = append(ss, scs)
	}

	return ss
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
