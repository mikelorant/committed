package theme_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	inttheme "github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/option/theme"
	"github.com/mikelorant/committed/internal/ui/uitest"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		model func(m theme.Model) theme.Model
	}

	type want struct {
		model func(m theme.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
		},
		{
			name: "down_enter",
			args: args{
				model: func(m theme.Model) theme.Model {
					m.Focus()
					m, _ = theme.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					m, _ = theme.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))

					return m
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			state := &commit.State{
				Theme: inttheme.New(config.ColourAdaptive),
			}

			m := theme.New(state)

			if tt.args.model != nil {
				m = tt.args.model(m)
			}

			if tt.want.model != nil {
				tt.want.model(m)
			}

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
