package leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "basic case",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "single zero",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "no zeros",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "zeros at beginning",
			nums:     []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "zeros at end",
			nums:     []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "alternating zeros",
			nums:     []int{0, 1, 0, 2, 0, 3},
			expected: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name:     "single non-zero",
			nums:     []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies in-place
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			moveZeroes(numsCopy)
			if !reflect.DeepEqual(numsCopy, tt.expected) {
				t.Errorf("moveZeroes(%v) = %v, expected %v", tt.nums, numsCopy, tt.expected)
			}
		})
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	nums := []int{0, 1, 0, 3, 12, 0, 0, 4, 5, 0, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		b.StartTimer()

		moveZeroes(numsCopy)
	}
}
