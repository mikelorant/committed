package filterlist

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func verticalPaginator(pos, total int) string {
	return strings.Join(dots(pos, total), "\n")
}

func horizontalPaginator(pos, total int) string {
	return strings.Join(dots(pos, total), "")
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
