package status

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/shortcut"
)

type Model struct {
	Shortcuts shortcut.Shortcuts
	shortcut  shortcut.Model
	state     *commit.State
}

func New(state *commit.State) Model {
	ds := shortcut.Shortcuts{
		Modifiers:   defaultModifiers(),
		KeyBindings: defaultKeyBindings(),
		State:       state,
	}

	return Model{
		Shortcuts: ds,
		shortcut:  shortcut.New(ds),
		state:     state,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Shortcuts.State = m.state
	m.shortcut.Shortcuts = m.Shortcuts
	m.shortcut, _ = shortcut.ToModel(m.shortcut.Update(nil))

	return m, nil
}

func (m Model) View() string {
	return m.shortcut.View()
}

func GlobalShortcuts(next, previous string) shortcut.Shortcuts {
	mods := defaultModifiers()
	kb := defaultKeyBindings()

	switch next {
	case "":
		mods = append(mods, shortcut.Modifier{
			Modifier: shortcut.NoModifier,
			Align:    shortcut.AlignRight,
		})

		kb = append(kb, shortcut.KeyBinding{
			Modifier: shortcut.NoModifier,
		})
	default:
		mods = append(mods, shortcut.Modifier{
			Modifier: shortcut.NoModifier,
			Align:    shortcut.AlignRight,
		})

		kb = append(kb, shortcut.KeyBinding{
			Modifier: shortcut.NoModifier,
			Key:      "tab",
			Label:    next,
		})
	}

	switch previous {
	case "":
		mods = append(mods, shortcut.Modifier{
			Modifier: shortcut.ShiftModifier,
			Align:    shortcut.AlignRight,
			Label:    strings.Repeat(" ", 6),
		})

		kb = append(kb, shortcut.KeyBinding{
			Modifier: shortcut.ShiftModifier,
		})
	default:
		mods = append(mods, shortcut.Modifier{
			Modifier: shortcut.ShiftModifier,
			Label:    "Shift", Align: shortcut.AlignRight,
		})

		kb = append(kb, shortcut.KeyBinding{
			Modifier: shortcut.ShiftModifier,
			Key:      "tab",
			Label:    previous,
		})
	}

	return shortcut.Shortcuts{
		Modifiers:   mods,
		KeyBindings: kb,
	}
}

func HelpShortcuts() shortcut.Shortcuts {
	kb := defaultKeyBindings()
	mods := defaultModifiers()

	mods = append(mods, shortcut.Modifier{
		Modifier: shortcut.NoModifier,
		Align:    shortcut.AlignRight,
	})

	kb = append(kb, shortcut.KeyBinding{
		Modifier: shortcut.NoModifier,
		Key:      "esc",
		Label:    "Exit",
	})

	return shortcut.Shortcuts{
		Modifiers:   mods,
		KeyBindings: kb,
	}
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func defaultModifiers() []shortcut.Modifier {
	return []shortcut.Modifier{
		{Modifier: shortcut.AltModifier, Label: "Alt", Align: shortcut.AlignLeft},
		{Modifier: shortcut.ControlModifier, Label: "Ctrl", Align: shortcut.AlignLeft},
	}
}

func defaultKeyBindings() []shortcut.KeyBinding {
	return []shortcut.KeyBinding{
		{Modifier: shortcut.ControlModifier, Key: "c", Label: "Cancel"},
		{Modifier: shortcut.AltModifier, Key: "enter", Label: "Commit"},
		{Modifier: shortcut.AltModifier, Key: "s", Label: "Sign-off"},
		{Modifier: shortcut.AltModifier, Key: "/", Label: "Help"},
	}
}
