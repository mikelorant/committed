package shortcut

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	boundary     lipgloss.Style
	blockLeft    lipgloss.Style
	blockRight   lipgloss.Style
	columnLeft   lipgloss.Style
	columnRight  lipgloss.Style
	key          lipgloss.Style
	label        lipgloss.Style
	modifierPlus lipgloss.Style
	angleBracket lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	colour := theme.Shortcut()

	s.boundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.blockLeft = lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Left)

	s.blockRight = lipgloss.NewStyle().
		Width(30).
		Align(lipgloss.Right)

	s.columnRight = lipgloss.NewStyle().
		MarginLeft(1)

	s.columnLeft = lipgloss.NewStyle().
		MarginRight(1)

	s.key = lipgloss.NewStyle().
		Foreground(colour.Key)

	s.label = lipgloss.NewStyle().
		Foreground(colour.Label)

	s.modifierPlus = lipgloss.NewStyle().
		Foreground(colour.Plus).
		SetString("+")

	s.angleBracket = lipgloss.NewStyle().
		Foreground(colour.AngleBracket)

	return s
}
