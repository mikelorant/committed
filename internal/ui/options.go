package ui

import (
	"github.com/mikelorant/committed/internal/ui/option/section"
)

func (m *Model) configureOptions() {
	m.models.option.SectionWidth = 30
	m.models.option.SectionHeight = 23

	m.setOptionSettings()
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
