package shortcut

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"
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

	clr := colour.New(th).Shortcut()

	s.boundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.columnRight = lipgloss.NewStyle().
		MarginLeft(1)

	s.columnLeft = lipgloss.NewStyle().
		MarginRight(1)

	s.key = lipgloss.NewStyle().
		Foreground(clr.Key)

	s.label = lipgloss.NewStyle().
		Foreground(clr.Label)

	s.modifierPlus = lipgloss.NewStyle().
		Foreground(clr.Plus).
		SetString("+")

	s.angleBracket = lipgloss.NewStyle().
		Foreground(clr.AngleBracket)

	return s
}
