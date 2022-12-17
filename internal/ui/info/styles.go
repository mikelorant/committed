package info

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	dateTimeFormat string = "Mon Jan 2 15:04:05 2006 -0700"
)

type Styles struct {
	infoBoundary lipgloss.Style

	filterListBoundary lipgloss.Style

	hashText     lipgloss.Style
	hashValue    lipgloss.Style
	hashBoundary lipgloss.Style

	branchHead     lipgloss.Style
	branchLocal    lipgloss.Style
	branchGrouping lipgloss.Style
	branchRemote   lipgloss.Style

	authorText  lipgloss.Style
	authorValue lipgloss.Style

	dateText  lipgloss.Style
	dateValue lipgloss.Style
}

const (
	black         = "0"
	red           = "1"
	green         = "2"
	yellow        = "3"
	blue          = "4"
	magenta       = "5"
	cyan          = "6"
	white         = "7"
	brightBlack   = "8"
	brightRed     = "9"
	brightGreen   = "10"
	brightYellow  = "11"
	brightBlue    = "12"
	brightMagenta = "13"
	brightCyan    = "14"
	brightWhite   = "15"
)

func defaultStyles() Styles {
	var s Styles

	s.infoBoundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.filterListBoundary = lipgloss.NewStyle().
		MarginTop(1)

	s.hashText = lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow))

	s.hashValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		SetString("commit")

	s.hashBoundary = lipgloss.NewStyle().
		MarginRight(1)

	s.branchHead = lipgloss.NewStyle().
		Foreground(lipgloss.Color(brightCyan)).
		Bold(true).
		SetString("HEAD ->")

	s.branchLocal = lipgloss.NewStyle().
		Foreground(lipgloss.Color(brightGreen)).
		Bold(true)

	s.branchGrouping = lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow))

	s.branchRemote = lipgloss.NewStyle().
		Foreground(lipgloss.Color(red)).
		Bold(true)

	s.authorText = lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		SetString("author")

	s.authorValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color(white))

	s.dateText = lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		SetString("date")

	s.dateValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color(white))

	return s
}
