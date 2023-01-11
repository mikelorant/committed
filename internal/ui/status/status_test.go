package status_test

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/ui/shortcut"
	"github.com/mikelorant/committed/internal/ui/status"
	"github.com/mikelorant/committed/internal/ui/uitest"
)

func TestModel(t *testing.T) {
	type args struct {
		shortcuts []shortcut.Shortcut
		modifiers []shortcut.Modifier
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scs := shortcut.NewShortcut(status.GlobalShortcuts(tt.args.next, tt.args.previous))
			if tt.args.modifiers != nil && tt.args.shortcuts != nil {
				scs = shortcut.NewShortcut(tt.args.modifiers, tt.args.shortcuts)
			}

			m := status.Model{
				Next:      tt.args.next,
				Previous:  tt.args.previous,
				Shortcuts: scs,
			}

			m, _ = status.ToModel(m.Update(nil))

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
