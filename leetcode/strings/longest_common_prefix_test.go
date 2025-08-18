package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "common prefix exists",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "no common prefix",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "single string",
			strs:     []string{"hello"},
			expected: "hello",
		},
		{
			name:     "empty array",
			strs:     []string{},
			expected: "",
		},
		{
			name:     "all same strings",
			strs:     []string{"test", "test", "test"},
			expected: "test",
		},
		{
			name:     "empty string in array",
			strs:     []string{"abc", "", "ab"},
			expected: "",
		},
		{
			name:     "one character common",
			strs:     []string{"a", "ab", "abc"},
			expected: "a",
		},
		{
			name:     "complete prefix match",
			strs:     []string{"abc", "abcdef", "abcxyz"},
			expected: "abc",
		},
		{
			name:     "no match from start",
			strs:     []string{"abc", "def", "ghi"},
			expected: "",
		},
		{
			name:     "single character strings",
			strs:     []string{"a", "a", "a"},
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.strs)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %q, expected %q", tt.strs, result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	strs := []string{"flower", "flow", "flight", "flock", "flute"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(strs)
	}
}