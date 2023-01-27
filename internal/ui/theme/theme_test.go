package theme_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/ui/theme"
	"github.com/stretchr/testify/assert"
)

func TestTint(t *testing.T) {
	tests := []struct {
		name string
		ids  []string
	}{
		{
			name: "ids",
			ids: []string{
				"builtin_dark",
				"dracula",
				"gruvbox_dark",
				"nord",
				"retrowave",
				"solarized_dark_higher_contrast",
				"tokyo_night",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := theme.Tint()
			reg.SetTintID(reg.TintIDs()[0])

			var ids []string
			for i := 0; i < len(reg.TintIDs()); i++ {
				ids = append(ids, reg.TintIDs()[i])
			}

			assert.Equal(t, tt.ids, ids)
		})
	}
}

func TestNextTint(t *testing.T) {
	tests := []struct {
		name  string
		count int
		id    string
	}{
		{
			name: "first",
			id:   "builtin_dark",
		},
		{
			name:  "one",
			count: 1,
			id:    "dracula",
		},
		{
			name:  "three",
			count: 3,
			id:    "nord",
		},
		{
			name:  "last",
			count: 6,
			id:    "tokyo_night",
		},
		{
			name:  "last_plus_one",
			count: 7,
			id:    "builtin_dark",
		},
		{
			name:  "last_plus_two",
			count: 8,
			id:    "dracula",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := theme.Tint()
			reg.SetTintID(reg.TintIDs()[0])

			for i := 0; i < tt.count; i++ {
				theme.NextTint()
			}

			assert.Equal(t, tt.id, reg.ID())
		})
	}
}

func TestListTints(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "default",
			want: []string{
				"builtin_dark",
				"dracula",
				"gruvbox_dark",
				"nord",
				"retrowave",
				"solarized_dark_higher_contrast",
				"tokyo_night",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := theme.Tint()
			reg.SetTintID(reg.TintIDs()[0])

			got := theme.ListTints()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTint(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{
			name: "valid",
			id:   "builtin_dark",
			want: "builtin_dark",
		},
		{
			name: "invalid",
			id:   "invalid",
			want: "builtin_dark",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := theme.Tint()
			reg.SetTintID(reg.TintIDs()[0])

			_ = theme.SetTint(tt.id)
			got := theme.GetTint()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetTint(t *testing.T) {
	type want struct {
		id string
		ok bool
	}

	tests := []struct {
		name string
		id   string
		want want
	}{
		{
			name: "valid",
			id:   "dracula",
			want: want{
				id: "dracula",
				ok: true,
			},
		},
		{
			name: "invalid",
			id:   "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := theme.Tint()
			reg.SetTintID(reg.TintIDs()[0])

			ok := theme.SetTint(tt.id)
			if !tt.want.ok {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)

			got := theme.GetTint()
			assert.Equal(t, tt.want.id, got)
		})
	}
}
