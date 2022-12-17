package status

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Model struct {
	styles Styles
}

type shortcut struct {
	key   string
	label string
}

var altShortcuts = []shortcut{
	{key: "enter", label: "commit"},
	{key: "1", label: "author"},
	{key: "2", label: "emoji"},
	{key: "3", label: "summary"},
	{key: "4", label: "body"},
}

var ctrlShortcuts = []shortcut{
	{key: "c", label: "cancel"},
}

const (
	altKey  = "Alt"
	ctrlKey = "Ctrl"
)

func New() Model {
	return Model{
		styles: defaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return m.styles.boundary.Render(m.statusRow())
}

func (m Model) statusRow() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.modifiers(),
		m.shortcuts(),
	)
}

func (m Model) modifiers() string {
	mods := []string{altKey, ctrlKey}
	w := sliceMaxLen(mods)

	var ss []string
	for _, mod := range mods {
		str := m.styles.modifiers.Width(w).Render(mod)

		ss = append(ss, fmt.Sprintf("%s +", str))
	}
	strs := strings.Join(ss, "\n")

	return m.styles.modifiersBoundary.Width(w + 2).Height(len(mods)).Render(strs)
}

func (m Model) shortcuts() string {
	col := len(altShortcuts)
	if len(ctrlShortcuts) > col {
		col = len(ctrlShortcuts)
	}

	var ss []string
	for i := 0; i < col; i++ {
		var alt, ctrl shortcut

		if i < len(altShortcuts) {
			alt = altShortcuts[i]
		}

		if i < len(ctrlShortcuts) {
			ctrl = ctrlShortcuts[i]
		}

		ss = append(ss, m.shortcutColumn(alt, ctrl))
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		ss...,
	)
}

func (m Model) shortcutColumn(alt, ctrl shortcut) string {
	keys := m.shortcutColumnKeys(alt.key, ctrl.key)
	labels := m.shortcutColumnLabels(alt.label, ctrl.label)

	sc := lipgloss.JoinHorizontal(lipgloss.Top, keys, labels)

	return m.styles.shortcutBoundary.Render(sc)
}

func (m Model) shortcutColumnKeys(alt, ctrl string) string {
	keys := []string{alt, ctrl}
	w := sliceMaxLen(keys)

	var ss []string
	for _, key := range keys {
		if key == "" {
			ss = append(ss, "")
			continue
		}

		str := m.styles.shortcutKey.Render(key)
		wrappedStr := m.styles.shortcutDecorateKey.Width(w + 2).Render(fmt.Sprintf("<%s>", str))

		ss = append(ss, wrappedStr)
	}
	return lipgloss.NewStyle().
		MarginRight(1).
		Render(strings.Join(ss, "\n"))
}

func (m Model) shortcutColumnLabels(alt, ctrl string) string {
	labels := []string{alt, ctrl}
	w := sliceMaxLen(labels)

	var ss []string
	for _, label := range labels {
		if label == "" {
			ss = append(ss, "")
			continue
		}

		str := m.styles.shortcutLabel.Width(w).Render(title(label))

		ss = append(ss, str)
	}
	return strings.Join(ss, "\n")
}

func sliceMaxLen(ss []string) int {
	var i int
	for _, v := range ss {
		if len(v) > i {
			i = len(v)
		}
	}

	return i
}

func title(str string) string {
	return cases.Title(language.English).String(str)
}
