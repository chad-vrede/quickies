package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "basic case",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "single negative element",
			nums:     []int{-1},
			expected: -1,
		},
		{
			name:     "all positive",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "all negative",
			nums:     []int{-5, -2, -8, -1},
			expected: -1,
		},
		{
			name:     "mixed with zeros",
			nums:     []int{-2, 0, -1, 3, -1},
			expected: 3,
		},
		{
			name:     "entire array is optimal",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			name:     "alternating positive negative",
			nums:     []int{1, -1, 1, -1, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSubArray(tt.nums)
			if result != tt.expected {
				t.Errorf("maxSubArray(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 3, -2, 6, -1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSubArray(nums)
	}
}
