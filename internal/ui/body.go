package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func bodyBlock(m model) string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(body(m.config.Body))
}

func body(str string) string {
	return lipgloss.NewStyle().
		Width(74).
		Height(19).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		Faint(true).
		Render(strings.TrimSpace(str))
}
