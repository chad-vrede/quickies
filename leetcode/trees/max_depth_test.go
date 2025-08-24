package leetcode

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "example 1 - balanced tree",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
		{
			name: "example 2 - right skewed",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			expected: 2,
		},
		{
			name:     "example 3 - empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "example 4 - single node",
			root:     &TreeNode{Val: 0},
			expected: 1,
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			expected: 4,
		},
		{
			name: "complete binary tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
		{
			name: "unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  4,
							Left: &TreeNode{Val: 5},
						},
					},
				},
				Right: &TreeNode{Val: 6},
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDepth(tt.root)
			if result != tt.expected {
				t.Errorf("maxDepth() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func BenchmarkMaxDepth(t *testing.B) {
	// Create a deep tree for benchmarking
	root := &TreeNode{Val: 1}
	current := root
	for i := 2; i <= 1000; i++ {
		current.Left = &TreeNode{Val: i}
		current = current.Left
	}

	for i := 0; i < t.N; i++ {
		maxDepth(root)
	}
}
