package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		expectedLength int
		expectedArray  []int
	}{
		{
			name:           "basic case",
			nums:           []int{1, 1, 2},
			expectedLength: 2,
			expectedArray:  []int{1, 2},
		},
		{
			name:           "multiple duplicates",
			nums:           []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedLength: 5,
			expectedArray:  []int{0, 1, 2, 3, 4},
		},
		{
			name:           "no duplicates",
			nums:           []int{1, 2, 3, 4, 5},
			expectedLength: 5,
			expectedArray:  []int{1, 2, 3, 4, 5},
		},
		{
			name:           "single element",
			nums:           []int{1},
			expectedLength: 1,
			expectedArray:  []int{1},
		},
		{
			name:           "all same elements",
			nums:           []int{1, 1, 1, 1},
			expectedLength: 1,
			expectedArray:  []int{1},
		},
		{
			name:           "empty array",
			nums:           []int{},
			expectedLength: 0,
			expectedArray:  []int{},
		},
		{
			name:           "negative numbers",
			nums:           []int{-3, -2, -2, -1, -1, 0, 0, 1},
			expectedLength: 5,
			expectedArray:  []int{-3, -2, -1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies in-place
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			length := removeDuplicates(numsCopy)

			if length != tt.expectedLength {
				t.Errorf("removeDuplicates(%v) length = %v, expected %v", tt.nums, length, tt.expectedLength)
			}

			if length > 0 && !reflect.DeepEqual(numsCopy[:length], tt.expectedArray) {
				t.Errorf("removeDuplicates(%v) array = %v, expected %v", tt.nums, numsCopy[:length], tt.expectedArray)
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 5, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		b.StartTimer()

		removeDuplicates(numsCopy)
	}
}
