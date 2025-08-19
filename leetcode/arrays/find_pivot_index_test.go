package leetcode

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "basic pivot found",
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			name:     "no pivot exists",
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "pivot at left edge",
			nums:     []int{2, 1, -1},
			expected: 0,
		},
		{
			name:     "single element",
			nums:     []int{5},
			expected: 0,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "pivot at right edge",
			nums:     []int{-1, 1, 0},
			expected: 2,
		},
		{
			name:     "two elements no pivot",
			nums:     []int{1, 2},
			expected: -1,
		},
		{
			name:     "large positive and negative",
			nums:     []int{1000, -1000, 0},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pivotIndex(tt.nums)
			if result != tt.expected {
				t.Errorf("pivotIndex(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkPivotIndex(b *testing.B) {
	nums := []int{1, 7, 3, 6, 5, 6, 2, 8, 4, 9, 1, 3, 5, 7, 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pivotIndex(nums)
	}
}
