/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	d := &ListNode{}
	c := d
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			c.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			c.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		c = c.Next
	}
	if list1 == nil {
		c.Next = list2
	} else {
		c.Next = list1
	}
	return d.Next
}