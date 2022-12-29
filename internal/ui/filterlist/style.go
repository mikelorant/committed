package filterlist

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	boundary      lipgloss.Style
	itemPrompt    lipgloss.Border
	selectedTitle lipgloss.Style
	promptMark    lipgloss.Style
	promptText    lipgloss.Style
	paginatorDots lipgloss.Style
}

const (
	listPrompt = "❯"

	paginatorDot       = "○"
	paginatorActiveDot = "●"
)

func defaultStyles() Styles {
	var s Styles

	tint := theme.Tint()

	s.boundary = lipgloss.NewStyle().
		Width(74).
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(tint.Fg())

	// Item prompt is set as a left border character.
	s.itemPrompt = lipgloss.Border{
		Left: listPrompt,
	}

	// Assign border style to the selected item.
	s.selectedTitle = lipgloss.NewStyle().
		Border(s.itemPrompt, false, false, false, true).
		BorderForeground(tint.Cyan()).
		Foreground(tint.Cyan()).
		Padding(0, 0, 0, 1)

	s.promptMark = lipgloss.NewStyle().
		Foreground(tint.Green()).
		MarginRight(1)

	s.promptText = lipgloss.NewStyle().
		Foreground(tint.Fg()).
		Bold(true).
		MarginRight(1)

	s.paginatorDots = lipgloss.NewStyle().
		Foreground(tint.Cyan())

	return s
}
