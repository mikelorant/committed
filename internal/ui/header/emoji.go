package header

import (
	"fmt"
	"strings"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/fuzzy"

	"github.com/charmbracelet/bubbles/list"
	"github.com/rivo/uniseg"
)

type listItem struct {
	emoji         emoji.Emoji
	compatibility config.Compatibility
}

type fuzzyItem struct {
	emoji emoji.Emoji
}

func (i listItem) Title() string {
	const maxEmojiWidth = 2

	padLen := maxEmojiWidth - uniseg.StringWidth(i.emoji.Character)
	padding := strings.Repeat(" ", padLen)

	return fmt.Sprintf("%s%s - %s", i.emoji.Character, padding, i.emoji.Description)
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
		i.emoji.Shortcode,
	}
}

func WithCompatibility(c config.Compatibility) func(*listItem) {
	return func(i *listItem) {
		i.compatibility = c
	}
}

func castToListItems(emojis []emoji.Emoji, opts ...func(*listItem)) []list.Item {
	res := make([]list.Item, len(emojis))
	for i, e := range emojis {
		var item listItem
		item.emoji = e
		for _, o := range opts {
			o(&item)
		}
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
