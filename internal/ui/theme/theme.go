package theme

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
)

type Theme struct {
	Registry *tint.Registry
}

type Msg int

func New() Theme {
	var reg *tint.Registry
	var t []tint.Tint

	switch lipgloss.HasDarkBackground() {
	case true:
		t = darkTints()
	case false:
		t = lightTints()
	}

	reg = tint.NewRegistry(t[0], t[1:]...)

	return Theme{
		Registry: reg,
	}
}

//nolint:ireturn
func (t *Theme) NextTint() {
	l := len(t.ListTints())
	ids := t.ListTints()

	if t.GetTint() == ids[l-1] {
		t.SetTint(ids[0])
		return
	}

	t.Registry.NextTint()
}

func (t *Theme) ListTints() []string {
	var tints []string

	for _, tint := range t.Registry.Tints() {
		tints = append(tints, tint.ID())
	}

	return tints
}

func (t *Theme) GetTint() string {
	return t.Registry.ID()
}

func (t *Theme) SetTint(id string) bool {
	return t.Registry.SetTintID(id)
}

//nolint:ireturn
func UpdateTheme() tea.Msg {
	var msg Msg

	return msg
}

func darkTints() []tint.Tint {
	return []tint.Tint{
		tint.TintBuiltinDark,
		tint.TintGruvboxDark,
		tint.TintSolarizedDarkHigherContrast,
		tint.TintRetrowave,
		tint.TintDracula,
		tint.TintNord,
		tint.TintTokyoNight,
	}
}

func lightTints() []tint.Tint {
	return []tint.Tint{
		tint.TintBuiltinLight,
		tint.TintGruvboxLight,
		tint.TintBuiltinSolarizedLight,
		tint.TintBuiltinTangoLight,
		tint.TintTokyoNightLight,
	}
}
