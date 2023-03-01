package option_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/option"
	"github.com/mikelorant/committed/internal/ui/option/section"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		model func(m option.Model) option.Model
	}

	type want struct {
		model func(m option.Model)
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
			name: "section",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings([]section.Setting{
						{Category: "First", Name: "1"},
						{Category: "First", Name: "2"},
						{Category: "First", Name: "3"},
						{Category: "First", Name: "4"},
						{Category: "Second", Name: "1"},
						{Category: "Third", Name: "1"},
						{Category: "Third", Name: "2"},
						{Category: "Third", Name: "3"},
						{Category: "Forth", Name: "1"},
						{Category: "Forth", Name: "2"},
					})

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
				Theme: theme.New(config.ColourAdaptive),
			}

			m := option.New(state)

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
