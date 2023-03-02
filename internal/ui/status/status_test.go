package status_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/status"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
)

const (
	globalShortcuts = iota
	helpShortcuts
	optionShortcuts
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		shortcuts int
		next      string
		previous  string
	}

	type want struct{}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
		},
		{
			name: "next",
			args: args{
				next: "next",
			},
		},
		{
			name: "previous",
			args: args{
				previous: "previous",
			},
		},
		{
			name: "next_previous",
			args: args{
				next:     "next",
				previous: "previous",
			},
		},
		{
			name: "next_previous",
			args: args{
				next:     "next",
				previous: "previous",
			},
		},
		{
			name: "help",
			args: args{
				shortcuts: helpShortcuts,
			},
		},
		{
			name: "option",
			args: args{
				shortcuts: optionShortcuts,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			state := &commit.State{
				Theme: theme.New(theme.Default(config.ColourAdaptive)),
			}

			m := status.New(state)

			switch tt.args.shortcuts {
			case helpShortcuts:
				m.Shortcuts = status.HelpShortcuts()
			case optionShortcuts:
				m.Shortcuts = status.OptionShortcuts()
			default:
				m.Shortcuts = status.GlobalShortcuts(tt.args.next, tt.args.previous)
			}

			m, _ = status.ToModel(m.Update(nil))

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
