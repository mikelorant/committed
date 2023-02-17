package ui

import (
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
)

func (m *Model) defaults(cfg config.Config) {
	m.defaultEmojiType(cfg.Commit.EmojiType)
	m.defaultFocus(cfg.View.Focus)
	m.defaultSignoff(cfg.Commit.Signoff)
	m.defaultTheme(cfg.View.Theme, cfg.View.Colour)
}

func (m *Model) defaultEmojiType(et config.EmojiType) {
	m.emojiType = et
}

func (m *Model) defaultFocus(focus config.Focus) {
	switch focus {
	case config.FocusAuthor:
		m.focus = authorComponent
	case config.FocusEmoji:
		m.focus = emojiComponent
	case config.FocusSummary:
		m.focus = summaryComponent
	default:
		m.focus = emojiComponent
	}
}

func (m *Model) defaultSignoff(signoff bool) {
	m.signoff = signoff
}

func (m *Model) defaultTheme(th string, clr config.Colour) {
	t := theme.New(clr)
	t.Set(th)

	m.state.Theme = t
}

func defaultAmendSave(st *commit.State) savedState {
	s := savedState{
		amend:   true,
		summary: commit.MessageToSummary(st.Repository.Head.Message),
		body:    commit.MessageToBody(st.Repository.Head.Message),
	}

	if e := commit.MessageToEmoji(st.Emojis, st.Repository.Head.Message); e.Valid {
		s.emoji = e.Emoji
	}

	return s
}
