package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "contains duplicates",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "all same elements",
			nums:     []int{1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "two identical elements",
			nums:     []int{1, 1},
			expected: true,
		},
		{
			name:     "negative numbers with duplicates",
			nums:     []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "large array with duplicate at end",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("containsDuplicate(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate(nums)
	}
}
