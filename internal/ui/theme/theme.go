package theme

import (
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
)

type Msg int

var (
	registry *tint.Registry
	once     sync.Once
)

var darkTints = []tint.Tint{
	tint.TintBuiltinDark,
	tint.TintGruvboxDark,
	tint.TintSolarizedDarkHigherContrast,
	tint.TintRetrowave,
	tint.TintDracula,
	tint.TintNord,
	tint.TintTokyoNight,
}

var lightTints = []tint.Tint{
	tint.TintBuiltinLight,
	tint.TintGruvboxLight,
	tint.TintBuiltinSolarizedLight,
	tint.TintBuiltinTangoLight,
	tint.TintTokyoNightLight,
}

func Tint() *tint.Registry {
	var t []tint.Tint

	once.Do(func() {
		switch lipgloss.HasDarkBackground() {
		case true:
			t = darkTints
		case false:
			t = lightTints
		}

		registry = tint.NewRegistry(t[0], t[1:]...)
	})

	return registry
}

//nolint:ireturn
func NextTint() tea.Msg {
	var msg Msg

	l := len(registry.TintIDs())
	ids := registry.TintIDs()

	if registry.ID() == ids[l-1] {
		registry.SetTintID(ids[0])
		return msg
	}

	registry.NextTint()

	return msg
}
