package leetcode

import (
	"reflect"
	"testing"
)

// Helper function to create linked list from slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

// Helper function to convert linked list to slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "basic addition",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "single digits",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "different lengths",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name:     "carry propagation",
			l1:       []int{5},
			l2:       []int{5},
			expected: []int{0, 1},
		},
		{
			name:     "one empty list",
			l1:       []int{1, 2, 3},
			l2:       []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "multiple carries",
			l1:       []int{9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 1},
		},
		{
			name:     "no carry",
			l1:       []int{1, 2},
			l2:       []int{3, 4},
			expected: []int{4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createLinkedList(tt.l1)
			l2 := createLinkedList(tt.l2)

			result := addTwoNumbers(l1, l2)
			resultSlice := linkedListToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, expected %v", tt.l1, tt.l2, resultSlice, tt.expected)
			}
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create fresh copies for each iteration since the function consumes the lists
		l1Copy := createLinkedList([]int{2, 4, 3, 6, 8})
		l2Copy := createLinkedList([]int{5, 6, 4, 2, 7})
		addTwoNumbers(l1Copy, l2Copy)
	}
}
