package option

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	sectionBoundary      lipgloss.Style
	sectionBoundaryFocus lipgloss.Style
	settingBoundary      lipgloss.Style
	settingBoundaryFocus lipgloss.Style
	helpBoundary         lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).Option()

	s.sectionBoundary = lipgloss.NewStyle().
		Align(lipgloss.Left, lipgloss.Top).
		BorderForeground(clr.SectionBoundary).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 0).
		Margin(0, 1, 1, 4)

	s.sectionBoundaryFocus = lipgloss.NewStyle().
		Align(lipgloss.Left, lipgloss.Top).
		BorderForeground(clr.SectionBoundaryFocus).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 0).
		Margin(0, 1, 1, 4)

	s.settingBoundary = lipgloss.NewStyle().
		Align(lipgloss.Left, lipgloss.Top).
		BorderForeground(clr.SettingBoundary).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		MarginBottom(1)

	s.settingBoundaryFocus = lipgloss.NewStyle().
		Align(lipgloss.Left, lipgloss.Top).
		BorderForeground(clr.SettingBoundaryFocus).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		MarginBottom(1)

	s.helpBoundary = lipgloss.NewStyle().
		Align(lipgloss.Left, lipgloss.Top).
		BorderForeground(clr.HelpBoundary).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		MarginBottom(1)
	return s
}
