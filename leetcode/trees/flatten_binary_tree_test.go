package leetcode

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int // expected values in flattened order (preorder)
	}{
		{
			name: "example tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "empty tree",
			root:     nil,
			expected: nil,
		},
		{
			name: "single node",
			root: &TreeNode{
				Val: 0,
			},
			expected: []int{0},
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the tree for testing
			originalRoot := copyTree(tt.root)

			flatten(originalRoot)

			// Convert flattened tree to slice for comparison
			result := flattenedTreeToSlice(originalRoot)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("flatten() = %v, expected %v", result, tt.expected)
			}

			// Verify structure: all left children should be nil
			verifyFlattenedStructure(t, originalRoot)
		})
	}
}

// Helper function to copy a tree for testing
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

// Helper function to convert flattened tree to slice
func flattenedTreeToSlice(root *TreeNode) []int {
	var result []int
	current := root
	for current != nil {
		result = append(result, current.Val)
		current = current.Right
	}
	return result
}

// Helper function to verify the flattened tree structure
func verifyFlattenedStructure(t *testing.T, root *TreeNode) {
	current := root
	for current != nil {
		if current.Left != nil {
			t.Errorf("Found non-nil left child at node with value %d", current.Val)
		}
		current = current.Right
	}
}

func BenchmarkFlatten(b *testing.B) {
	// Create a balanced tree for benchmarking
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testRoot := copyTree(root)
		flatten(testRoot)
	}
}
