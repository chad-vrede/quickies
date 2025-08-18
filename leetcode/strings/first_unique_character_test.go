package leetcode

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "first unique at beginning",
			s:        "leetcode",
			expected: 0,
		},
		{
			name:     "first unique in middle",
			s:        "loveleetcode",
			expected: 2,
		},
		{
			name:     "no unique characters",
			s:        "aabb",
			expected: -1,
		},
		{
			name:     "single character",
			s:        "a",
			expected: 0,
		},
		{
			name:     "all same characters",
			s:        "aaaa",
			expected: -1,
		},
		{
			name:     "first unique at end",
			s:        "aabbccdd",
			expected: -1,
		},
		{
			name:     "empty string",
			s:        "",
			expected: -1,
		},
		{
			name:     "unique at last position",
			s:        "aabbz",
			expected: 4,
		},
		{
			name:     "repeated pattern",
			s:        "abccba",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := firstUniqChar(tt.s)
			if result != tt.expected {
				t.Errorf("firstUniqChar(%q) = %v, expected %v", tt.s, result, tt.expected)
			}
		})
	}
}

func BenchmarkFirstUniqChar(b *testing.B) {
	s := "loveleetcode"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		firstUniqChar(s)
	}
}
