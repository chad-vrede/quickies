package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "all same frequency",
			nums:     []int{1, 2, 3},
			k:        2,
			expected: []int{1, 2}, // any 2 elements are valid
		},
		{
			name:     "k equals array length",
			nums:     []int{4, 1, -1, 2, -1, 2, 3},
			k:        4,
			expected: []int{-1, 2, 1, 3}, // order may vary for same frequency
		},
		{
			name:     "single element repeated",
			nums:     []int{5, 5, 5, 5},
			k:        1,
			expected: []int{5},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -2, -2, -2, -3},
			k:        2,
			expected: []int{-2, -1},
		},
		{
			name:     "large numbers",
			nums:     []int{100, 200, 100, 200, 200, 300},
			k:        2,
			expected: []int{200, 100},
		},
		{
			name:     "mixed positive negative",
			nums:     []int{0, 0, 0, 1, 1, -1},
			k:        2,
			expected: []int{0, 1},
		},
		{
			name:     "zero frequency edge case",
			nums:     []int{0, 1, 2, 0, 1, 0},
			k:        2,
			expected: []int{0, 1},
		},
		{
			name:     "k is 1",
			nums:     []int{3, 0, 1, 0},
			k:        1,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)

			// Since the order of elements with same frequency can vary,
			// we need to check if the result contains the expected elements
			if len(result) != len(tt.expected) {
				t.Errorf("topKFrequent(%v, %d) returned %d elements, expected %d",
					tt.nums, tt.k, len(result), len(tt.expected))
				return
			}

			// Sort both slices to compare regardless of order
			sortedResult := make([]int, len(result))
			copy(sortedResult, result)
			sort.Ints(sortedResult)

			sortedExpected := make([]int, len(tt.expected))
			copy(sortedExpected, tt.expected)
			sort.Ints(sortedExpected)

			// For cases where frequency matters, we need to verify the actual frequencies
			freqMap := make(map[int]int)
			for _, num := range tt.nums {
				freqMap[num]++
			}

			// Check if all returned elements are among the most frequent
			resultFreqs := make([]int, len(result))
			for i, num := range result {
				resultFreqs[i] = freqMap[num]
			}
			sort.Sort(sort.Reverse(sort.IntSlice(resultFreqs)))

			// Verify we got k elements and they're among the top frequencies
			allFreqs := make([]int, 0, len(freqMap))
			for _, freq := range freqMap {
				allFreqs = append(allFreqs, freq)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(allFreqs)))

			for i, freq := range resultFreqs {
				if i < len(allFreqs) && freq < allFreqs[tt.k-1] {
					t.Errorf("topKFrequent(%v, %d) = %v, but element with frequency %d is not among top %d",
						tt.nums, tt.k, result, freq, tt.k)
				}
			}
		})
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i % 100 // Create frequency distribution
	}
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topKFrequent(nums, k)
	}
}

// Helper function to check if two slices contain the same elements (ignoring order)
func containsSameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	countA := make(map[int]int)
	countB := make(map[int]int)

	for _, v := range a {
		countA[v]++
	}
	for _, v := range b {
		countB[v]++
	}

	return reflect.DeepEqual(countA, countB)
}
