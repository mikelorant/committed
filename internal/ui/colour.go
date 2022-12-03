package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type ansiColour int

const (
	black         ansiColour = 0
	red           ansiColour = 1
	green         ansiColour = 2
	yellow        ansiColour = 3
	blue          ansiColour = 4
	magenta       ansiColour = 5
	cyan          ansiColour = 6
	white         ansiColour = 7
	brightBlack   ansiColour = 8
	brightRed     ansiColour = 9
	brightGreen   ansiColour = 10
	brightYellow  ansiColour = 11
	brightBlue    ansiColour = 12
	brightMagenta ansiColour = 13
	brightCyan    ansiColour = 14
	brightWhite   ansiColour = 15
)

func colour(str string, c ansiColour) string {
	clr := lipgloss.Color(fmt.Sprintf("%d", c))

	return lipgloss.NewStyle().Foreground(clr).Render(str)
}
