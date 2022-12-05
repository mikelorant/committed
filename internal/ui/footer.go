package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func footerBlock(m model) string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(footerRow(m.config.Name, m.config.Email))
}

func footerRow(n, e string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		signoff(n, e),
	)
}

func signoff(name, email string) string {
	s := colour("Signed-off-by", white)
	n := colour(name, white)
	e := colour(email, white)

	str := fmt.Sprintf("%s: %s <%s>", s, n, e)

	return lipgloss.NewStyle().
		Width(74).
		Height(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.HiddenBorder(), false, true).
		Padding(0, 1, 0, 1).
		Render(str)
}
