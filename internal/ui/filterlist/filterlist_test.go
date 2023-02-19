package filterlist_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/filterlist"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
)

type MockItem struct {
	title string
}

func (i MockItem) Title() string {
	return i.title
}

func (i MockItem) Description() string {
	return ""
}

func (i MockItem) FilterValue() string {
	return ""
}

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		height int
		items  []MockItem
		title  string
		model  func(m filterlist.Model) filterlist.Model
	}

	type want struct {
		model func(m filterlist.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			args: args{
				height: 10,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
					{title: "item 3"},
				},
				title: "test",
			},
		},
		{
			name: "focus",
			args: args{
				height: 10,
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					assert.Equal(t, true, m.Focused())
				},
			},
		},
		{
			name: "blur",
			args: args{
				height: 10,
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(nil))
					m.Blur()
					m, _ = filterlist.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					assert.Equal(t, false, m.Focused())
				},
			},
		},
		{
			name: "one",
			args: args{
				height: 1,
				items: []MockItem{
					{title: "item 1"},
				},
				title: "test",
			},
		},
		{
			name: "no_items",
			args: args{
				height: 1,
				items:  nil,
				title:  "test",
			},
		},
		{
			name: "no_title",
			args: args{
				height: 1,
				items:  nil,
				title:  "",
			},
		},
		{
			name: "multiple_pages",
			args: args{
				height: 1,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
					{title: "item 3"},
				},
				title: "",
			},
		},
		{
			name: "overflow_pages",
			args: args{
				height: 1,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
					{title: "item 3"},
					{title: "item 4"},
				},
				title: "",
			},
		},
		{
			name: "down",
			args: args{
				height: 2,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
				},
				title: "title",
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					return m
				},
			},
		},
		{
			name: "up",
			args: args{
				height: 2,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
				},
				title: "title",
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyUp}))
					return m
				},
			},
		},
		{
			name: "pagedown",
			args: args{
				height: 2,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
					{title: "item 3"},
					{title: "item 4"},
				},
				title: "title",
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					return m
				},
			},
		},
		{
			name: "pageup",
			args: args{
				height: 2,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
					{title: "item 3"},
					{title: "item 4"},
				},
				title: "title",
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgUp}))
					return m
				},
			},
		},
		{
			name: "pagedown_lastpage",
			args: args{
				height: 2,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
					{title: "item 3"},
					{title: "item 4"},
				},
				title: "title",
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyPgDown}))
					return m
				},
			},
		},
		{
			name: "enter",
			args: args{
				height: 2,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
				},
				title: "title",
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					title := m.SelectedItem().(MockItem).title
					assert.Equal(t, "item 1", title)
				},
			},
		},
		{
			name: "escape",
			args: args{
				height: 2,
				title:  "Prompt:",
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
				},
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEsc}))
					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					f := m.Filter()
					assert.Equal(t, "", f)
				},
			},
		},
		{
			name: "setitems",
			args: args{
				height: 2,
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
				},
				title: "title",
				model: func(m filterlist.Model) filterlist.Model {
					items := castToListItems([]MockItem{
						{title: "newitem 1"},
						{title: "newitem 2"},
					})
					m.Focus()
					m.SetItems(items)
					m, _ = filterlist.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					title := m.SelectedItem().(MockItem).title
					assert.Equal(t, "newitem 1", title)
				},
			},
		},
		{
			name: "type",
			args: args{
				height: 2,
				title:  "Prompt:",
				items: []MockItem{
					{title: "item 1"},
					{title: "item 2"},
				},
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = filterlist.ToModel(m.Update(nil))
					m, _ = filterlist.ToModel(uitest.SendString(m, "item"), nil)

					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					f := m.Filter()
					assert.Equal(t, "item", f)
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

			m := filterlist.New(castToListItems(tt.args.items), tt.args.title, tt.args.height, state)
			m, _ = filterlist.ToModel(m.Update(nil))

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

func castToListItems(m []MockItem) []list.Item {
	res := make([]list.Item, len(m))
	for i, e := range m {
		item := e
		res[i] = item
	}

	return res
}
