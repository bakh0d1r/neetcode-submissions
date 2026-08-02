/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := head
	curr := head.Next
	for curr != nil {
		if curr.Val > prev.Val {
			prev = curr
			curr = curr.Next
			continue
		}
		node := dummy
		for node.Next.Val < curr.Val {
			node = node.Next
		}
		prev.Next = curr.Next
		curr.Next = node.Next
		node.Next = curr
		curr = prev.Next
	}
	return dummy.Next
}