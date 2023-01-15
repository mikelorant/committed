package shortcut

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	shortcutBoundary     lipgloss.Style
	shortcutBlockLeft    lipgloss.Style
	shortcutBlockRight   lipgloss.Style
	shortcutColumnLeft   lipgloss.Style
	shortcutColumnRight  lipgloss.Style
	shortcutKey          lipgloss.Style
	shortcutLabel        lipgloss.Style
	shortcutPlus         lipgloss.Style
	shortcutAngleBracket lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	colour := theme.Shortcut()

	s.shortcutBoundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.shortcutBlockLeft = lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Left)

	s.shortcutBlockRight = lipgloss.NewStyle().
		Width(30).
		Align(lipgloss.Right)

	s.shortcutColumnRight = lipgloss.NewStyle().
		MarginLeft(1)

	s.shortcutColumnLeft = lipgloss.NewStyle().
		MarginRight(1)

	s.shortcutKey = lipgloss.NewStyle().
		Foreground(colour.Key)

	s.shortcutLabel = lipgloss.NewStyle().
		Foreground(colour.Label)

	s.shortcutPlus = lipgloss.NewStyle().
		Foreground(colour.Plus).
		SetString("+")

	s.shortcutAngleBracket = lipgloss.NewStyle().
		Foreground(colour.AngleBracket)

	return s
}
