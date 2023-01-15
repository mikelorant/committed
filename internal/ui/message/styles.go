package message

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	message lipgloss.Style
	summary lipgloss.Style
	body    lipgloss.Style
	footer  lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	colour := theme.Message()

	s.message = lipgloss.NewStyle().
		MarginLeft(4).
		MarginBottom(2)

	s.summary = lipgloss.NewStyle().
		Foreground(colour.Summary)

	s.body = lipgloss.NewStyle().
		MarginTop(1).
		Foreground(colour.Body)

	s.footer = lipgloss.NewStyle().
		MarginTop(1).
		Foreground(colour.Footer)

	return s
}
