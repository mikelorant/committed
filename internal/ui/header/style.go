package header

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	emojiBoundary   lipgloss.Style
	summaryBoundary lipgloss.Style
	counterLimit    lipgloss.Style
	counterBoundary lipgloss.Style
	emojiConnector  lipgloss.Style
}

const (
	emptyCounter   = 0
	minimumCounter = 5
	warningCounter = 40
	maximumCounter = 50
)

func defaultStyles() Styles {
	var s Styles

	tint := theme.Tint()

	s.emojiBoundary = lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(tint.Fg())

	s.summaryBoundary = lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(tint.Fg())

	s.counterLimit = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.counterBoundary = lipgloss.NewStyle().
		Width(5).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center)

	s.emojiConnector = lipgloss.NewStyle().
		MarginLeft(6)

	return s
}

func counterStyle(i int) lipgloss.Style {
	var clr lipgloss.TerminalColor

	tint := theme.Tint()

	switch {
	case i > emptyCounter && i < minimumCounter:
		clr = tint.Yellow()
	case i >= minimumCounter && i <= warningCounter:
		clr = tint.Green()
	case i > warningCounter && i <= maximumCounter:
		clr = tint.Yellow()
	case i > maximumCounter:
		clr = tint.BrightRed()
	default:
		clr = tint.Fg()
	}

	bold := false
	if i > maximumCounter {
		bold = true
	}

	return lipgloss.NewStyle().
		Foreground(clr).
		Bold(bold)
}
