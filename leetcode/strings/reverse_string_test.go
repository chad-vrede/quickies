package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		s        []byte
		expected []byte
	}{
		{
			name:     "basic string",
			s:        []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:     "even length",
			s:        []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name:     "single character",
			s:        []byte{'a'},
			expected: []byte{'a'},
		},
		{
			name:     "two characters",
			s:        []byte{'a', 'b'},
			expected: []byte{'b', 'a'},
		},
		{
			name:     "empty string",
			s:        []byte{},
			expected: []byte{},
		},
		{
			name:     "palindrome",
			s:        []byte{'a', 'b', 'a'},
			expected: []byte{'a', 'b', 'a'},
		},
		{
			name:     "numbers",
			s:        []byte{'1', '2', '3', '4'},
			expected: []byte{'4', '3', '2', '1'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies in-place
			sCopy := make([]byte, len(tt.s))
			copy(sCopy, tt.s)
			
			reverseString(sCopy)
			if !reflect.DeepEqual(sCopy, tt.expected) {
				t.Errorf("reverseString(%v) = %v, expected %v", tt.s, sCopy, tt.expected)
			}
		})
	}
}

func BenchmarkReverseString(b *testing.B) {
	s := []byte{'h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd'}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		sCopy := make([]byte, len(s))
		copy(sCopy, s)
		b.StartTimer()
		
		reverseString(sCopy)
	}
}