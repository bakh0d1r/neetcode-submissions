/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow,fast:=head,head

	for fast.Next != nil && fast.Next.Next != nil {
		fast=fast.Next.Next
		slow=slow.Next
	}
	second:=slow.Next
	slow.Next=nil
	
	var prev *ListNode

	for second != nil {
		next := second.Next
		second.Next = prev
		prev = second
		second = next	
	}

	first , second := head , prev
	for second != nil {
		firstNext:= first.Next
		secondNext := second.Next

		first.Next = second
		second.Next = firstNext

		second = secondNext
		first = firstNext
	}
}
