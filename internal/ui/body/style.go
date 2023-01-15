package body

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Styles struct {
	boundary            lipgloss.Style
	textAreaPlaceholder lipgloss.Style
	textAreaPrompt      lipgloss.Style
	textAreaFocusedText lipgloss.Style
	textAreaBlurredText lipgloss.Style
	textAreaCursorStyle lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	colour := theme.Body()

	s.boundary = lipgloss.NewStyle().
		Width(74).
		MarginTop(1).
		MarginBottom(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colour.Boundary).
		Padding(0, 1, 0, 1)

	s.textAreaPlaceholder = lipgloss.NewStyle().
		Foreground(colour.TextAreaPlaceholder)

	s.textAreaPrompt = lipgloss.NewStyle().
		Foreground(colour.TextAreaPrompt)

	s.textAreaFocusedText = lipgloss.NewStyle().
		Foreground(colour.TextAreaFocusedText)

	s.textAreaBlurredText = lipgloss.NewStyle().
		Foreground(colour.TextAreaBlurredText)

	s.textAreaCursorStyle = lipgloss.NewStyle().
		Foreground(colour.TextAreaCursorStyle)

	return s
}
