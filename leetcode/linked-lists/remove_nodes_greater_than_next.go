package leetcode

// Remove Nodes From Linked List
// You are given the head of a linked list. Remove all nodes from the linked list
// that have a value greater than the next node's value. Return the head of the modified linked list.
//
// Example 1:
// Input: head = [5,2,13,3,8]
// Output: [2,3,8]
// Explanation:
// - 5 > 2, so remove 5
// - 13 > 3, so remove 13
// - 2 < 13 (before removal), keep 2
// - 3 < 8, keep 3
// - 8 has no next, keep 8
//
// Example 2:
// Input: head = [1,1,1,1]
// Output: [1,1,1,1]
// Explanation: No node is greater than its next, so keep all.
//
// Example 3:
// Input: head = [5,4,3,2,1]
// Output: [1]
// Explanation: Each node (except the last) is greater than the next, so only keep the last.
//
// Constraints:
// - The number of nodes in the list is in the range [1, 10^5]
// - 1 ≤ Node.val ≤ 10^5
func removeNodes(head *ListNode) *ListNode {
	var stack []*ListNode

	for head != nil {
		// Remove nodes from stack that are greater than current
		for len(stack) > 0 && stack[len(stack)-1].Val > head.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head)
		head = head.Next
	}

	if len(stack) == 0 {
		return nil
	}

	for i := 0; i < len(stack)-1; i++ {
		stack[i].Next = stack[i+1]
	}

	stack[len(stack)-1].Next = nil

	return stack[0]
}
