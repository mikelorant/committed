package message

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"
)

type Styles struct {
	message lipgloss.Style
	summary lipgloss.Style
	body    lipgloss.Style
	footer  lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).Message()

	s.message = lipgloss.NewStyle().
		MarginLeft(4).
		MarginBottom(2)

	s.summary = lipgloss.NewStyle().
		Foreground(clr.Summary)

	s.body = lipgloss.NewStyle().
		MarginTop(1).
		Foreground(clr.Body)

	s.footer = lipgloss.NewStyle().
		MarginTop(1).
		Foreground(clr.Footer)

	return s
}
