package help_test

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/ui/help"
	"github.com/mikelorant/committed/internal/ui/uitest"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	type args struct {
		content string
		model   func(m help.Model) help.Model
	}

	type want struct {
		model func(m help.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			want: want{
				model: func(m help.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "singleline",
			args: args{
				content: "line",
			},
			want: want{
				model: func(m help.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "multiline",
			args: args{
				content: "line 1\nline 2\nline 3\n",
			},
			want: want{
				model: func(m help.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "focus",
			args: args{
				model: func(m help.Model) help.Model {
					m.Focus()
					m, _ = help.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m help.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
		{
			name: "blur",
			args: args{
				model: func(m help.Model) help.Model {
					m.Focus()
					m, _ = help.ToModel(m.Update(nil))
					m.Blur()
					m, _ = help.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m help.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			help.Content = tt.args.content

			m := help.New()

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
