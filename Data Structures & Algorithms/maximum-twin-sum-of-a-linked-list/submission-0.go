/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	mx := 0
	curr := head
	currHalf := head
	var prev *ListNode
	for currHalf != nil && currHalf.Next != nil {
		currHalf = currHalf.Next.Next
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	for curr != nil {
		mx = max(mx, curr.Val+prev.Val)
		curr = curr.Next
		prev = prev.Next
	}

	return mx
}