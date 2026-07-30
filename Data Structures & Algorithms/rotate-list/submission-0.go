/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head 
    }
	n := 1
	tail := head
	for ; tail.Next != nil; tail = tail.Next {
		n++
	}
	k = k % n
    if k == 0 {
        return head
    }
	curr := head
	c := 0
	for ; curr != nil; curr = curr.Next {
		c++
		if n-k == c {
			break
		}
	}
	node := curr.Next
	curr.Next = nil
	tail.Next = head
	return node
}
