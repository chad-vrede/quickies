package leetcode

// Binary Tree Level Order Traversal
// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).
//
// Example 1:
//        3
//       / \
//      9  20
//        /  \
//       15   7
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]
//
// Example 2:
// Input: root = [1]
// Output: [[1]]
//
// Example 3:
// Input: root = []
// Output: []
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 2000].
// - -1000 <= Node.val <= 1000

func levelOrder(root *TreeNode) (levelOrder [][]int) {
	levelOrder = make([][]int, 0)
	visitationQueue := make(TreeNodeQueue, 0)

	if root == nil {
		return
	}
	visitationQueue = append(visitationQueue, root)

	for visitationQueue.Len() > 0 {
		level := []int{}
		queueLength := visitationQueue.Len()
		// size of queue is the amount of nodes currently on this level
		for i := 0; i < queueLength; i++ {
			node := visitationQueue.Dequeue()
			if node != nil {
				level = append(level, node.Val)
				if node.Left != nil {
					visitationQueue.Enqueue(node.Left)
				}
				if node.Right != nil {
					visitationQueue.Enqueue(node.Right)
				}
			}
		}
		levelOrder = append(levelOrder, level)
	}
	return
}
