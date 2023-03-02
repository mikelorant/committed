package section_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/option/section"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		model func(section.Model) section.Model
	}

	type want struct {
		model    func(section.Model)
		category string
		setting  string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "first",
			want: want{
				category: "First",
				setting:  "1",
			},
		},
		{
			name: "last",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 3
					m.SetIndex = 1

					m, _ = section.ToModel(m.Update(nil))

					return m
				},
			},
			want: want{
				category: "Forth",
				setting:  "2",
			},
		},
		{
			name: "next",
			args: args{
				model: func(m section.Model) section.Model {
					m.Next()

					return m
				},
			},
			want: want{
				category: "First",
				setting:  "2",
			},
		},
		{
			name: "previous",
			args: args{
				model: func(m section.Model) section.Model {
					m.SetIndex = 3
					m.Previous()

					return m
				},
			},
			want: want{
				category: "First",
				setting:  "3",
			},
		},
		{
			name: "next_category",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 2
					m.SetIndex = 2
					m.Next()

					return m
				},
			},
			want: want{
				category: "Forth",
				setting:  "1",
			},
		},
		{
			name: "previous_category",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 3
					m.SetIndex = 0
					m.Previous()

					return m
				},
			},
			want: want{
				category: "Third",
				setting:  "3",
			},
		},
		{
			name: "next_limit",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 3
					m.SetIndex = 1
					m.Next()

					return m
				},
			},
			want: want{
				category: "Forth",
				setting:  "2",
			},
		},
		{
			name: "previous_limit",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 0
					m.SetIndex = 0
					m.Previous()

					return m
				},
			},
			want: want{
				category: "First",
				setting:  "1",
			},
		},
		{
			name: "category_only",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 1
					m.SetIndex = 0

					return m
				},
			},
			want: want{
				category: "Second",
				setting:  "",
			},
		},
		{
			name: "reset",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 2
					m.SetIndex = 2
					m.Reset()

					return m
				},
			},
			want: want{
				category: "First",
				setting:  "1",
			},
		},
		{
			name: "invalid",
			args: args{
				model: func(m section.Model) section.Model {
					m.CatIndex = 9
					m.SetIndex = 9

					return m
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			state := &commit.State{
				Theme: theme.New(theme.Default(config.ColourAdaptive)),
			}

			m := section.New(state)
			m.Settings = testSettings()

			if tt.args.model != nil {
				m = tt.args.model(m)
			}

			assert.Equal(t, tt.want.category, m.SelectedCategory())
			assert.Equal(t, tt.want.setting, m.SelectedSetting())

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}

func testSettings() []section.Setting {
	return []section.Setting{
		{Category: "First", Name: "1"},
		{Category: "First", Name: "2"},
		{Category: "First", Name: "3"},
		{Category: "First", Name: "4"},
		{Category: "Second", Name: "1"},
		{Category: "Third", Name: "1"},
		{Category: "Third", Name: "2"},
		{Category: "Third", Name: "3"},
		{Category: "Forth", Name: "1"},
		{Category: "Forth", Name: "2"},
	}
}
