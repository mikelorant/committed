package header

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/mattn/go-runewidth"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/fuzzy"
)

type listItem struct {
	emoji emoji.Emoji
}

type fuzzyItem struct {
	emoji emoji.Emoji
}

func (i listItem) Title() string {
	var space string
	if runewidth.StringWidth(i.emoji.Character) == 1 {
		space = " "
	}

	return fmt.Sprintf("%s%s - %s", i.emoji.Character, space, i.emoji.Description)
}

func (i listItem) Description() string {
	return i.emoji.Name
}

func (i listItem) FilterValue() string {
	return i.emoji.Name
}

func (i fuzzyItem) Terms() []string {
	return []string{
		i.emoji.Description,
	}
}

func castToListItems(emojis []emoji.Emoji) []list.Item {
	res := make([]list.Item, len(emojis))
	for i, e := range emojis {
		var item listItem
		item.emoji = e
		res[i] = item
	}

	return res
}

func castToFuzzyItems(emojis []emoji.Emoji) []fuzzy.Item {
	res := make([]fuzzy.Item, len(emojis))
	for i, e := range emojis {
		var item fuzzyItem
		item.emoji = e
		res[i] = item
	}

	return res
}
