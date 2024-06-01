package theme

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	themeTitle             lipgloss.Style
	themeTitleFocus        lipgloss.Style
	themeTitleLabel        lipgloss.Style
	themeTitleText         lipgloss.Style
	themeListBoundary      lipgloss.Style
	themeListBoundaryFocus lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).OptionTheme()

	s.themeTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(clr.Title).
		Width(40).
		Padding(0, 1, 0, 1).
		MarginBottom(1)

	s.themeTitleFocus = s.themeTitle.
		BorderForeground(clr.TitleFocus)

	s.themeTitleLabel = lipgloss.NewStyle().
		Foreground(clr.TitleLabel).
		SetString("Theme:")

	s.themeTitleText = lipgloss.NewStyle().
		Foreground(clr.TitleText)

	s.themeListBoundary = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(40).
		BorderForeground(clr.Boundary)

	s.themeListBoundaryFocus = s.themeListBoundary.
		BorderForeground(clr.BoundaryFocus)

	return s
}
