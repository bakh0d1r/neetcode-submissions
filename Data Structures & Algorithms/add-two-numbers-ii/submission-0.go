/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := []int{}
	s2 := []int{}
	for l1 != nil || l2 != nil {
		if l1 != nil {
			s1 = append(s1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			s2 = append(s2, l2.Val)
			l2 = l2.Next
		}
	}
	e := 0
	var l3 *ListNode
	for len(s1) > 0 || len(s2) > 0 || e != 0 {
		n := 0
		if len(s1) > 0 {
			n += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			n += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		n += e
		e = n / 10
		n = n % 10
		l3 = &ListNode{
			Val:  n,
			Next: l3,
		}
	}
	return l3
}