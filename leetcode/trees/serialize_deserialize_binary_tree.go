package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// serialize encodes a tree to a single string.
// Serialization is the process of converting a data structure or object into a sequence of bits
// so that it can be stored in a file or memory buffer, or transmitted across a network connection link
// to be reconstructed later in the same or another computer environment.
//
// Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how
// your serialization/deserialization algorithm should work. You just need to ensure that a binary
// tree can be serialized to a string and this string can be deserialized to the original tree structure.
//
// Example:
// You may serialize the following tree:
//
//	  1
//	 / \
//	2   3
//	   / \
//	  4   5
//
// as "1,2,null,null,3,4,null,null,5,null,null"
//
// Time: O(n) where n is the number of nodes
// Space: O(n) for the recursion stack and output string
func serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return fmt.Sprintf("%s,%s,%s", strconv.Itoa(root.Val), serialize(root.Left), serialize(root.Right))
}

// deserialize decodes your encoded data to tree.
// Takes a serialized string and reconstructs the original binary tree structure.
//
// Time: O(n) where n is the number of nodes
// Space: O(n) for the recursion stack and the tokens slice
func deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	index := 0
	return deserializeHelper(tokens, &index)
}

func deserializeHelper(tokens []string, index *int) *TreeNode {
	if *index >= len(tokens) || tokens[*index] == "null" {
		*index++
		return nil
	}

	val, _ := strconv.Atoi(tokens[*index])
	*index++

	node := &TreeNode{Val: val}
	node.Left = deserializeHelper(tokens, index)
	node.Right = deserializeHelper(tokens, index)

	return node
}
