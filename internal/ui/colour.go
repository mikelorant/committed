package ui

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Colour struct {
	bold bool
}

type ansiColour int

const (
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
	brightBlack
	brightRed
	brightGreen
	brightYellow
	brightBlue
	brightMagenta
	brightCyan
	brightWhite
)

func colour(str string, clr int, opts ...func(*Colour)) string {
	c := &Colour{}

	for _, o := range opts {
		o(c)
	}

	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color(strconv.Itoa(clr))).
		Bold(c.bold).
		Render(str)
}

func WithBold(b bool) func(*Colour) {
	return func(c *Colour) {
		c.bold = b
	}
}
