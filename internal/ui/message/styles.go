package message

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
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
		MarginBottom(2).
		Foreground(clr.Message)

	return s
}
