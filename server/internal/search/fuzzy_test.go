package search

import (
	"testing"
)

func TestFuzzySearch(t *testing.T) {
	content := `apple
banana
grape
orange
pineapple`

	tests := []struct {
		query    string
		expected []string
	}{
		{
			query:    "apple",
			expected: []string{"apple", "pineapple"},
		},
		{
			query:    "banana",
			expected: []string{"banana"},
		},
		{
			query:    "berry",
			expected: []string{},
		},
		{
			query:    "grape",
			expected: []string{"grape"},
		},
		{
			query:    "Orange",
			expected: []string{"orange"},
		},
	}

	for _, test := range tests {
		results := FuzzySearch(content, test.query)
		if len(results) != len(test.expected) {
			t.Errorf("For query '%s', expected %v but got %v", test.query, test.expected, results)
			continue
		}

		for i, result := range results {
			if result != test.expected[i] {
				t.Errorf("For query '%s', expected %v but got %v", test.query, test.expected, results)
				break
			}
		}
	}
}
