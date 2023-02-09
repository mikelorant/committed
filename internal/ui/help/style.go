package help

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"
)

type Styles struct {
	boundary lipgloss.Style
	viewport lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).Help()

	s.boundary = lipgloss.NewStyle().
		Width(74).
		MarginBottom(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(clr.Boundary).
		Padding(0, 1, 0, 1)

	s.viewport = lipgloss.NewStyle().
		Foreground(clr.Viewport)

	return s
}
