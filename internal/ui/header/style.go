package header

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	emojiBoundary                lipgloss.Style
	emojiFocusBoundary           lipgloss.Style
	summaryBoundary              lipgloss.Style
	summaryFocusBoundary         lipgloss.Style
	counterDivider               lipgloss.Style
	counterLimit                 lipgloss.Style
	counterBoundary              lipgloss.Style
	emojiConnector               lipgloss.Style
	summaryInputPromptStyle      lipgloss.Style
	summaryInputTextStyle        lipgloss.Style
	summaryInputPlaceholderStyle lipgloss.Style
	summaryInputCursorStyle      lipgloss.Style
	readyCommitTypeBoundary      lipgloss.Style
	readyError                   lipgloss.Style
	readyIncomplete              lipgloss.Style
	readyOK                      lipgloss.Style
	commitTypeNew                lipgloss.Style
	commitTypeAmend              lipgloss.Style
	spacer                       lipgloss.Style
}

const (
	emptyCounter   = 0
	minimumCounter = 5
	warningCounter = 40
	maximumCounter = 50
)

const readyDot = "●"

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).Header()

	s.emojiBoundary = lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(clr.EmojiBoundary)

	s.emojiFocusBoundary = s.emojiBoundary.Copy().
		BorderForeground(clr.EmojiFocusBoundary)

	s.summaryBoundary = lipgloss.NewStyle().
		Width(53).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(clr.SummaryBoundary)

	s.summaryFocusBoundary = s.summaryBoundary.Copy().
		BorderForeground(clr.SummaryFocusBoundary)

	s.counterDivider = lipgloss.NewStyle().
		Foreground(clr.CounterDivider).
		SetString("/")

	s.counterLimit = lipgloss.NewStyle().
		Foreground(clr.CounterLimit)

	s.counterBoundary = lipgloss.NewStyle().
		Width(5).
		Height(3).
		MarginRight(1).
		Align(lipgloss.Right, lipgloss.Center)

	s.summaryInputPromptStyle = lipgloss.NewStyle().
		Foreground(clr.SummaryInputPromptStyle)

	s.summaryInputTextStyle = lipgloss.NewStyle().
		Foreground(clr.SummaryInputTextStyle)

	s.summaryInputPlaceholderStyle = lipgloss.NewStyle().
		Foreground(clr.SummaryInputPlaceholderStyle)

	s.summaryInputCursorStyle = lipgloss.NewStyle().
		Foreground(clr.SummaryInputCursorStyle)

	s.readyCommitTypeBoundary = lipgloss.NewStyle().
		Width(7).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center)

	s.readyError = lipgloss.NewStyle().
		Foreground(clr.ReadyError).
		MarginRight(1).
		SetString(readyDot)

	s.readyIncomplete = lipgloss.NewStyle().
		Foreground(clr.ReadyIncomplete).
		MarginRight(1).
		SetString(readyDot)

	s.readyOK = lipgloss.NewStyle().
		Foreground(clr.ReadyOK).
		MarginRight(1).
		SetString(readyDot)

	s.commitTypeNew = lipgloss.NewStyle().
		Foreground(clr.CommitTypeNew).
		Align(lipgloss.Right).
		SetString("New")

	s.commitTypeAmend = lipgloss.NewStyle().
		Foreground(clr.CommitTypeAmend).
		SetString("Amend")

	s.spacer = lipgloss.NewStyle().
		Height(1)

	return s
}

func counterStyle(i int, th theme.Theme) lipgloss.Style {
	var clr lipgloss.TerminalColor

	colour := colour.New(th).Header()

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
