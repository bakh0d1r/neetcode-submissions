/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{Next:head}
	var prev = dummy

	for   range left-1{
		prev=prev.Next
	}

	curr:=prev.Next
	var next *ListNode
	tail:=curr

	for range right - left + 1{
		next = curr.Next
		curr.Next = prev.Next
		prev.Next = curr
		curr=next
	}
	tail.Next=curr
	return dummy.Next
}
