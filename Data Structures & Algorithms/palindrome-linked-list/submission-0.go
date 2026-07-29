/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	s:=[]int{}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		s=append(s,slow.Val)
		slow = slow.Next
	}
	if fast != nil {
		slow=slow.Next
	}
	for slow != nil {
		if len(s) > 0 && slow.Val != s[len(s)-1] {
			return false
		}
		s=s[:len(s)-1]
		slow=slow.Next
	}
	return true
}
