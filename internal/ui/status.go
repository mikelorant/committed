package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/mikelorant/committed/internal/commit"
)

type StatusModel struct{}

type keys []key

type key struct {
	key   string
	label string
}

var shortcuts = keys{
	{key: "enter", label: "commit"},
	{key: "c", label: "cancel"},
	{key: "a", label: "author"},
	{key: "e", label: "emoji"},
	{key: "s", label: "summary"},
	{key: "b", label: "body"},
}

func NewStatus(cfg commit.Config) StatusModel {
	return StatusModel{}
}

func (m StatusModel) render() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(statusRow())
}

func statusRow() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		modifier(),
		commands(shortcuts),
	)
}

func modifier() string {
	c := colour("Control +", cyan)

	return lipgloss.NewStyle().
		Width(9).
		Height(1).
		MarginRight(1).
		Render(c)
}

func commands(k keys) string {
	var str strings.Builder
	for _, v := range k {
		fmt.Fprintf(&str, "%s ", command(v.key, v.label))
	}

	return strings.TrimSpace(str.String())
}

func command(key, label string) string {
	label = cases.Title(language.English).String(label)

	k := colour(key, cyan)
	l := colour(label, green)

	return fmt.Sprintf("<%s> %s", k, l)
}
