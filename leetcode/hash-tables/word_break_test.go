package leetcode

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "example 1 - simple break",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "example 2 - reuse words",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "example 3 - cannot break",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "single character match",
			s:        "a",
			wordDict: []string{"a"},
			expected: true,
		},
		{
			name:     "single character no match",
			s:        "a",
			wordDict: []string{"b"},
			expected: false,
		},
		{
			name:     "empty string",
			s:        "",
			wordDict: []string{"a"},
			expected: true,
		},
		{
			name:     "multiple ways to break",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			expected: true,
		},
		{
			name:     "overlapping prefixes",
			s:        "aaaaaa",
			wordDict: []string{"aaaa", "aa"},
			expected: true,
		},
		{
			name:     "complex case",
			s:        "goalspecial",
			wordDict: []string{"go", "goal", "goals", "special"},
			expected: true,
		},
		{
			name:     "prefix exists but not complete",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aa", "a"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordBreak(tt.s, tt.wordDict)
			if result != tt.expected {
				t.Errorf("wordBreak(%q, %v) = %v, expected %v", tt.s, tt.wordDict, result, tt.expected)
			}
		})
	}
}

func BenchmarkWordBreak(t *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}

	for i := 0; i < t.N; i++ {
		wordBreak(s, wordDict)
	}
}
