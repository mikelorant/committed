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

func defaultStyles() Styles {
	var s Styles

	tint := theme.Tint()

	s.boundary = lipgloss.NewStyle().
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(tint.Fg())

	s.paginatorBoundary = lipgloss.NewStyle().
		MarginRight(1)

	// Item prompt is set as a left border character.
	s.listItemPrompt = lipgloss.Border{
		Left: listPrompt,
	}

	s.listNormalTitle = lipgloss.NewStyle().
		Foreground(tint.Fg()).
		Padding(0, 0, 0, 2)

	// Assign border style to the selected item.
	s.listSelectedTitle = lipgloss.NewStyle().
		Border(s.listItemPrompt, false, false, false, true).
		BorderForeground(tint.Cyan()).
		Foreground(tint.Cyan()).
		Padding(0, 0, 0, 1)

	s.listNoItems = lipgloss.NewStyle().
		Foreground(tint.BrightBlack())

	s.textInputPromptMark = lipgloss.NewStyle().
		Foreground(tint.Green()).
		MarginRight(1)

	s.textInputPromptText = lipgloss.NewStyle().
		Foreground(tint.Fg()).
		Bold(true).
		MarginRight(1)

	s.paginatorDots = lipgloss.NewStyle().
		Foreground(tint.Cyan())

	s.textInputPromptStyle = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.textInputTextStyle = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.textInputPlaceholderStyle = lipgloss.NewStyle().
		Foreground(tint.BrightBlack())

	s.textInputCursorStyle = lipgloss.NewStyle().
		Foreground(tint.Fg())

	return s
}
