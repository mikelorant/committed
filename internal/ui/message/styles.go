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

	tint := theme.Tint()

	s.message = lipgloss.NewStyle().
		MarginLeft(4).
		MarginBottom(2)

	s.summary = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.body = lipgloss.NewStyle().
		MarginTop(1).
		Foreground(tint.Fg())

	s.footer = lipgloss.NewStyle().
		MarginTop(1).
		Foreground(tint.Fg())

	return s
}
