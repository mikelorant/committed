package fuzzy_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/fuzzy"
	"github.com/stretchr/testify/assert"
)

type MockItem struct {
	terms []string
}

func (m MockItem) Terms() []string {
	return m.terms
}

func TestRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		inputTerm  string
		inputTerms [][]string
		want       []int
	}{
		{
			name:      "terms_one",
			inputTerm: "one",
			inputTerms: [][]string{
				{"one"},
				{"two"},
				{"three"},
			},
			want: []int{0},
		},
		{
			name:      "terms_multiple",
			inputTerm: "one",
			inputTerms: [][]string{
				{"two", "three", "four"},
				{"three", "four", "five"},
				{"one", "two", "three"},
			},
			want: []int{2},
		},
		{
			name:      "terms_multiple_matches",
			inputTerm: "one",
			inputTerms: [][]string{
				{"one", "two", "three"},
				{"two", "three", "four"},
				{"three", "one", "five"},
			},
			want: []int{0, 2},
		},
		{
			name:      "terms_multiple_fuzzy",
			inputTerm: "one",
			inputTerms: [][]string{
				{"one", "two", "three"},
				{"two", "three", "four"},
				{"three", "four", "bone"},
			},
			want: []int{0, 2},
		},
		{
			name:      "terms_match_none",
			inputTerm: "none",
			inputTerms: [][]string{
				{"one", "two", "three"},
				{"two", "three", "four"},
				{"three", "four", "five"},
			},
			want: []int{},
		},
		{
			name:      "terms_duplicate",
			inputTerm: "one",
			inputTerms: [][]string{
				{"one", "one", "one"},
				{"two", "three", "four"},
				{"three", "four", "five"},
			},
			want: []int{0},
		},
		{
			name:      "short_term",
			inputTerm: "on",
			inputTerms: [][]string{
				{"one"},
				{"two"},
				{"three"},
			},
			want: []int{0, 1, 2},
		},
		{
			name:       "empty_term",
			inputTerm:  "",
			inputTerms: [][]string{},
			want:       []int{},
		},
		{
			name:       "empty_terms",
			inputTerm:  "none",
			inputTerms: [][]string{},
			want:       []int{},
		},
		{
			name:       "nil_terms",
			inputTerm:  "none",
			inputTerms: nil,
			want:       []int{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			it := castToFuzzyItems(tt.inputTerms)
			got := fuzzy.Rank(tt.inputTerm, it)
			assert.Equal(t, tt.want, got)
		})
	}
}

func castToFuzzyItems(items [][]string) []fuzzy.Item {
	res := make([]fuzzy.Item, len(items))
	for i, it := range items {
		var item MockItem
		item.terms = it
		res[i] = item
	}

	return res
}
