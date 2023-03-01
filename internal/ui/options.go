package ui

import (
	"github.com/mikelorant/committed/internal/ui/option/section"
	"github.com/mikelorant/committed/internal/ui/option/setting"
)

func (m *Model) configureOptions() {
	m.models.option.SectionWidth = 30
	m.models.option.SettingWidth = 41
	m.models.option.HelpWidth = 41
	m.models.option.SectionHeight = 23
	m.models.option.SettingHeight = 17
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
				Values: []string{"Above", "Below"},
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
				Values: []string{"Default", "ttyd", "kitty"},
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
