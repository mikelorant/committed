package ui

import (
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/option/section"
	"github.com/mikelorant/committed/internal/ui/option/setting"
)

func (m *Model) configureOptions() {
	m.models.option.SectionWidth = 30
	m.models.option.SettingWidth = 41
	m.models.option.HelpWidth = 41
	m.models.option.ThemeWidth = 44
	m.models.option.SectionHeight = 23
	m.models.option.SettingHeight = 17
	m.models.option.ThemeHeight = 18
	m.models.option.HelpHeight = 3

	m.setOptionSettings()
	m.setOptionPaneSets()

	m.models.option.SetHelp("Help text for settings.")
}

func (m *Model) setOptionSettings() {
	m.models.option.SetSettings([]section.Setting{
		{Category: "General", Name: "Focus"},
		{Category: "General", Name: "Emoji Selector"},
		{Category: "General", Name: "Emoji Set"},
		{Category: "General", Name: "Ignore Global Author"},
		{Category: "Theme", Name: "Theme"},
		{Category: "Visual", Name: "Colour"},
		{Category: "Visual", Name: "Compatibility"},
		{Category: "Visual", Name: "Highlight Active"},
		{Category: "Commit", Name: "Emoji Type"},
		{Category: "Commit", Name: "Sign-off"},
	})
}

func (m *Model) setOptionPaneSets() {
	m.GeneralPaneSet()
	m.VisualPaneSet()
	m.ThemePaneSet()
	m.CommitPaneSet()
}

func (m *Model) GeneralPaneSet() {
	cfg := m.state.Config

	m.models.option.AddPaneSet("General",
		[]setting.Paner{
			&setting.Radio{
				Title:  "Focus",
				Values: []string{"Author", "Emoji", "Summary"},
				Index:  cfg.View.Focus.Index() - 1,
			},
			&setting.Radio{
				Title:  "Emoji Selector",
				Values: []string{"Below", "Above"},
				Index:  cfg.View.EmojiSelector.Index() - 1,
			},
			&setting.Radio{
				Title:  "Emoji Set",
				Values: []string{"Gitmoji", "Devmoji", "Emoji-Log"},
				Index:  cfg.View.EmojiSet.Index() - 1,
			},
			&setting.Toggle{
				Title:  "Ignore Global Author",
				Enable: bool(cfg.View.IgnoreGlobalAuthor),
			},
		},
	)
}

func (m *Model) VisualPaneSet() {
	cfg := m.state.Config

	m.models.option.AddPaneSet("Visual",
		[]setting.Paner{
			&setting.Radio{
				Title:  "Colour",
				Values: []string{"Adaptive", "Dark", "Light"},
				Index:  cfg.View.Colour.Index() - 1,
			},
			&setting.Radio{
				Title:  "Compatibility",
				Values: []string{"Unicode 14", "Unicode 9"},
				Index:  cfg.View.Compatibility.Index() - 1,
			},
			&setting.Toggle{
				Title:  "Highlight Active",
				Enable: bool(cfg.View.HighlightActive),
			},
		},
	)
}

func (m *Model) ThemePaneSet() {
	m.models.option.AddPaneSet("Theme",
		[]setting.Paner{
			&setting.Noop{},
		},
	)
}

func (m *Model) CommitPaneSet() {
	cfg := m.state.Config

	m.models.option.AddPaneSet("Commit",
		[]setting.Paner{
			&setting.Radio{
				Title:  "Emoji Type",
				Values: []string{"Shortcode", "Character"},
				Index:  cfg.Commit.EmojiType.Index() - 1,
			},
			&setting.Toggle{
				Title:  "Sign-off",
				Enable: bool(cfg.Commit.Signoff),
			},
		},
	)
}

func ToConfig(cfg config.Config, ps map[string][]setting.Paner, th theme.Theme) config.Config {
	view := config.View{
		Focus:              config.Focus(ps["General"][0].(*setting.Radio).Index) + 1,
		EmojiSelector:      config.EmojiSelector(ps["General"][1].(*setting.Radio).Index) + 1,
		EmojiSet:           config.EmojiSet(ps["General"][2].(*setting.Radio).Index) + 1,
		IgnoreGlobalAuthor: ps["General"][3].(*setting.Toggle).Enable,
		Colour:             config.Colour(ps["Visual"][0].(*setting.Radio).Index) + 1,
		Compatibility:      config.Compatibility(ps["Visual"][1].(*setting.Radio).Index) + 1,
		HighlightActive:    ps["Visual"][2].(*setting.Toggle).Enable,
		Theme:              th.ID,
	}

	commit := config.Commit{
		EmojiType: config.EmojiType(ps["Commit"][0].(*setting.Radio).Index) + 1,
		Signoff:   ps["Commit"][1].(*setting.Toggle).Enable,
	}

	return config.Config{
		View:    view,
		Commit:  commit,
		Authors: cfg.Authors,
	}
}
