package fuzzy

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Item interface {
	Terms() []string
}

func Rank(term string, items []Item) []int {
	idx := make([]int, len(items))
	null := string(rune(0))

	if len(term) < 3 {
		for i := 0; i < len(items); i++ {
			idx[i] = i
		}
		return idx
	}

	var lines []string
	for i, item := range items {
		for _, t := range item.Terms() {
			line := fmt.Sprintf("%v%v%v", i, null, t)
			lines = append(lines, line)
		}
	}

	ranks := fuzzy.RankFindFold(term, lines)
	sort.Sort(ranks)

	idx = make([]int, 0)
	for _, v := range ranks {
		pos := strings.Split(v.Target, null)[0]
		i, _ := strconv.Atoi(pos)
		if !contains(idx, i) {
			//nolint:makezero
			idx = append(idx, i)
		}
	}

	return idx
}

func contains[T comparable](vs []T, val T) bool {
	for _, v := range vs {
		if v == val {
			return true
		}
	}
	return false
}
