package leetcode

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "valid BST example 1",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "invalid BST example 1",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			want: false,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: true,
		},
		{
			name: "empty tree",
			root: nil,
			want: true,
		},
		{
			name: "left subtree violation",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 15},
				},
			},
			want: false,
		},
		{
			name: "right subtree violation",
			root: &TreeNode{
				Val: 10,
				Right: &TreeNode{
					Val:  15,
					Left: &TreeNode{Val: 6},
				},
			},
			want: false,
		},
		{
			name: "duplicate values",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 1},
			},
			want: false,
		},
		{
			name: "edge case with min/max int",
			root: &TreeNode{
				Val:   -2147483648,
				Right: &TreeNode{Val: 2147483647},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsValidBST(b *testing.B) {
	// Create a balanced BST for benchmarking
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 18},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isValidBST(root)
	}
}
