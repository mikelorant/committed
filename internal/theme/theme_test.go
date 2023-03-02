package theme_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/theme/themetest"

	tint "github.com/lrstanley/bubbletint"
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
			ids:    testIDs(15)[0:5],
		},
		{
			name:   "dark",
			colour: config.ColourDark,
			ids:    testIDs(15)[5:10],
		},
		{
			name:   "light",
			colour: config.ColourLight,
			ids:    testIDs(15)[10:15],
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var d tint.Tint
			var ds []tint.Tint

			switch tt.colour {
			case config.ColourAdaptive:
				d = themetest.NewStubTints(15)[0]
				ds = themetest.NewStubTints(15)[0:5]
			case config.ColourDark:
				d = themetest.NewStubTints(15)[5]
				ds = themetest.NewStubTints(15)[5:10]
			case config.ColourLight:
				d = themetest.NewStubTints(15)[10]
				ds = themetest.NewStubTints(15)[10:15]
			}

			tints := theme.Tint{
				Default:  d,
				Defaults: ds,
			}

			th := theme.New(tints)

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
			id:   "id0",
		},
		{
			name:  "one",
			count: 1,
			id:    "id1",
		},
		{
			name:  "three",
			count: 2,
			id:    "id2",
		},
		{
			name:  "last",
			count: 4,
			id:    "id4",
		},
		{
			name:  "last_plus_one",
			count: 5,
			id:    "id0",
		},
		{
			name:  "last_plus_two",
			count: 6,
			id:    "id1",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tints := theme.Tint{
				Default:  themetest.NewStubTints(5)[0],
				Defaults: themetest.NewStubTints(5),
			}

			th := theme.New(tints)

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
			want: []string{"id0", "id1", "id2", "id3", "id4"},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tints := theme.Tint{
				Default:  themetest.NewStubTints(5)[0],
				Defaults: themetest.NewStubTints(5),
			}

			th := theme.New(tints)

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
			id:   "id1",
			want: want{
				id: "id1",
				ok: true,
			},
		},
		{
			name: "invalid",
			id:   "test",
			want: want{
				id: "id0",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tints := theme.Tint{
				Default:  themetest.NewStubTints(5)[0],
				Defaults: themetest.NewStubTints(5),
			}

			th := theme.New(tints)

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

func testIDs(n int) []string {
	tints := themetest.NewStubTints(n)

	ids := make([]string, n)
	for idx, tint := range tints {
		ids[idx] = tint.ID()
	}

	return ids
}
