package setting

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	setting              lipgloss.Style
	settingTitle         lipgloss.Style
	settingSelected      lipgloss.Style
	settingTitleSelected lipgloss.Style
	settingDotEmpty      lipgloss.Style
	settingDotFilled     lipgloss.Style
	settingSquareEmpty   lipgloss.Style
	settingSquareFilled  lipgloss.Style
}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	clr := colour.New(th).OptionSetting()

	s.setting = lipgloss.NewStyle().
		Foreground(clr.Setting)

	s.settingTitle = lipgloss.NewStyle().
		Foreground(clr.SettingTitle)

	s.settingSelected = lipgloss.NewStyle().
		Foreground(clr.SettingSelected)

	s.settingTitleSelected = lipgloss.NewStyle().
		Foreground(clr.SettingTitleSelected)

	s.settingDotEmpty = lipgloss.NewStyle().
		Foreground(clr.SettingDotEmpty).
		SetString("○")

	s.settingDotFilled = lipgloss.NewStyle().
		Foreground(clr.SettingDotFilled).
		SetString("●")

	s.settingSquareEmpty = lipgloss.NewStyle().
		Foreground(clr.SettingSquareEmpty).
		SetString("▢")

	s.settingSquareFilled = lipgloss.NewStyle().
		Foreground(clr.SettingSquareFilled).
		SetString("▣")

	return s
}
