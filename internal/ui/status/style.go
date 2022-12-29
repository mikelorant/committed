package status

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	shortcutBoundary    lipgloss.Style
	shortcutBlockLeft   lipgloss.Style
	shortcutBlockRight  lipgloss.Style
	shortcutColumnLeft  lipgloss.Style
	shortcutColumnRight lipgloss.Style
	shortcutKey         lipgloss.Style
	shortcutLabel       lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	tint := theme.Tint()

	s.shortcutBoundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.shortcutBlockLeft = lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Left)

	s.shortcutBlockRight = lipgloss.NewStyle().
		Width(30).
		Align(lipgloss.Right)

	s.shortcutColumnRight = lipgloss.NewStyle().
		MarginRight(1)

	s.shortcutColumnLeft = lipgloss.NewStyle().
		MarginLeft(1)

	s.shortcutKey = lipgloss.NewStyle().
		Foreground(tint.Cyan())

	s.shortcutLabel = lipgloss.NewStyle().
		Foreground(tint.Green())

	return s
}
