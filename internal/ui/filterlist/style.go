package filterlist

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	boundary      lipgloss.Style
	itemPrompt    lipgloss.Border
	selectedTitle lipgloss.Style
	promptMark    lipgloss.Style
	promptText    lipgloss.Style
	paginatorDots lipgloss.Style
}

const (
	listPrompt = "❯"

	paginatorDot       = "○"
	paginatorActiveDot = "●"
)

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
		Width(74).
		MarginLeft(4).
		BorderStyle(lipgloss.NormalBorder())

	// Item prompt is set as a left border character.
	s.itemPrompt = lipgloss.Border{
		Left: listPrompt,
	}

	// Assign border style to the selected item.
	s.selectedTitle = lipgloss.NewStyle().
		Border(s.itemPrompt, false, false, false, true).
		BorderForeground(lipgloss.Color(cyan)).
		Foreground(lipgloss.Color(cyan)).
		Padding(0, 0, 0, 1)

	s.promptMark = lipgloss.NewStyle().
		Foreground(lipgloss.Color(green)).
		MarginRight(1)

	s.promptText = lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Bold(true).
		MarginRight(1)

	s.paginatorDots = lipgloss.NewStyle().
		Foreground(lipgloss.Color(cyan))

	return s
}
