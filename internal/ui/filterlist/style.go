package filterlist

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	boundary                  lipgloss.Style
	focusBoundary             lipgloss.Style
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

	clr := colour.New(th).FilterList()

	s.boundary = lipgloss.NewStyle().
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(clr.Boundary)

	s.focusBoundary = s.boundary.
		BorderForeground(clr.FocusBoundary)

	s.paginatorBoundary = lipgloss.NewStyle().
		MarginRight(1)

	// Item prompt is set as a left border character.
	s.listItemPrompt = lipgloss.Border{
		Left: listPrompt,
	}

	s.listNormalTitle = lipgloss.NewStyle().
		Foreground(clr.ListNormalTitle).
		Padding(0, 0, 0, 2)

	// Assign border style to the selected item.
	s.listSelectedTitle = lipgloss.NewStyle().
		Border(s.listItemPrompt, false, false, false, true).
		BorderForeground(clr.ListSelectedTitle).
		Foreground(clr.ListSelectedTitle).
		Padding(0, 0, 0, 1)

	s.listNoItems = lipgloss.NewStyle().
		Foreground(clr.ListNoItems)

	s.textInputPromptMark = lipgloss.NewStyle().
		Foreground(clr.TextInputPromptMark).
		MarginRight(1)

	s.textInputPromptText = lipgloss.NewStyle().
		Foreground(clr.TextInputPromptText).
		Bold(true).
		MarginRight(1)

	s.paginatorDots = lipgloss.NewStyle().
		Foreground(clr.PaginatorDots)

	s.textInputPromptStyle = lipgloss.NewStyle().
		Foreground(clr.TextInputPromptStyle)

	s.textInputTextStyle = lipgloss.NewStyle().
		Foreground(clr.TextInputTextStyle)

	s.textInputPlaceholderStyle = lipgloss.NewStyle().
		Foreground(clr.TextInputPlaceholderStyle)

	s.textInputCursorStyle = lipgloss.NewStyle().
		Foreground(clr.TextInputCursorStyle)

	return s
}
