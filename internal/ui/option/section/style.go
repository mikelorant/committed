package section

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	category         lipgloss.Style
	categorySelected lipgloss.Style
	categorySpacer   lipgloss.Style
	categoryPrompt   lipgloss.Style
	setting          lipgloss.Style
	settingSelected  lipgloss.Style
	settingSpacer    lipgloss.Style
	settingPrompt    lipgloss.Style
	settingJoiner    lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).OptionSection()

	s.category = lipgloss.NewStyle().
		Foreground(clr.Category)

	s.categorySelected = lipgloss.NewStyle().
		Foreground(clr.CategorySelected)

	s.categorySpacer = lipgloss.NewStyle().
		Foreground(clr.CategorySpacer).
		SetString(" ")

	s.categoryPrompt = lipgloss.NewStyle().
		Foreground(clr.CategoryPrompt).
		SetString("❯")

	s.setting = lipgloss.NewStyle().
		Foreground(clr.Setting)

	s.settingSelected = lipgloss.NewStyle().
		Foreground(clr.SettingSelected)

	s.settingSpacer = lipgloss.NewStyle().
		Foreground(clr.SettingSpacer).
		SetString("  ")

	s.settingPrompt = lipgloss.NewStyle().
		Foreground(clr.SettingPrompt).
		SetString("└▸")

	s.settingJoiner = lipgloss.NewStyle().
		Foreground(clr.SettingJoiner).
		SetString("│ ")

	return s
}
