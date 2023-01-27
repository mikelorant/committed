package ui

import (
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/ui/theme"
)

func (m *Model) defaults(cfg config.Config) {
	m.defaultEmojiType(cfg.Commit.EmojiType)
	m.defaultFocus(cfg.View.Focus)
	m.defaultSignoff(cfg.Commit.Signoff)
	m.defaultTheme(cfg.View.Theme)
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

func (m *Model) defaultTheme(t string) {
	th := theme.New()
	th.SetTint(t)

	m.state.Theme = th
}