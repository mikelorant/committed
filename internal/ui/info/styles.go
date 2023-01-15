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

	colour := theme.Info()

	s.infoBoundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.filterListBoundary = lipgloss.NewStyle().
		MarginTop(1)

	s.hashText = lipgloss.NewStyle().
		Foreground(colour.HashText).
		SetString("commit")

	s.hashValue = lipgloss.NewStyle().
		Foreground(colour.HashValue)

	s.hashBoundary = lipgloss.NewStyle().
		MarginRight(1)

	s.branchHead = lipgloss.NewStyle().
		Foreground(colour.BranchHead).
		Bold(true).
		SetString("HEAD ->")

	s.branchLocal = lipgloss.NewStyle().
		Foreground(colour.BranchLocal).
		Bold(true)

	s.branchGrouping = lipgloss.NewStyle().
		Foreground(colour.BranchGrouping)

	s.branchRemote = lipgloss.NewStyle().
		Foreground(colour.BranchRemote).
		Bold(true)

	s.colon = lipgloss.NewStyle().
		Foreground(colour.Colon).
		SetString(":")

	s.authorAngledBracket = lipgloss.NewStyle().
		Foreground(colour.AuthorAngledBracket)

	s.authorText = lipgloss.NewStyle().
		Foreground(colour.AuthorText).
		SetString("author")

	s.authorValue = lipgloss.NewStyle().
		Foreground(colour.AuthorValue)

	s.dateText = lipgloss.NewStyle().
		Foreground(colour.DateText).
		SetString("date")

	s.dateValue = lipgloss.NewStyle().
		Foreground(colour.DateValue)

	return s
}
