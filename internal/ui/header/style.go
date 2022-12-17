package header

import (
	"github.com/charmbracelet/lipgloss"
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

	s.emojiBoundary = lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder())

	s.summaryBoundary = lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder())

	s.counterLimit = lipgloss.NewStyle().
		Foreground(lipgloss.Color(white))

	s.counterBoundary = lipgloss.NewStyle().
		Width(5).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center)

	s.emojiConnector = lipgloss.NewStyle().
		MarginLeft(6)

	return s
}

func counterStyle(i int) lipgloss.Style {
	var clr string
	switch {
	case i > emptyCounter && i < minimumCounter:
		clr = yellow
	case i >= minimumCounter && i <= warningCounter:
		clr = green
	case i > warningCounter && i <= maximumCounter:
		clr = yellow
	case i > maximumCounter:
		clr = brightRed
	default:
		clr = white
	}

	bold := false
	if i > maximumCounter {
		bold = true
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(clr)).
		Bold(bold)
}
