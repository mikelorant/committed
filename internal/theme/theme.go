package theme

import (
	"os"

	"github.com/mikelorant/committed/internal/config"

	tint "github.com/lrstanley/bubbletint"
	"github.com/muesli/termenv"
)

type Theme struct {
	ID       string
	Registry *tint.Registry
}

func New(clr config.Colour) Theme {
	ts := tints(clr)
	reg := tint.NewRegistry(ts[0], ts[1:]...)

	return Theme{
		ID:       reg.ID(),
		Registry: reg,
	}
}

func (t *Theme) Next() {
	ids := t.ListID()
	l := len(t.ListID())

	switch {
	case t.ID == ids[l-1]:
		t.Set(ids[0])
	default:
		t.Registry.NextTint()
	}

	t.ID = t.Registry.ID()
}

func (t *Theme) Set(id string) bool {
	if ok := t.Registry.SetTintID(id); !ok {
		return false
	}

	t.ID = id

	return true
}

func (t *Theme) ListID() []string {
	var ts []string

	for _, t := range t.Registry.Tints() {
		ts = append(ts, t.ID())
	}

	return ts
}

func (t *Theme) List() []tint.Tint {
	return t.Registry.Tints()
}

func tints(clr config.Colour) []tint.Tint {
	if clr == config.ColourDark {
		return dark()
	}

	if clr == config.ColourLight {
		return light()
	}

	if termenv.NewOutput(os.Stdout).HasDarkBackground() {
		return dark()
	}

	return light()
}

func dark() []tint.Tint {
	return []tint.Tint{
		tint.TintBuiltinDark,
		tint.TintGruvboxDark,
		tint.TintSolarizedDarkHigherContrast,
		tint.TintRetrowave,
		tint.TintDracula,
		tint.TintNord,
		tint.TintTokyoNight,
		tint.TintSnazzy,
	}
}

func light() []tint.Tint {
	return []tint.Tint{
		tint.TintBuiltinLight,
		tint.TintGruvboxLight,
		tint.TintBuiltinSolarizedLight,
		tint.TintBuiltinTangoLight,
		tint.TintTokyoNightLight,
	}
}
