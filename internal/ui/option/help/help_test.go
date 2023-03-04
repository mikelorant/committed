package help_test

import (
	"strings"
	"testing"

	"github.com/mikelorant/committed/internal/ui/option/help"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		model func(help.Model) help.Model
	}

	type want struct {
		model   func(help.Model)
		content string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "help",
			args: args{
				model: func(m help.Model) help.Model {
					m.SetContent("help")

					return m
				},
			},
		},
		{
			name: "help_long",
			args: args{
				model: func(m help.Model) help.Model {
					m.SetContent(strings.Repeat("1234567890", 10))

					return m
				},
			},
		},
		{
			name: "help_multiline",
			args: args{
				model: func(m help.Model) help.Model {
					m.SetContent(strings.Repeat("1234567890\n", 10))

					return m
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := help.New()

			if tt.args.model != nil {
				m = tt.args.model(m)
			}

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
