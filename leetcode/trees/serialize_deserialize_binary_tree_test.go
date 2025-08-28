package leetcode

import (
	"testing"
)

// Helper function to create a binary tree for testing
func createBinaryTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	var queue TreeNodeQueue
	queue.Enqueue(root)

	i := 1
	for i < len(vals) && queue.Len() > 0 {
		node := queue.Dequeue()

		// Left child
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue.Enqueue(node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue.Enqueue(node.Right)
		}
		i++
	}

	return root
}

// Helper function to compare two binary trees
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestSerializeDeserialize(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		expected string
	}{
		{
			name:     "example tree",
			tree:     []interface{}{1, 2, 3, nil, nil, 4, 5},
			expected: "1,2,null,null,3,4,null,null,5,null,null",
		},
		{
			name:     "empty tree",
			tree:     []interface{}{},
			expected: "null",
		},
		{
			name:     "single node",
			tree:     []interface{}{1},
			expected: "1,null,null",
		},
		{
			name:     "left skewed tree",
			tree:     []interface{}{1, 2, nil, 3, nil},
			expected: "1,2,3,null,null,null,null",
		},
		{
			name:     "right skewed tree",
			tree:     []interface{}{1, nil, 2, nil, 3},
			expected: "1,null,2,null,3,null,null",
		},
		{
			name:     "perfect binary tree",
			tree:     []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: "1,2,4,null,null,5,null,null,3,6,null,null,7,null,null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the original tree
			originalTree := createBinaryTree(tt.tree)

			// Test serialization
			serialized := serialize(originalTree)
			if serialized != tt.expected {
				t.Errorf("serialize() = %v, want %v", serialized, tt.expected)
			}

			// Test deserialization
			deserialized := deserialize(serialized)

			// Verify the deserialized tree matches the original
			if !isSameTree(originalTree, deserialized) {
				t.Errorf("deserialized tree does not match original tree")
			}

			// Test round trip: serialize the deserialized tree again
			reSerialized := serialize(deserialized)
			if reSerialized != tt.expected {
				t.Errorf("re-serialize() = %v, want %v", reSerialized, tt.expected)
			}
		})
	}
}

func TestDeserializeInvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "empty string",
			input: "",
		},
		{
			name:  "only commas",
			input: ",,,",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Should not panic and should handle gracefully
			result := deserialize(tt.input)
			_ = result // Just ensure no panic occurs
		})
	}
}

func BenchmarkSerialize(b *testing.B) {
	// Create a moderately sized tree for benchmarking
	tree := createBinaryTree([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serialize(tree)
	}
}

func BenchmarkDeserialize(b *testing.B) {
	// Pre-serialize a tree for benchmarking deserialization
	tree := createBinaryTree([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	serialized := serialize(tree)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deserialize(serialized)
	}
}
