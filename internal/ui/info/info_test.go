package info_test

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/ui/info"
	"github.com/mikelorant/committed/internal/ui/theme"
	"github.com/mikelorant/committed/internal/ui/uitest"
	"github.com/stretchr/testify/assert"
)

const dateTimeFormat = "Mon Jan 2 15:04:05 2006 -0700"

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		state func(c *commit.State)
		model func(m info.Model) info.Model
	}

	type want struct {
		model func(m info.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			want: want{
				model: func(m info.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "focus",
			args: args{
				model: func(m info.Model) info.Model {
					m.Focus()
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m info.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
		{
			name: "blur",
			args: args{
				model: func(m info.Model) info.Model {
					m.Focus()
					m, _ = info.ToModel(m.Update(nil))
					m.Blur()
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m info.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		{
			name: "expand",
			args: args{
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					return m
				},
			},
			want: want{
				model: func(m info.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
		{
			name: "remote",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Branch.Remote = "origin/master"
				},
			},
		},
		{
			name: "tags",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Branch.Refs = []string{"v1.0.0"}
				},
			},
		},
		{
			name: "no_users",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = nil
				},
			},
		},
		{
			name: "no_local",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Branch.Local = ""
				},
			},
		},
		{
			name: "multiple_users",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = testRepositoryUsers(2)
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					return m
				},
			},
		},
		{
			name: "multiple_users_selected",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = testRepositoryUsers(2)
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					m, _ = info.ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
			want: want{
				model: func(m info.Model) {
					assert.Equal(t, m.Author, testRepositoryUsers(2)[1])
				},
			},
		},
		{
			name: "multiple_users_filtered",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = testRepositoryUsers(2)
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					m, _ = info.ToModel(uitest.SendString(m, "example.org"), nil)
					return m
				},
			},
		},
		{
			name: "users_mixed",
			args: args{
				state: func(c *commit.State) {
					c.Config.Authors = testConfigUsers(1)
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "users_mixed_multiple",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = testRepositoryUsers(2)
					c.Config.Authors = testConfigUsers(2)
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "repository_user_only",
			args: args{
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "config_user_only",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = nil
					c.Config.Authors = testConfigUsers(1)
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "users_both_nil",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = nil
					c.Config.Authors = nil
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "users_both_empty",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = []repository.User{}
					c.Config.Authors = []repository.User{}
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "multiple_users_config_default",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = nil
					c.Config.Authors = testConfigUsers(2)
					c.Config.Authors[1].Default = true
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "multiple_users_repository_config_default",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = testRepositoryUsers(1)
					c.Config.Authors = testConfigUsers(1)
					c.Config.Authors[0].Default = true
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "multiple_users_repository_config_default_multiple",
			args: args{
				state: func(c *commit.State) {
					c.Repository.Users = testRepositoryUsers(2)
					c.Config.Authors = testConfigUsers(2)
					c.Config.Authors[0].Default = true
					c.Config.Authors[1].Default = true
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					return m
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := testState()
			if tt.args.state != nil {
				tt.args.state(&c)
			}

			m := info.New(&c)
			m.Date = time.Date(2022, time.January, 1, 1, 0, 0, 0, time.UTC).Format(dateTimeFormat)

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

func testState() commit.State {
	return commit.State{
		Placeholders: commit.Placeholders{
			Hash: "1",
		},
		Repository: repository.Description{
			Branch: repository.Branch{
				Local: "master",
			},
			Users: testRepositoryUsers(1),
			Head: repository.Head{
				Hash: "1",
				When: time.Date(2022, time.January, 1, 1, 0, 0, 0, time.UTC),
			},
		},
		Theme: theme.New(config.ColourAdaptive),
		Options: commit.Options{
			Amend: true,
		},
	}
}

func testRepositoryUsers(n int) []repository.User {
	return []repository.User{
		{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		{
			Name:  "John Doe",
			Email: "jdoe@example.org",
		},
	}[0:n]
}

func testConfigUsers(n int) []repository.User {
	return []repository.User{
		{
			Name:  "John Doe",
			Email: "jd@example.net",
		},
		{
			Name:  "John Doe",
			Email: "j@example.id",
		},
	}[0:n]
}
