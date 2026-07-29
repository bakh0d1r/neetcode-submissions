/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	e := 0
	for l1 != nil || l2 != nil  || e > 0 {
		n := 0
		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}
		n += e
		e = n / 10
		n = n % 10
		curr.Next = &ListNode{
			Val: n,
		}
		curr = curr.Next
	}
   
	return dummy.Next
}