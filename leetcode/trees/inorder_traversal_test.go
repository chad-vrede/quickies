package leetcode

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "example 1 - right skewed with left child",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			expected: []int{1, 3, 2},
		},
		{
			name:     "example 2 - empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "example 3 - single node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name:     "complete tree",
			root:     &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}}},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "left skewed tree",
			root:     &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "right skewed tree",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "unbalanced tree",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}}}},
			expected: []int{2, 1, 4, 5, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := inorderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("inorderTraversal() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkInorderTraversal(t *testing.B) {
	// Create a larger tree for benchmarking
	root := &TreeNode{Val: 50}
	for i := 1; i <= 30; i++ {
		insertIntoBST(root, i)
		insertIntoBST(root, 100-i)
	}

	for i := 0; i < t.N; i++ {
		inorderTraversal(root)
	}
}

// Helper function to create BST for testing
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
