package filterlist

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	boundary                  lipgloss.Style
	listItemPrompt            lipgloss.Border
	listNormalTitle           lipgloss.Style
	listSelectedTitle         lipgloss.Style
	listNoItems               lipgloss.Style
	textInputPromptMark       lipgloss.Style
	textInputPromptText       lipgloss.Style
	textInputPromptStyle      lipgloss.Style
	textInputTextStyle        lipgloss.Style
	textInputPlaceholderStyle lipgloss.Style
	textInputCursorStyle      lipgloss.Style
	paginatorBoundary         lipgloss.Style
	paginatorDots             lipgloss.Style
}

const (
	listPrompt = "❯"

	paginatorDot       = "○"
	paginatorActiveDot = "●"
)

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	colour := th.FilterList()

	s.boundary = lipgloss.NewStyle().
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colour.Boundary)

	s.paginatorBoundary = lipgloss.NewStyle().
		MarginRight(1)

	// Item prompt is set as a left border character.
	s.listItemPrompt = lipgloss.Border{
		Left: listPrompt,
	}

	s.listNormalTitle = lipgloss.NewStyle().
		Foreground(colour.ListNormalTitle).
		Padding(0, 0, 0, 2)

	// Assign border style to the selected item.
	s.listSelectedTitle = lipgloss.NewStyle().
		Border(s.listItemPrompt, false, false, false, true).
		BorderForeground(colour.ListSelectedTitle).
		Foreground(colour.ListSelectedTitle).
		Padding(0, 0, 0, 1)

	s.listNoItems = lipgloss.NewStyle().
		Foreground(colour.ListNoItems)

	s.textInputPromptMark = lipgloss.NewStyle().
		Foreground(colour.TextInputPromptMark).
		MarginRight(1)

	s.textInputPromptText = lipgloss.NewStyle().
		Foreground(colour.TextInputPromptText).
		Bold(true).
		MarginRight(1)

	s.paginatorDots = lipgloss.NewStyle().
		Foreground(colour.PaginatorDots)

	s.textInputPromptStyle = lipgloss.NewStyle().
		Foreground(colour.TextInputPromptStyle)

	s.textInputTextStyle = lipgloss.NewStyle().
		Foreground(colour.TextInputTextStyle)

	s.textInputPlaceholderStyle = lipgloss.NewStyle().
		Foreground(colour.TextInputPlaceholderStyle)

	s.textInputCursorStyle = lipgloss.NewStyle().
		Foreground(colour.TextInputCursorStyle)

	return s
}
