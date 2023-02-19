package shortcut_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/shortcut"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		keybindings []shortcut.KeyBinding
		modifiers   []shortcut.Modifier
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
			name: "left",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t", Label: "test"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignLeft, Label: "Test"},
				},
			},
		},
		{
			name: "empty_left",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t", Label: "test"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignLeft},
				},
			},
		},
		{
			name: "empty_left_top",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.AltModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignLeft},
					{Modifier: shortcut.AltModifier, Align: shortcut.AlignLeft, Label: "Test2"},
				},
			},
		},
		{
			name: "empty_left_bottom",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.AltModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.NoModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.AltModifier, Align: shortcut.AlignLeft, Label: "Test1"},
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignLeft},
				},
			},
		},
		{
			name: "multiple_same_left",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.NoModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignLeft, Label: "Test1"},
				},
			},
		},
		{
			name: "multiple_different_left",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.AltModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignLeft, Label: "Test1"},
					{Modifier: shortcut.AltModifier, Align: shortcut.AlignLeft, Label: "Test2"},
				},
			},
		},
		{
			name: "right",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t", Label: "test"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignRight, Label: "Test"},
				},
			},
		},
		{
			name: "empty_right",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t", Label: "test"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignRight},
				},
			},
		},
		{
			name: "empty_right_top",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.AltModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignRight},
					{Modifier: shortcut.AltModifier, Align: shortcut.AlignRight, Label: "Test2"},
				},
			},
		},
		{
			name: "empty_right_bottom",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.AltModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.NoModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.AltModifier, Align: shortcut.AlignRight, Label: "Test"},
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignRight},
				},
			},
		},
		{
			name: "multiple_same_right",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.NoModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignRight, Label: "Test1"},
				},
			},
		},
		{
			name: "multiple_different_right",
			args: args{
				keybindings: []shortcut.KeyBinding{
					{Modifier: shortcut.NoModifier, Key: "t1", Label: "test1"},
					{Modifier: shortcut.AltModifier, Key: "t2", Label: "test2"},
				},
				modifiers: []shortcut.Modifier{
					{Modifier: shortcut.NoModifier, Align: shortcut.AlignRight, Label: "Test1"},
					{Modifier: shortcut.AltModifier, Align: shortcut.AlignRight, Label: "Test2"},
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

			m := shortcut.New(shortcut.Shortcuts{
				Modifiers:   tt.args.modifiers,
				KeyBindings: tt.args.keybindings,
				State:       state,
			})

			m, _ = shortcut.ToModel(m.Update(nil))

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
