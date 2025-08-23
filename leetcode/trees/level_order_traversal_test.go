package leetcode

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
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
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "example 2 - single node",
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name:     "example 3 - empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name: "right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: [][]int{{1}, {2}, {3}},
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
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name: "unbalanced tree",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:   4,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			expected: [][]int{{1}, {2, 3}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrder(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkLevelOrder(t *testing.B) {
	// Create a larger balanced tree for benchmarking
	root := &TreeNode{Val: 1}
	queue := []*TreeNode{root}
	val := 2

	// Build a tree with multiple levels
	for level := 0; level < 10 && len(queue) > 0; level++ {
		levelSize := len(queue)
		for i := 0; i < levelSize && val <= 1000; i++ {
			node := queue[0]
			queue = queue[1:]

			if val <= 1000 {
				node.Left = &TreeNode{Val: val}
				queue = append(queue, node.Left)
				val++
			}
			if val <= 1000 {
				node.Right = &TreeNode{Val: val}
				queue = append(queue, node.Right)
				val++
			}
		}
	}

	for i := 0; i < t.N; i++ {
		levelOrder(root)
	}
}
