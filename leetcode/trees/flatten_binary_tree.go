package leetcode

// Flatten Binary Tree to Linked List
// Given the root of a binary tree, flatten the tree into a "linked list":
// - The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
// - The "linked list" should be in the same order as a pre-order traversal of the binary tree.
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	nodes := getAllNodesPreorder(root)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Right = nodes[i+1]
		nodes[i].Left = nil
	}
	if len(nodes) > 0 {
		nodes[len(nodes)-1].Left = nil
	}

}

func getAllNodesPreorder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	var result []*TreeNode

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}
