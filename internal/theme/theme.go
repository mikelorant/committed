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

type Tint struct {
	Default  tint.Tint
	Defaults []tint.Tint
}

func New(t Tint) Theme {
	reg := tint.NewRegistry(t.Default, t.Defaults...)

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

func Default(clr config.Colour) Tint {
	dark := Tint{
		Default:  tint.DefaultTints()[29],
		Defaults: tint.DefaultTints(),
	}

	light := Tint{
		Default:  tint.DefaultTints()[30],
		Defaults: tint.DefaultTints(),
	}

	switch {
	case clr == config.ColourDark:
		return dark
	case clr == config.ColourLight:
		return light
	case termenv.NewOutput(os.Stdout).HasDarkBackground():
		return dark
	}

	return light
}
