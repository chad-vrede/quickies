package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			name:     "basic increment",
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "no carry needed",
			digits:   []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			name:     "single digit 9",
			digits:   []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "multiple nines",
			digits:   []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "carry in middle",
			digits:   []int{1, 9, 9},
			expected: []int{2, 0, 0},
		},
		{
			name:     "single zero",
			digits:   []int{0},
			expected: []int{1},
		},
		{
			name:     "ending with 9",
			digits:   []int{1, 2, 9},
			expected: []int{1, 3, 0},
		},
		{
			name:     "large number",
			digits:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := plusOne(tt.digits)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("plusOne(%v) = %v, expected %v", tt.digits, result, tt.expected)
			}
		})
	}
}

func BenchmarkPlusOne(b *testing.B) {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration since the function might modify the slice
		digitsCopy := make([]int, len(digits))
		copy(digitsCopy, digits)
		plusOne(digitsCopy)
	}
}
