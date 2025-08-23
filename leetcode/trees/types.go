package leetcode

// TreeNode definition for binary tree problems
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TreeNodeQueue implementation
type TreeNodeQueue []*TreeNode

// Generic Queue interface
type Queue[T any] interface {
	Dequeue() T
	Enqueue(v T)
	Len() int
}

func (q *TreeNodeQueue) Dequeue() *TreeNode {
	if q == nil || len(*q) == 0 {
		return nil
	}
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

func (q *TreeNodeQueue) Enqueue(t *TreeNode) {
	*q = append(*q, t)
}

func (q *TreeNodeQueue) Len() int {
	return len(*q)
}
