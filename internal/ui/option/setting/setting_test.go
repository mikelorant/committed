package setting_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/option/setting"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
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
		{
			name: "radio",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "Title", Values: []string{"1", "2", "3"}},
				},
				model: func(m setting.Model) setting.Model {
					m.SelectPane("Title")

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "1", ap.Value())
				},
			},
		},
		{
			name: "radio_select",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "Title", Values: []string{"1", "2", "3"}},
				},
				model: func(m setting.Model) setting.Model {
					m.SelectPane("Title")

					ap := setting.ToRadio(m.ActivePane())
					ap.Select(2)

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					assert.Equal(t, "3", m.ActivePane().(*setting.Radio).Value())
				},
			},
		},
		{
			name: "radio_next",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "Title", Values: []string{"1", "2", "3"}},
				},
				model: func(m setting.Model) setting.Model {
					ap := setting.ToRadio(m.ActivePane())

					ap.Next()

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					assert.Equal(t, "2", setting.ToRadio(m.ActivePane()).Value())
				},
			},
		},
		{
			name: "radio_next_last",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "Title", Values: []string{"1", "2", "3"}},
				},
				model: func(m setting.Model) setting.Model {
					ap := setting.ToRadio(m.ActivePane())

					ap.Select(2)
					ap.Next()

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					assert.Equal(t, "1", setting.ToRadio(m.ActivePane()).Value())
				},
			},
		},
		{
			name: "radio_previous",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "Title", Values: []string{"1", "2", "3"}},
				},
				model: func(m setting.Model) setting.Model {
					ap := setting.ToRadio(m.ActivePane())

					ap.Select(2)
					ap.Previous()

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "2", ap.Value())
				},
			},
		},
		{
			name: "radio_previous_first",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "Title", Values: []string{"1", "2", "3"}},
				},
				model: func(m setting.Model) setting.Model {
					ap := setting.ToRadio(m.ActivePane())

					ap.Previous()

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "3", ap.Value())
				},
			},
		},
		{
			name: "radio_multiple",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "First", Values: []string{"1", "2", "3"}},
					&setting.Radio{Title: "Second", Values: []string{"4", "5", "6"}},
				},
			},
		},
		{
			name: "radio_multiple_select",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "First", Values: []string{"1", "2", "3"}},
					&setting.Radio{Title: "Second", Values: []string{"4", "5", "6"}},
				},
				model: func(m setting.Model) setting.Model {
					m.SelectPane("Second")

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "4", ap.Value())
				},
			},
		},
		{
			name: "radio_multiple_select_next",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "First", Values: []string{"1", "2", "3"}},
					&setting.Radio{Title: "Second", Values: []string{"4", "5", "6"}},
				},
				model: func(m setting.Model) setting.Model {
					m.SelectPane("Second")

					ap := setting.ToRadio(m.ActivePane())
					ap.Next()

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "5", ap.Value())
				},
			},
		},
		{
			name: "radio_invalid",
			args: args{
				panes: []setting.Paner{
					&setting.Radio{Title: "First", Values: []string{"1", "2", "3"}},
				},
				model: func(m setting.Model) setting.Model {
					m.SelectPane("Invalid")

					return m
				},
			},
			want: want{
				model: func(m setting.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "1", ap.Value())
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
