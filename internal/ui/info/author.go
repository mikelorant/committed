package info

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/fuzzy"
)

type listItem struct {
	author commit.Author
}

type fuzzyItem struct {
	author commit.Author
}

func (i listItem) Title() string {
	return fmt.Sprintf("%s <%s>", i.author.Name, i.author.Email)
}

func (i listItem) Description() string {
	return i.author.Name
}

func (i listItem) FilterValue() string {
	return i.author.Name
}

func (i fuzzyItem) Terms() []string {
	return []string{
		i.author.Name,
		i.author.Email,
	}
}

func castToListItems(authors []commit.Author) []list.Item {
	res := make([]list.Item, len(authors))
	for i, a := range authors {
		var item listItem
		item.author = a
		res[i] = item
	}

	return res
}

func castToFuzzyItems(authors []commit.Author) []fuzzy.Item {
	res := make([]fuzzy.Item, len(authors))
	for i, e := range authors {
		var item fuzzyItem
		item.author = e
		res[i] = item
	}

	return res
}
