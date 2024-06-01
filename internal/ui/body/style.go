package body

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	boundary            lipgloss.Style
	focusBoundary       lipgloss.Style
	textAreaPlaceholder lipgloss.Style
	textAreaPrompt      lipgloss.Style
	textAreaFocusedText lipgloss.Style
	textAreaBlurredText lipgloss.Style
	textAreaCursorStyle lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).Body()

	s.boundary = lipgloss.NewStyle().
		Width(74).
		MarginTop(1).
		MarginBottom(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(clr.Boundary).
		Padding(0, 1, 0, 1)

	s.focusBoundary = s.boundary.
		BorderForeground(clr.FocusBoundary)

	s.textAreaPlaceholder = lipgloss.NewStyle().
		Foreground(clr.TextAreaPlaceholder)

	s.textAreaPrompt = lipgloss.NewStyle().
		Foreground(clr.TextAreaPrompt)

	s.textAreaFocusedText = lipgloss.NewStyle().
		Foreground(clr.TextAreaFocusedText)

	s.textAreaBlurredText = lipgloss.NewStyle().
		Foreground(clr.TextAreaBlurredText)

	s.textAreaCursorStyle = lipgloss.NewStyle().
		Foreground(clr.TextAreaCursorStyle)

	return s
}
