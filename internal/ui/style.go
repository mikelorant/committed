package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

func hash(str string) string {
	k := colour("commit", yellow)
	h := colour(str, yellow)

	return lipgloss.NewStyle().
		MarginRight(1).
		Render(fmt.Sprintf("%s %s", k, h))
}

func branchRefs(lb, rb string, brefs []string) string {
	h := colour("HEAD ->", brightCyan)

	l := colour(lb, brightGreen,
		WithBold(true),
	)

	lp := colour("(", yellow)
	rp := colour(")", yellow)
	c := colour(",", yellow)

	str := fmt.Sprintf("%s %s", h, l)

	if rb != "" {
		str += fmt.Sprintf("%s %s", c, colour(rb, red))
	}

	return fmt.Sprintf("%s%s%s", lp, str, rp)
}

func author(name, email string) string {
	k := colour("author", white)
	n := colour(name, white)
	e := colour(email, white)

	return fmt.Sprintf("%s: %s <%s>", k, n, e)
}

func date(str string) string {
	k := colour("date", white)
	d := colour(str, white)

	return fmt.Sprintf("%s:   %s", k, d)
}

func emoji(str string) string {
	return lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(str)
}

func summary(str string) string {
	return lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		Faint(true).
		Render(str)
}

func counter(count, total int) string {
	c := colour(fmt.Sprintf("%d", count), white)
	t := colour(fmt.Sprintf("%d", total), white)

	return lipgloss.NewStyle().
		Width(5).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center).
		Render(fmt.Sprintf("%s/%s", c, t))
}

func body(str string) string {
	return lipgloss.NewStyle().
		Width(74).
		Height(19).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		Faint(true).
		Render(strings.TrimSpace(str))
}

func signoff(name, email string) string {
	s := colour("Signed-off-by", white)
	n := colour(name, white)
	e := colour(email, white)

	str := fmt.Sprintf("%s: %s <%s>", s, n, e)

	return lipgloss.NewStyle().
		Width(74).
		Height(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.HiddenBorder(), false, true).
		Padding(0, 1, 0, 1).
		Render(str)
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
