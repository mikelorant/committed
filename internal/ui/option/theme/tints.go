package theme

import (
	"github.com/mikelorant/committed/internal/fuzzy"

	"github.com/charmbracelet/bubbles/list"
	tint "github.com/lrstanley/bubbletint"
)

type listItem struct {
	tint tint.Tint
}

type fuzzyItem struct {
	tint tint.Tint
}

func (i listItem) Title() string {
	return i.tint.DisplayName()
}

func (i listItem) Description() string {
	return i.tint.ID()
}

func (i listItem) FilterValue() string {
	return i.tint.DisplayName()
}

func (i fuzzyItem) Terms() []string {
	return []string{
		i.tint.DisplayName(),
		i.tint.ID(),
	}
}

func castToListItems(tints []tint.Tint) []list.Item {
	res := make([]list.Item, len(tints))
	for i, t := range tints {
		var item listItem
		item.tint = t
		res[i] = item
	}

	return res
}

func castToFuzzyItems(tints []tint.Tint) []fuzzy.Item {
	res := make([]fuzzy.Item, len(tints))
	for i, t := range tints {
		var item fuzzyItem
		item.tint = t
		res[i] = item
	}

	return res
}
