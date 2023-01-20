package info_test

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/ui/info"
	"github.com/mikelorant/committed/internal/ui/uitest"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	type args struct {
		config func(c *commit.Config)
		model  func(m info.Model) info.Model
	}

	type want struct {
		state func(m info.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			want: want{
				state: func(m info.Model) {
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
				state: func(m info.Model) {
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
				state: func(m info.Model) {
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
				state: func(m info.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
		{
			name: "remote",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Branch.Remote = "origin/master"
				},
			},
		},
		{
			name: "tags",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Branch.Refs = []string{"v1.0.0"}
				},
			},
		},
		{
			name: "no_users",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Users = nil
				},
			},
		},
		{
			name: "no_local",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Branch.Local = ""
				},
			},
		},
		{
			name: "multiple_users",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Users = testConfigUsers(2)
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
				config: func(c *commit.Config) {
					c.Repository.Users = testConfigUsers(2)
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
				state: func(m info.Model) {
					assert.Equal(t, m.Author, testConfigUsers(2)[1])
				},
			},
		},
		{
			name: "multiple_users_filtered",
			args: args{
				config: func(c *commit.Config) {
					c.Repository.Users = []repository.User{
						testConfigUsers(3)[0],
						testConfigUsers(3)[2],
					}
				},
				model: func(m info.Model) info.Model {
					m.Focus()
					m.Expand = true
					m, _ = info.ToModel(m.Update(nil))
					m, _ = info.ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testConfig()
			if tt.args.config != nil {
				tt.args.config(&c)
			}

			m := info.New(&c)

			if tt.args.model != nil {
				m = tt.args.model(m)
			}

			if tt.want.state != nil {
				tt.want.state(m)
			}

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}

func testConfig() commit.Config {
	return commit.Config{
		Placeholders: commit.Placeholders{
			Hash: "1",
		},
		Repository: repository.Description{
			Branch: repository.Branch{
				Local: "master",
			},
			Users: testConfigUsers(1),
			Head: repository.Head{
				Hash: "1",
				When: time.Date(2022, time.January, 1, 1, 0, 0, 0, time.UTC),
			},
		},
		Amend: true,
	}
}

func testConfigUsers(n int) []repository.User {
	return []repository.User{
		{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		{
			Name:  "John Doe",
			Email: "jdoe@example.org",
		},
		{
			Name:  "John Doe",
			Email: "jdoe@test",
		},
	}[0:n]
}
