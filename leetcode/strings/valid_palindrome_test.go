package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "valid palindrome with punctuation",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "not a palindrome",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "empty string",
			s:        "",
			expected: true,
		},
		{
			name:     "single character",
			s:        "a",
			expected: true,
		},
		{
			name:     "single non-alphanumeric",
			s:        ".",
			expected: true,
		},
		{
			name:     "simple palindrome",
			s:        "racecar",
			expected: true,
		},
		{
			name:     "palindrome with numbers",
			s:        "A1B2b1a",
			expected: true,
		},
		{
			name:     "mixed case palindrome",
			s:        "Madam",
			expected: true,
		},
		{
			name:     "punctuation only",
			s:        ".,!",
			expected: true,
		},
		{
			name:     "spaces and punctuation",
			s:        " ",
			expected: true,
		},
		{
			name:     "numbers palindrome",
			s:        "12321",
			expected: true,
		},
		{
			name:     "not palindrome with numbers",
			s:        "12345",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.s)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v, expected %v", tt.s, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	s := "A man, a plan, a canal: Panama"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
}