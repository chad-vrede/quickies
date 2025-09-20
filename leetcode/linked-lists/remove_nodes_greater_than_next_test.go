package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveNodes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1: Mixed values with removals",
			input:    []int{5, 2, 13, 3, 8},
			expected: []int{2, 3, 8},
		},
		{
			name:     "Example 2: All same values",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "Example 3: Strictly decreasing",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1},
		},
		{
			name:     "Single node",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Two nodes - first greater",
			input:    []int{5, 3},
			expected: []int{3},
		},
		{
			name:     "Two nodes - first smaller",
			input:    []int{3, 5},
			expected: []int{3, 5},
		},
		{
			name:     "Strictly increasing",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All nodes removed except last",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1},
		},
		{
			name:     "Large values at beginning",
			input:    []int{100, 50, 25, 75, 30, 60},
			expected: []int{25, 30, 60},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedList(tt.input)
			result := removeNodes(head)
			actual := linkedListToSlice(result)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("removeNodes(%v) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}

// Test helper functions
func TestHelperFunctions(t *testing.T) {
	t.Run("createLinkedList and linkedListToSlice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		head := createLinkedList(input)
		output := linkedListToSlice(head)

		if !reflect.DeepEqual(input, output) {
			t.Errorf("Helper functions failed: input %v, output %v", input, output)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		head := createLinkedList([]int{})
		output := linkedListToSlice(head)

		if len(output) != 0 {
			t.Errorf("Empty list should return empty slice, got %v", output)
		}
	})
}

// Benchmark test
func BenchmarkRemoveNodes(b *testing.B) {
	// Create a large test case
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i % 100 // Creates a pattern with many removals
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Need to recreate the list each time since the function modifies it
		testHead := createLinkedList(input)
		removeNodes(testHead)
	}
}
