package info

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"
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
	branchTag      lipgloss.Style

	colon lipgloss.Style

	authorAngledBracket lipgloss.Style
	authorText          lipgloss.Style
	authorValue         lipgloss.Style

	dateText  lipgloss.Style
	dateValue lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).Info()

	s.infoBoundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.filterListBoundary = lipgloss.NewStyle().
		MarginTop(1)

	s.hashText = lipgloss.NewStyle().
		Foreground(clr.HashText).
		SetString("commit")

	s.hashValue = lipgloss.NewStyle().
		Foreground(clr.HashValue)

	s.hashBoundary = lipgloss.NewStyle().
		MarginRight(1)

	s.branchHead = lipgloss.NewStyle().
		Foreground(clr.BranchHead).
		Bold(true).
		SetString("HEAD ->")

	s.branchLocal = lipgloss.NewStyle().
		Foreground(clr.BranchLocal).
		Bold(true)

	s.branchGrouping = lipgloss.NewStyle().
		Foreground(clr.BranchGrouping)

	s.branchRemote = lipgloss.NewStyle().
		Foreground(clr.BranchRemote).
		Bold(true)

	s.branchTag = lipgloss.NewStyle().
		Foreground(clr.BranchTag).
		Bold(true)

	s.colon = lipgloss.NewStyle().
		Foreground(clr.Colon).
		SetString(":")

	s.authorAngledBracket = lipgloss.NewStyle().
		Foreground(clr.AuthorAngledBracket)

	s.authorText = lipgloss.NewStyle().
		Foreground(clr.AuthorText).
		SetString("author")

	s.authorValue = lipgloss.NewStyle().
		Foreground(clr.AuthorValue)

	s.dateText = lipgloss.NewStyle().
		Foreground(clr.DateText).
		SetString("date")

	s.dateValue = lipgloss.NewStyle().
		Foreground(clr.DateValue)

	return s
}
