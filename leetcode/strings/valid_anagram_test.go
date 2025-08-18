package leetcode

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "valid anagram",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "not an anagram",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "different lengths",
			s:        "hello",
			t:        "world",
			expected: false,
		},
		{
			name:     "empty strings",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "single characters same",
			s:        "a",
			t:        "a",
			expected: true,
		},
		{
			name:     "single characters different",
			s:        "a",
			t:        "b",
			expected: false,
		},
		{
			name:     "repeated characters",
			s:        "aab",
			t:        "aba",
			expected: true,
		},
		{
			name:     "different character counts",
			s:        "aab",
			t:        "aaa",
			expected: false,
		},
		{
			name:     "unicode characters",
			s:        "café",
			t:        "face",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("isAnagram(%q, %q) = %v, expected %v", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	s := "anagram"
	t := "nagaram"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}
