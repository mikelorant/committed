package status_test

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/ui/status"
	"github.com/mikelorant/committed/internal/ui/theme"
	"github.com/mikelorant/committed/internal/ui/uitest"
)

const (
	globalShortcuts = iota
	helpShortcuts
)

func TestModel(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &commit.State{
				Theme: theme.New(config.ColourAdaptive),
			}

			m := status.New(state)

			switch tt.args.shortcuts {
			case helpShortcuts:
				m.Shortcuts = status.HelpShortcuts()
			default:
				m.Shortcuts = status.GlobalShortcuts(tt.args.next, tt.args.previous)
			}

			m, _ = status.ToModel(m.Update(nil))

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
