package info

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
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

	colon lipgloss.Style

	authorAngledBracket lipgloss.Style
	authorText          lipgloss.Style
	authorValue         lipgloss.Style

	dateText  lipgloss.Style
	dateValue lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	tint := theme.Tint()

	s.infoBoundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.filterListBoundary = lipgloss.NewStyle().
		MarginTop(1)

	s.hashText = lipgloss.NewStyle().
		Foreground(tint.Yellow()).
		SetString("commit")

	s.hashValue = lipgloss.NewStyle().
		Foreground(tint.Yellow())

	s.hashBoundary = lipgloss.NewStyle().
		MarginRight(1)

	s.branchHead = lipgloss.NewStyle().
		Foreground(tint.BrightCyan()).
		Bold(true).
		SetString("HEAD ->")

	s.branchLocal = lipgloss.NewStyle().
		Foreground(tint.BrightGreen()).
		Bold(true)

	s.branchGrouping = lipgloss.NewStyle().
		Foreground(tint.Yellow())

	s.branchRemote = lipgloss.NewStyle().
		Foreground(tint.BrightRed()).
		Bold(true)

	s.colon = lipgloss.NewStyle().
		Foreground(tint.Fg()).
		SetString(":")

	s.authorAngledBracket = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.authorText = lipgloss.NewStyle().
		Foreground(tint.Fg()).
		SetString("author")

	s.authorValue = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.dateText = lipgloss.NewStyle().
		Foreground(tint.Fg()).
		SetString("date")

	s.dateValue = lipgloss.NewStyle().
		Foreground(tint.Fg())

	return s
}
