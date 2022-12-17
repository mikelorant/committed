package status

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	boundary            lipgloss.Style
	modifiers           lipgloss.Style
	modifiersBoundary   lipgloss.Style
	shortcutBoundary    lipgloss.Style
	shortcutKey         lipgloss.Style
	shortcutDecorateKey lipgloss.Style
	shortcutLabel       lipgloss.Style
}

const (
	black         = "0"
	red           = "1"
	green         = "2"
	yellow        = "3"
	blue          = "4"
	magenta       = "5"
	cyan          = "6"
	white         = "7"
	brightBlack   = "8"
	brightRed     = "9"
	brightGreen   = "10"
	brightYellow  = "11"
	brightBlue    = "12"
	brightMagenta = "13"
	brightCyan    = "14"
	brightWhite   = "15"
)

func defaultStyles() Styles {
	var s Styles

	s.boundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.modifiers = lipgloss.NewStyle().
		Foreground(lipgloss.Color(cyan)).
		Align(lipgloss.Right)

	s.modifiersBoundary = lipgloss.NewStyle().
		MarginRight(1)

	s.shortcutBoundary = lipgloss.NewStyle().
		MarginRight(1)

	s.shortcutKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color(cyan))

	s.shortcutDecorateKey = lipgloss.NewStyle().
		Align(lipgloss.Right)

	s.shortcutLabel = lipgloss.NewStyle().
		Foreground(lipgloss.Color(green)).
		Align(lipgloss.Left)

	return s
}
