package filterlist

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	paginatorDot       = "○"
	paginatorActiveDot = "●"
)

func verticalPaginator(pos, total int) string {
	d := dots(pos, total)

	return strings.Join(d, "\n")
}

func horizontalPaginator(pos, total int) string {
	d := dots(pos, total)

	return strings.Join(d, "")
}

func dots(pos, total int) []string {
	dots := make([]string, total)
	for i := range dots {
		dots[i] = paginatorDot
	}

	dots = append(dots[:pos], dots[pos:]...)
	dots[pos] = lipgloss.NewStyle().
		Foreground(lipgloss.Color(cyan)).
		Render(paginatorActiveDot)

	return dots
}
