package setting_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/option/setting"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		panes []setting.Paner
		model func(setting.Model) setting.Model
	}

	type want struct {
		model func(setting.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "noop",
			args: args{
				panes: []setting.Paner{
					&setting.Noop{},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			state := &commit.State{
				Theme: theme.New(config.ColourAdaptive),
			}

			m := setting.New(state)
			m.AddPaneSet("Set", tt.args.panes)

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
