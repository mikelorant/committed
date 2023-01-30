package shortcut

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	boundary     lipgloss.Style
	columnLeft   lipgloss.Style
	columnRight  lipgloss.Style
	key          lipgloss.Style
	label        lipgloss.Style
	modifierPlus lipgloss.Style
	angleBracket lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	colour := th.Shortcut()

	s.boundary = lipgloss.NewStyle().
		MarginBottom(1)

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
