package body_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/body"
	"github.com/mikelorant/committed/internal/ui/uitest"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		body   string
		height int
		model  func(m body.Model) body.Model
	}

	type want struct {
		model func(m body.Model)
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
			name: "placeholder",
			args: args{
				body: "placeholder",
			},
			want: want{
				model: func(m body.Model) {
					m, _ = body.ToModel(m.Update(nil))
					assert.Equal(t, "", m.Value())
				},
			},
		},
		{
			name: "focus",
			args: args{
				model: func(m body.Model) body.Model {
					m.Focus()
					m, _ = body.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m body.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
		{
			name: "blur",
			args: args{
				model: func(m body.Model) body.Model {
					m.Focus()
					m, _ = body.ToModel(m.Update(nil))
					m.Blur()
					m, _ = body.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m body.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "body",
			args: args{
				model: func(m body.Model) body.Model {
					m.SetValue("body")
					return m
				},
			},
			want: want{
				model: func(m body.Model) {
					assert.Equal(t, "body", m.Value())
				},
			},
		},
		{
			name: "body_multiline",
			args: args{
				model: func(m body.Model) body.Model {
					m.SetValue("line 1\nline 2\nline 3\n")
					return m
				},
			},
			want: want{
				model: func(m body.Model) {
					assert.Equal(t, "line 1\nline 2\nline 3", uitest.StripString(m.Value()))
				},
			},
		},
		{
			name: "dimensions",
			args: args{
				height: 10,
				model: func(m body.Model) body.Model {
					m.Height = 3
					m.SetValue("body")
					m, _ = body.ToModel(m.Update(m))
					return m
				},
			},
			want: want{
				model: func(m body.Model) {
					assert.Equal(t, 3, m.Height)
				},
			},
		},
		{
			name: "tab",
			args: args{
				model: func(m body.Model) body.Model {
					m.Focus()
					m, _ = body.ToModel(m.Update(m))
					m, _ = body.ToModel(uitest.SendString(m, "before"), nil)
					m, _ = body.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = body.ToModel(uitest.SendString(m, "after"), nil)
					m, _ = body.ToModel(m.Update(m))
					return m
				},
			},
			want: want{
				model: func(m body.Model) {
					assert.Equal(t, "before    after", m.Value())
				},
			},
		},
		{
			name: "empty",
			want: want{
				model: func(m body.Model) {
					assert.Equal(t, "", m.Value())
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := commit.State{
				Placeholders: commit.Placeholders{
					Body: tt.args.body,
				},
				Theme: theme.New(config.ColourAdaptive),
			}

			m := body.New(&c, tt.args.height)

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
