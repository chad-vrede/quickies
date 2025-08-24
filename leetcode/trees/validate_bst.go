package leetcode

import "math"

// Validate Binary Search Tree
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// A valid BST is defined as follows:
// - The left subtree of a node contains only nodes with keys less than the node's key.
// - The right subtree of a node contains only nodes with keys greater than the node's key.
// - Both the left and right subtrees must also be binary search trees.
func isValidBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return validateBST(root.Left, min, root.Val) && validateBST(root.Right, root.Val, max)
}
