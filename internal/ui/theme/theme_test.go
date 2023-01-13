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
