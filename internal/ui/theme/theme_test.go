package theme_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/ui/theme"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		colour config.Colour
		ids    []string
	}{
		{
			name:   "adaptive",
			colour: config.ColourAdaptive,
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
		{
			name:   "dark",
			colour: config.ColourDark,
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
		{
			name:   "light",
			colour: config.ColourLight,
			ids: []string{
				"builtin_light",
				"builtin_solarized_light",
				"builtin_tango_light",
				"gruvbox_light",
				"tokyo_night_light",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(tt.colour)

			var ids []string
			for i := 0; i < len(th.ListTints()); i++ {
				ids = append(ids, th.ListTints()[i])
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
			id:   "builtin_light",
		},
		{
			name:  "one",
			count: 1,
			id:    "builtin_solarized_light",
		},
		{
			name:  "three",
			count: 3,
			id:    "gruvbox_light",
		},
		{
			name:  "last",
			count: 4,
			id:    "tokyo_night_light",
		},
		{
			name:  "last_plus_one",
			count: 5,
			id:    "builtin_light",
		},
		{
			name:  "last_plus_two",
			count: 6,
			id:    "builtin_solarized_light",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)

			for i := 0; i < tt.count; i++ {
				th.NextTint()
			}

			assert.Equal(t, tt.id, th.GetTint())
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
				"builtin_light",
				"builtin_solarized_light",
				"builtin_tango_light",
				"gruvbox_light",
				"tokyo_night_light",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)

			got := th.ListTints()
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
			id:   "builtin_light",
			want: "builtin_light",
		},
		{
			name: "invalid",
			id:   "invalid",
			want: "builtin_light",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)

			_ = th.SetTint(tt.id)
			got := th.GetTint()

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
			id:   "builtin_tango_light",
			want: want{
				id: "builtin_tango_light",
				ok: true,
			},
		},
		{
			name: "invalid",
			id:   "test",
			want: want{
				id: "builtin_light",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)

			ok := th.SetTint(tt.id)
			if !tt.want.ok {
				assert.False(t, ok)
				got := th.GetTint()
				assert.Equal(t, tt.want.id, got)
				return
			}
			assert.True(t, ok)

			got := th.GetTint()
			assert.Equal(t, tt.want.id, got)
		})
	}
}
