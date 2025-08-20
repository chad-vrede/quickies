package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort the result for comparison since order doesn't matter
func sortGroupedAnagrams(groups [][]string) [][]string {
	// Sort each group internally
	for _, group := range groups {
		sort.Strings(group)
	}
	// Sort groups by their first element for consistent comparison
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 && len(groups[j]) == 0 {
			return false
		}
		if len(groups[i]) == 0 {
			return true
		}
		if len(groups[j]) == 0 {
			return false
		}
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name: "basic anagrams",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"ate", "eat", "tea"},
				{"bat"},
				{"nat", "tan"},
			},
		},
		{
			name:     "empty string",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "single character",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name: "no anagrams",
			strs: []string{"abc", "def", "ghi"},
			expected: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
		{
			name: "all same characters",
			strs: []string{"aab", "aba", "baa"},
			expected: [][]string{
				{"aab", "aba", "baa"},
			},
		},
		{
			name: "mixed lengths",
			strs: []string{"a", "aa", "aaa", "b", "bb", "bbb"},
			expected: [][]string{
				{"a"},
				{"aa"},
				{"aaa"},
				{"b"},
				{"bb"},
				{"bbb"},
			},
		},
		{
			name: "duplicate strings",
			strs: []string{"abc", "bca", "abc", "cab"},
			expected: [][]string{
				{"abc", "abc", "bca", "cab"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.strs)

			// Sort both result and expected for comparison
			sortedResult := sortGroupedAnagrams(result)
			sortedExpected := sortGroupedAnagrams(tt.expected)

			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("groupAnagrams(%v) = %v, expected %v", tt.strs, sortedResult, sortedExpected)
			}
		})
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "tab", "arc", "car", "race", "care"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams(strs)
	}
}
