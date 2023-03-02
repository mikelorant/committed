package option_test

import (
	"fmt"
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/option"
	"github.com/mikelorant/committed/internal/ui/option/section"
	"github.com/mikelorant/committed/internal/ui/option/setting"
	"github.com/mikelorant/committed/internal/ui/uitest"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		model func(m option.Model) option.Model
	}

	type want struct {
		model func(m option.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
		},

		// Section
		{
			name: "section",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings(testSections())

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "A", m.Category())
					assert.Equal(t, "1", m.Setting())
				},
			},
		},
		{
			name: "section_category_only",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings([]section.Setting{
						{Category: "First"},
					})

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "First", m.Category())
					assert.Equal(t, "", m.Setting())
				},
			},
		},
		{
			name: "section_set",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings(testSections())

					m.SectionIndex(2, 1)

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "C", m.Category())
					assert.Equal(t, "2", m.Setting())
				},
			},
		},
		{
			name: "section_set_category_only",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings(testSections())

					m.SectionIndex(1, 0)

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "B", m.Category())
					assert.Equal(t, "", m.Setting())
				},
			},
		},

		// Setting
		{
			name: "setting",
			args: args{
				model: func(m option.Model) option.Model {
					m.AddPaneSet("Test", testSettings(3, 3))

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "0.0", ap.Value())
				},
			},
		},
		{
			name: "setting_next",
			args: args{
				model: func(m option.Model) option.Model {
					m.AddPaneSet("Test", testSettings(3, 3))

					ap := setting.ToRadio(m.ActivePane())
					ap.Next()

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "0.1", ap.Value())
				},
			},
		},
		{
			name: "setting_select",
			args: args{
				model: func(m option.Model) option.Model {
					m.AddPaneSet("Test", testSettings(3, 3))

					m.SelectPane("B")

					ap := setting.ToRadio(m.ActivePane())

					ap.Next()

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					ap := setting.ToRadio(m.ActivePane())

					assert.Equal(t, "1.1", ap.Value())
				},
			},
		},
		{
			name: "setting_previous",
			args: args{
				model: func(m option.Model) option.Model {
					m.AddPaneSet("Test", testSettings(3, 3))

					m.SelectPane("C")
					setting.ToRadio(m.ActivePane()).Select(2)
					setting.ToRadio(m.ActivePane()).Previous()

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "2.1", setting.ToRadio(m.ActivePane()).Value())
				},
			},
		},

		// Help
		{
			name: "help",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetHelp("help")

					return m
				},
			},
		},

		// Theme
		{
			name: "theme",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings([]section.Setting{
						{Category: "Theme"},
					})

					m.SectionWidth = 20
					m.SectionHeight = 23
					m.ThemeWidth = 50

					m, _ = option.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRight}))
					m, _ = option.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					m, _ = option.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))

					return m
				},
			},
		},

		// Combined
		{
			name: "combined",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings(testSections())

					m.AddPaneSet("Test", testSettings(3, 3))

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "A", m.Category())
					assert.Equal(t, "1", m.Setting())
					assert.Equal(t, "0.0", setting.ToRadio(m.ActivePane()).Value())
				},
			},
		},
		{
			name: "combined_modified",
			args: args{
				model: func(m option.Model) option.Model {
					m.SetSettings(testSections())
					m.SectionIndex(2, 1)

					m.AddPaneSet("Test", testSettings(3, 3))

					m.SelectPane("C")
					setting.ToRadio(m.ActivePane()).Select(2)
					m.SetHelp("help")

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "C", m.Category())
					assert.Equal(t, "2", m.Setting())
					assert.Equal(t, "2.2", setting.ToRadio(m.ActivePane()).Value())
				},
			},
		},
		{
			name: "combined_modified_width_height",
			args: args{
				model: func(m option.Model) option.Model {
					m.SectionWidth = 20
					m.SettingWidth = 40
					m.HelpWidth = 40
					m.SectionHeight = 20
					m.SettingHeight = 14
					m.HelpHeight = 3

					m.SetSettings(testSections())
					m.SectionIndex(2, 1)
					m.SetHelp("help")

					panes := testSettings(2, 3)

					tog := setting.Toggle{
						Title: "Toggle",
					}

					panes = append(panes, &tog)

					m.AddPaneSet("Test", panes)

					m.SelectPane("B")
					ap := setting.ToRadio(m.ActivePane())
					ap.Next()
					ap.Next()

					m.SelectPane("Toggle")
					setting.ToToggle(m.ActivePane()).Enable = true

					m.SelectPane("B")

					return m
				},
			},
			want: want{
				model: func(m option.Model) {
					assert.Equal(t, "C", m.Category())
					assert.Equal(t, "2", m.Setting())
					assert.Equal(t, "1.2", setting.ToRadio(m.ActivePane()).Value())
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

			m := option.New(state)

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

func testSections() []section.Setting {
	return []section.Setting{
		{Category: "A", Name: "1"},
		{Category: "A", Name: "2"},
		{Category: "A", Name: "3"},
		{Category: "B", Name: "1"},
		{Category: "C", Name: "1"},
		{Category: "C", Name: "2"},
		{Category: "C", Name: "3"},
		{Category: "D", Name: "1"},
		{Category: "D", Name: "2"},
	}
}

func testSettings(count int, items int) []setting.Paner {
	title := 'A'

	var panes []setting.Paner

	for i := 0; i < count; i++ {
		var vals []string
		for j := 0; j < items; j++ {
			vals = append(vals, fmt.Sprintf("%v.%v", i, j))
		}

		r := setting.Radio{
			Title:  string(title),
			Values: vals,
		}

		panes = append(panes, &r)

		title++
	}

	return panes
}
