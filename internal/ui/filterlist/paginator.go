package filterlist

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"
)

func verticalPaginator(pos, total int) string {
	return strings.Join(dots(pos, total), "\n")
}

func horizontalPaginator(pos, total int) string {
	return strings.Join(dots(pos, total), "")
}

func dots(pos, total int) []string {
	tint := theme.Tint()

	dots := make([]string, total)
	for i := range dots {
		dots[i] = paginatorDot
	}

	dots = append(dots[:pos], dots[pos:]...)
	dots[pos] = lipgloss.NewStyle().
		Foreground(tint.Cyan()).
		Render(paginatorActiveDot)

	return dots
}
