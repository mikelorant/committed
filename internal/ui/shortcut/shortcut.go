package shortcut

import (
	"strings"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/colour"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Shortcuts Shortcuts
	state     *commit.State
	styles    Styles
	view      string
}

type Shortcuts struct {
	State       *commit.State
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
		state:     s.State,
		styles:    defaultStyles(s.State.Theme),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg.(type) {
	case colour.Msg:
		m.styles = defaultStyles(m.Shortcuts.State.Theme)
	}

	m.view = m.shortcutRow()

	return m, nil
}

func (m Model) View() string {
	return m.view
}

func (m Model) shortcutRow() string {
	l := newKeySet(keySetConfig{
		align:       AlignLeft,
		modifiers:   m.Shortcuts.Modifiers,
		keyBindings: m.Shortcuts.KeyBindings,
		decorate:    true,
		state:       m.Shortcuts.State,
	})

	r := newKeySet(keySetConfig{
		align:       AlignRight,
		modifiers:   m.Shortcuts.Modifiers,
		keyBindings: m.Shortcuts.KeyBindings,
		decorate:    true,
		state:       m.Shortcuts.State,
	})

	ml := newKeySet(keySetConfig{
		align:       AlignLeft,
		modifiers:   m.Shortcuts.Modifiers,
		keyBindings: modifierToKeyBinding(AlignLeft, m.Shortcuts.Modifiers, m.Shortcuts.State),
		decorate:    false,
		state:       m.Shortcuts.State,
	})

	mr := newKeySet(keySetConfig{
		align:       AlignRight,
		modifiers:   m.Shortcuts.Modifiers,
		keyBindings: modifierToKeyBinding(AlignRight, m.Shortcuts.Modifiers, m.Shortcuts.State),
		decorate:    false,
		state:       m.Shortcuts.State,
	})

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
	const lineWidth = 80

	hleft := lipgloss.JoinHorizontal(lipgloss.Top, left...)
	hright := lipgloss.JoinHorizontal(lipgloss.Top, right...)

	bleft := lipgloss.PlaceHorizontal(lipgloss.Width(hleft), lipgloss.Left, hleft)
	bright := lipgloss.PlaceHorizontal(lineWidth-lipgloss.Width(hleft), lipgloss.Right, hright)

	block := lipgloss.JoinHorizontal(lipgloss.Top, bleft, bright)

	return m.styles.boundary.Render(block)
}

func modifierToKeyBinding(a int, ms []Modifier, state *commit.State) []KeyBinding {
	var ss []KeyBinding

	for _, v := range ms {
		var label string

		if v.Align != a {
			continue
		}

		if strings.TrimSpace(v.Label) != "" {
			label = defaultStyles(state.Theme).modifierPlus.String()
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
