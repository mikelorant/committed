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

	tint := theme.Tint()

	s.boundary = lipgloss.NewStyle().
		Width(74).
		MarginTop(1).
		MarginBottom(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(tint.Fg()).
		Padding(0, 1, 0, 1)

	s.textAreaPlaceholder = lipgloss.NewStyle().
		Foreground(tint.BrightBlack())

	s.textAreaPrompt = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.textAreaFocusedText = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.textAreaBlurredText = lipgloss.NewStyle().
		Foreground(tint.Fg())

	s.textAreaCursorStyle = lipgloss.NewStyle().
		Foreground(tint.Fg())

	return s
}
