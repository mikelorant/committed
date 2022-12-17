package body

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	boundary lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	s.boundary = lipgloss.NewStyle().
		Width(74).
		MarginTop(1).
		MarginBottom(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1)

	return s
}
