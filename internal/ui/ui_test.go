package ui_test

import (
	"fmt"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/ui"
	"github.com/mikelorant/committed/internal/ui/theme"
	"github.com/mikelorant/committed/internal/ui/uitest"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		state func(*commit.State)
		model func(ui.Model) ui.Model
	}

	type want struct {
		model func(ui.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
		},
		{
			name: "alt+1",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					return m
				},
			},
		},
		{
			name: "alt+1_twice",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					return m
				},
			},
		},
		{
			name: "alt+2",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}, Alt: true}))
					return m
				},
			},
		},
		{
			name: "alt+2_twice",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}, Alt: true}))
					m, _ = ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "alt+3",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "alt+3_twice",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "alt+4",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "alt+4_twice",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "enter_author",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
		},
		{
			name: "enter_emoji",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					return m
				},
			},
		},
		{
			name: "enter_summary",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "enter_body",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "alt+s",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}, Alt: true}))
					return m
				},
			},
		},
		{
			name: "alt+s_twice",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}, Alt: true}))
					return m
				},
			},
		},
		{
			name: "alt+s_change_author",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}, Alt: true}))
					return m
				},
			},
		},
		{
			name: "alt+t",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'t'}, Alt: true}))
					return m
				},
			},
		},
		{
			name: "ctrl+h",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyCtrlH}))
					return m
				},
			},
		},
		{
			name: "ctrl+h_twice",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyCtrlH}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyCtrlH}))
					return m
				},
			},
		},
		{
			name: "ctrl+h_twice_body",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyCtrlH}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyCtrlH}))
					return m
				},
			},
		},
		{
			name: "escape_help",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyCtrlH}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEscape}))
					return m
				},
			},
		},
		{
			name: "tab_author",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					return m
				},
			},
		},
		{
			name: "tab_emoji",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "tab_summary",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "tab_body",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "shift_tab_author",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyShiftTab}))
					return m
				},
			},
		},
		{
			name: "shift_tab_emoji",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyShiftTab}))
					return m
				},
			},
		},
		{
			name: "shift_tab_summary",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyShiftTab}))
					return m
				},
			},
		},
		{
			name: "shift_tab_body",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyShiftTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "ctrl+c",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					_, cmd := ToModel(m.Update(tea.KeyMsg{Type: tea.KeyCtrlC}))
					assert.Equal(t, "tea.quitMsg", fmt.Sprintf("%T", cmd()))
				},
			},
		},
		{
			name: "alt+enter_summary",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))

					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					req := commit.Request{
						Apply:   true,
						Summary: "test",
						Author: repository.User{
							Name:  "John Doe",
							Email: "john.doe@example.com",
						},
						Amend: true,
					}

					assert.Equal(t, &req, m.Request)
				},
			},
		},
		{
			name: "alt+enter_summary_emoji",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))

					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					req := commit.Request{
						Apply:   true,
						Emoji:   ":art:",
						Summary: "test",
						Author: repository.User{
							Name:  "John Doe",
							Email: "john.doe@example.com",
						},
						Amend: true,
					}

					assert.Equal(t, &req, m.Request)
				},
			},
		},
		{
			name: "alt+enter_summary_body",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))

					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					req := commit.Request{
						Apply:   true,
						Summary: "test",
						Body:    "test",
						RawBody: "test",
						Author: repository.User{
							Name:  "John Doe",
							Email: "john.doe@example.com",
						},
						Amend: true,
					}

					assert.Equal(t, &req, m.Request)
				},
			},
		},
		{
			name: "alt+enter_summary_footer",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))

					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					req := commit.Request{
						Apply:   true,
						Summary: "test",
						Footer:  "Signed-off-by: John Doe <john.doe@example.com>",
						Author: repository.User{
							Name:  "John Doe",
							Email: "john.doe@example.com",
						},
						Amend: true,
					}

					assert.Equal(t, &req, m.Request)
				},
			},
		},
		{
			name: "alt+enter_summary_author",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyTab}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyDown}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))

					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					req := commit.Request{
						Apply:   true,
						Summary: "test",
						Author: repository.User{
							Name:  "John Doe",
							Email: "jdoe@example.org",
						},
						Amend: true,
					}

					assert.Equal(t, &req, m.Request)
				},
			},
		},
		{
			name: "alt+enter_invalid",
			args: args{
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}, Alt: true}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))
					return m
				},
			},
		},
		{
			name: "config_author",
			args: args{
				state: func(s *commit.State) {
					s.Config.View.Focus = config.FocusAuthor
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "config_emoji",
			args: args{
				state: func(s *commit.State) {
					s.Config.View.Focus = config.FocusEmoji
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "config_summary",
			args: args{
				state: func(s *commit.State) {
					s.Config.View.Focus = config.FocusSummary
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					return m
				},
			},
		},
		{
			name: "config_signoff_off",
			args: args{
				state: func(s *commit.State) {
					s.Config.Commit.Signoff = false
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "config_signoff_on",
			args: args{
				state: func(s *commit.State) {
					s.Config.Commit.Signoff = true
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					return m
				},
			},
		},
		{
			name: "config_emoji_type_shortcode",
			args: args{
				state: func(s *commit.State) {
					s.Config.Commit.EmojiType = config.EmojiTypeShortcode
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))

					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					req := commit.Request{
						Apply:   true,
						Emoji:   ":art:",
						Summary: "test",
						Author: repository.User{
							Name:  "John Doe",
							Email: "john.doe@example.com",
						},
						Amend: true,
					}

					assert.Equal(t, &req, m.Request)
				},
			},
		},
		{
			name: "config_emoji_type_character",
			args: args{
				state: func(s *commit.State) {
					s.Config.Commit.EmojiType = config.EmojiTypeCharacter
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter, Alt: true}))

					return m
				},
			},
			want: want{
				model: func(m ui.Model) {
					req := commit.Request{
						Apply:   true,
						Emoji:   "🎨",
						Summary: "test",
						Author: repository.User{
							Name:  "John Doe",
							Email: "john.doe@example.com",
						},
						Amend: true,
					}

					assert.Equal(t, &req, m.Request)
				},
			},
		},
		{
			name: "amend_empty",
			args: args{
				state: func(s *commit.State) {
					s.Repository.Head.Message = ":art: summary\n\nbody\n"
					s.Options.Amend = false
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}, Alt: true}))

					return m
				},
			},
		},
		{
			name: "amend_existing",
			args: args{
				state: func(s *commit.State) {
					s.Repository.Head.Message = ":art: summary\n\nbody\n"
					s.Options.Amend = false
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "test"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}, Alt: true}))

					return m
				},
			},
		},
		{
			name: "snapshot_load_from_new_to_amend",
			args: args{
				state: func(s *commit.State) {
					s.Snapshot.Summary = "amend"
					s.Snapshot.Amend = true
					s.Options.Amend = false
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "new"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}, Alt: true}))
					return m
				},
			},
			want: want{},
		},
		{
			name: "snapshot_load_from_new_to_amend_to_new",
			args: args{
				state: func(s *commit.State) {
					s.Snapshot.Summary = "amend"
					s.Snapshot.Amend = true
					s.Options.Amend = false
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyEnter}))
					m, _ = ToModel(uitest.SendString(m, "new"), nil)
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}, Alt: true}))
					return m
				},
			},
			want: want{},
		},
		{
			name: "snapshot_load_from_amend_to_new",
			args: args{
				state: func(s *commit.State) {
					s.Repository.Head.Message = "summary"
					s.Snapshot.Amend = true
					s.Options.Amend = true
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}, Alt: true}))

					return m
				},
			},
		},
		{
			name: "snapshot_load_from_amend_to_new_to_amend",
			args: args{
				state: func(s *commit.State) {
					s.Repository.Head.Message = "amend"
					s.Snapshot.Amend = true
					s.Options.Amend = true
				},
				model: func(m ui.Model) ui.Model {
					m, _ = ToModel(m.Update(nil))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}, Alt: true}))
					m, _ = ToModel(m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}, Alt: true}))

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

			m := ui.New()
			m.Date = time.Date(2022, time.January, 1, 1, 0, 0, 0, time.UTC)
			m.Configure(&c)

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
			Summary: "placeholder",
			Body:    "placeholder",
		},
		Repository: repository.Description{
			Branch: repository.Branch{
				Local: "master",
			},
			Users: []repository.User{
				{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				},
				{
					Name:  "John Doe",
					Email: "jdoe@example.org",
				},
			},
			Head: repository.Head{
				Hash: "1",
				When: time.Date(2022, time.January, 1, 1, 0, 0, 0, time.UTC),
			},
		},
		Emojis: &emoji.Set{
			Emojis: []emoji.Emoji{
				{
					Character:   "🎨",
					Description: "Improve structure / format of the code.",
					Shortcode:   ":art:",
				},
			},
		},
		Theme: theme.New(config.ColourAdaptive),
		Options: commit.Options{
			Amend: true,
		},
	}
}

func ToModel(m tea.Model, c tea.Cmd) (ui.Model, tea.Cmd) {
	return m.(ui.Model), c
}
