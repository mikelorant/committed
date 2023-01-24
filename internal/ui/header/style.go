package header

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	emojiBoundary                lipgloss.Style
	summaryBoundary              lipgloss.Style
	counterDivider               lipgloss.Style
	counterLimit                 lipgloss.Style
	counterBoundary              lipgloss.Style
	emojiConnector               lipgloss.Style
	summaryInputPromptStyle      lipgloss.Style
	summaryInputTextStyle        lipgloss.Style
	summaryInputPlaceholderStyle lipgloss.Style
	summaryInputCursorStyle      lipgloss.Style
	spacer                       lipgloss.Style
}

const (
	emptyCounter   = 0
	minimumCounter = 5
	warningCounter = 40
	maximumCounter = 50
)

func defaultStyles() Styles {
	var s Styles

	colour := theme.Header()

	s.emojiBoundary = lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colour.EmojiBoundary)

	s.summaryBoundary = lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colour.SummaryBoundary)

	s.counterDivider = lipgloss.NewStyle().
		Foreground(colour.CounterDivider).
		SetString("/")

	s.counterLimit = lipgloss.NewStyle().
		Foreground(colour.CounterLimit)

	s.counterBoundary = lipgloss.NewStyle().
		Width(5).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center)

	s.summaryInputPromptStyle = lipgloss.NewStyle().
		Foreground(colour.SummaryInputPromptStyle)

	s.summaryInputTextStyle = lipgloss.NewStyle().
		Foreground(colour.SummaryInputTextStyle)

	s.summaryInputPlaceholderStyle = lipgloss.NewStyle().
		Foreground(colour.SummaryInputPlaceholderStyle)

	s.summaryInputCursorStyle = lipgloss.NewStyle().
		Foreground(colour.SummaryInputCursorStyle)

	s.spacer = lipgloss.NewStyle().
		Height(1)

	return s
}

func counterStyle(i int) lipgloss.Style {
	var clr lipgloss.TerminalColor

	colour := theme.Header()

	switch {
	case i > emptyCounter && i < minimumCounter:
		clr = colour.CounterLow
	case i >= minimumCounter && i <= warningCounter:
		clr = colour.CounterNormal
	case i > warningCounter && i <= maximumCounter:
		clr = colour.CounterWarning
	case i > maximumCounter:
		clr = colour.CounterHigh
	default:
		clr = colour.CounterDefault
	}

	bold := false
	if i > maximumCounter {
		bold = true
	}

	return lipgloss.NewStyle().
		Foreground(clr).
		Bold(bold)
}
