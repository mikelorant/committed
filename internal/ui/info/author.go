package info

import (
	"fmt"

	"github.com/mikelorant/committed/internal/fuzzy"
	"github.com/mikelorant/committed/internal/repository"

	"github.com/charmbracelet/bubbles/list"
)

type listItem struct {
	author repository.User
}

type fuzzyItem struct {
	author repository.User
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

func castToListItems(authors []repository.User) []list.Item {
	res := make([]list.Item, len(authors))
	for i, a := range authors {
		var item listItem
		item.author = a
		res[i] = item
	}

	return res
}

func castToFuzzyItems(authors []repository.User) []fuzzy.Item {
	res := make([]fuzzy.Item, len(authors))
	for i, e := range authors {
		var item fuzzyItem
		item.author = e
		res[i] = item
	}

	return res
}
