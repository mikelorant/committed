package status

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/ui/shortcut"
)

type Model struct {
	Next      string
	Previous  string
	Shortcuts shortcut.Model
}

func New() Model {
	return Model{
		Shortcuts: shortcut.NewShortcut(defaultModifiers(), defaultShortcuts()),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Shortcuts, _ = shortcut.ToModel(m.Shortcuts.Update(nil))

	return m, nil
}

func (m Model) View() string {
	return m.Shortcuts.View()
}

func GlobalShortcuts(next, previous string) ([]shortcut.Modifier, []shortcut.Shortcut) {
	scs := defaultShortcuts()
	mods := defaultModifiers()

	switch next {
	case "":
		mods = append(mods, shortcut.Modifier{
			Modifier: shortcut.NoModifier,
			Align:    shortcut.AlignRight,
		})

		scs = append(scs, shortcut.Shortcut{
			Modifier: shortcut.NoModifier,
		})
	default:
		mods = append(mods, shortcut.Modifier{
			Modifier: shortcut.NoModifier,
			Align:    shortcut.AlignRight,
		})

		scs = append(scs, shortcut.Shortcut{
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

		scs = append(scs, shortcut.Shortcut{
			Modifier: shortcut.ShiftModifier,
		})
	default:
		mods = append(mods, shortcut.Modifier{
			Modifier: shortcut.ShiftModifier,
			Label:    "Shift", Align: shortcut.AlignRight,
		})

		scs = append(scs, shortcut.Shortcut{
			Modifier: shortcut.ShiftModifier,
			Key:      "tab",
			Label:    previous,
		})
	}

	return mods, scs
}

func HelpShortcuts() ([]shortcut.Modifier, []shortcut.Shortcut) {
	scs := defaultShortcuts()
	mods := defaultModifiers()

	mods = append(mods, shortcut.Modifier{
		Modifier: shortcut.NoModifier,
		Align:    shortcut.AlignRight,
	})

	scs = append(scs, shortcut.Shortcut{
		Modifier: shortcut.NoModifier,
		Key:      "esc",
		Label:    "Exit",
	})

	return mods, scs
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

func defaultShortcuts() []shortcut.Shortcut {
	return []shortcut.Shortcut{
		{Modifier: shortcut.ControlModifier, Key: "c", Label: "Cancel"},
		{Modifier: shortcut.AltModifier, Key: "enter", Label: "Commit"},
		{Modifier: shortcut.AltModifier, Key: "s", Label: "Sign-off"},
		{Modifier: shortcut.AltModifier, Key: "/", Label: "Help"},
	}
}
