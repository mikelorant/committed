package uitest

import (
	"strings"
	"unicode"

	"github.com/acarl005/stripansi"
	tea "github.com/charmbracelet/bubbletea"
)

//nolint:ireturn
func KeyPress(key rune) tea.Msg {
	return tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{key},
	}
}

func StripString(str string) string {
	s := stripansi.Strip(str)
	ss := strings.Split(s, "\n")

	var lines []string
	for _, l := range ss {
		trim := strings.TrimRightFunc(l, unicode.IsSpace)
		lines = append(lines, trim)
	}

	return strings.Join(lines, "\n")
}

//nolint:ireturn
func SendString(m tea.Model, str string) tea.Model {
	for _, r := range str {
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
		m, _ = m.Update(msg)
	}

	return m
}
