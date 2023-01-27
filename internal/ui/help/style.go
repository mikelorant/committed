package help

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	boundary lipgloss.Style
	viewport lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	colour := th.Help()

	s.boundary = lipgloss.NewStyle().
		Width(74).
		MarginBottom(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colour.Boundary).
		Padding(0, 1, 0, 1)

	s.viewport = lipgloss.NewStyle().
		Foreground(colour.Viewport)

	return s
}
