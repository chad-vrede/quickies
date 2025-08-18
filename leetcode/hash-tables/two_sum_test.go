package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "basic case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "target with negative numbers",
			nums:     []int{3, 2, -4},
			target:   -1,
			expected: []int{0, 2},
		},
		{
			name:     "duplicate numbers",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "larger array",
			nums:     []int{1, 5, 3, 7, 9, 2},
			target:   10,
			expected: []int{2, 3},
		},
		{
			name:     "zero target",
			nums:     []int{-1, 0, 1, 2},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "negative target",
			nums:     []int{-3, -2, -1},
			target:   -5,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSum(%v, %d) = %v, expected %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func TestTwoSumNoSolution(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 10
	result := twoSum(nums, target)
	expected := []int{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("twoSum(%v, %d) = %v, expected %v (no solution case)", nums, target, result, expected)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 6, 8, 1, 4, 9}
	target := 9

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
