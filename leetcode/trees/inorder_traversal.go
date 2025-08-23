package leetcode

// Binary Tree Inorder Traversal
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
//
// Inorder traversal visits nodes in this order: Left -> Root -> Right
//
// Example 1:
//        1
//         \
//          2
//         /
//        3
// Input: root = [1,null,2,3]
// Output: [1,3,2]
//
// Example 2:
// Input: root = []
// Output: []
//
// Example 3:
// Input: root = [1]
// Output: [1]
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 100].
// - -100 <= Node.val <= 100

func inorderTraversal(root *TreeNode) (traversalPath []int) {
	/** Inorder means you visit nodes in this specific order:
	  1. Left subtree first (all nodes to the left)
	  2. Root (current node)
	  3. Right subtree last (all nodes to the right)
	*/
	return processNode(root)

}

func processNode(root *TreeNode) (traversalPath []int) {
	traversalPath = make([]int, 0)

	if root == nil {
		return
	}

	traversalPath = append(traversalPath, processNode(root.Left)...)
	traversalPath = append(traversalPath, root.Val)
	traversalPath = append(traversalPath, processNode(root.Right)...)
	return
}
