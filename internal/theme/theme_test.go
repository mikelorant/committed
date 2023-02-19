package theme_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			th := theme.New(tt.colour)

			var ids []string
			for i := 0; i < len(th.ListID()); i++ {
				ids = append(ids, th.ListID()[i])
			}

			assert.Equal(t, tt.ids, ids)
		})
	}
}

func TestNext(t *testing.T) {
	t.Parallel()

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
			count: 4,
			id:    "retrowave",
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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			th := theme.New(config.ColourAdaptive)

			for i := 0; i < tt.count; i++ {
				th.Next()
			}

			assert.Equal(t, tt.id, th.ID)
		})
	}
}

func TestList(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			th := theme.New(config.ColourAdaptive)

			got := th.ListID()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSet(t *testing.T) {
	t.Parallel()

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
			want: want{
				id: "builtin_dark",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			th := theme.New(config.ColourAdaptive)

			ok := th.Set(tt.id)
			if !tt.want.ok {
				assert.False(t, ok)
				assert.Equal(t, tt.want.id, th.ID)
				return
			}
			assert.True(t, ok)

			assert.Equal(t, tt.want.id, th.ID)
		})
	}
}
