package search

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func FuzzySearch(content, query string) []string {
	lines := strings.Split(content, "\n")
	results := []string{}

	for _, line := range lines {
		if fuzzy.MatchFold(query, line) {
			results = append(results, line)
		}
	}

	return results
}
