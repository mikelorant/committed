package filterlist

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

func verticalPaginator(pos, total int, state *commit.State) string {
	return strings.Join(dots(pos, total, state), "\n")
}

func horizontalPaginator(pos, total int, state *commit.State) string {
	return strings.Join(dots(pos, total, state), "")
}

func dots(pos, total int, state *commit.State) []string {
	colour := state.Theme.FilterList()

	dots := make([]string, total)
	for i := range dots {
		dots[i] = paginatorDot
	}

	dots = append(dots[:pos], dots[pos:]...)
	dots[pos] = lipgloss.NewStyle().
		Foreground(colour.PaginatorDots).
		Render(paginatorActiveDot)

	return dots
}
