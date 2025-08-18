package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "no repeating characters",
			s:        "abcabcbb",
			expected: 3,
		},
		{
			name:     "all same characters",
			s:        "bbbbb",
			expected: 1,
		},
		{
			name:     "mixed pattern",
			s:        "pwwkew",
			expected: 3,
		},
		{
			name:     "empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "single character",
			s:        "a",
			expected: 1,
		},
		{
			name:     "no repeated characters",
			s:        "abcdef",
			expected: 6,
		},
		{
			name:     "space character",
			s:        " ",
			expected: 1,
		},
		{
			name:     "alternating pattern",
			s:        "abba",
			expected: 2,
		},
		{
			name:     "special characters",
			s:        "!@#$%",
			expected: 5,
		},
		{
			name:     "numbers and letters",
			s:        "a1b2c3a1",
			expected: 6,
		},
		{
			name:     "long string with cycle",
			s:        "dvdf",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %v, expected %v", tt.s, result, tt.expected)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "abcabcbb"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s)
	}
}