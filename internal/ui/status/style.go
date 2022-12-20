package status

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	shortcutBoundary    lipgloss.Style
	shortcutBlockLeft   lipgloss.Style
	shortcutBlockRight  lipgloss.Style
	shortcutColumnLeft  lipgloss.Style
	shortcutColumnRight lipgloss.Style
	shortcutKey         lipgloss.Style
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

	s.shortcutBoundary = lipgloss.NewStyle().
		MarginBottom(1)

	s.shortcutBlockLeft = lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Left)

	s.shortcutBlockRight = lipgloss.NewStyle().
		Width(30).
		Align(lipgloss.Right)

	s.shortcutColumnRight = lipgloss.NewStyle().
		MarginRight(1)

	s.shortcutColumnLeft = lipgloss.NewStyle().
		MarginLeft(1)

	s.shortcutKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color(cyan))

	s.shortcutLabel = lipgloss.NewStyle().
		Foreground(lipgloss.Color(green))

	return s
}
