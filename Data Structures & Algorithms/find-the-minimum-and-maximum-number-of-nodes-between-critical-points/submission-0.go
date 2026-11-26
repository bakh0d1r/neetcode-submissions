/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func nodesBetweenCriticalPoints(head *ListNode) []int {
    previous, current := head, head.Next
    index, first, last, minimum := 1, -1, -1, 100000
    for current.Next != nil {
        value := current.Val
        if (value > previous.Val && value > current.Next.Val) || (value < previous.Val && value < current.Next.Val) {
            if last != -1 {
                minimum = min(minimum, index-last)
            }
            if first == -1 {
                first = index
            }
            last = index
        }
        previous = current
        current = current.Next
        index++
    }
    if first == last {
        return []int{-1, -1}
    }
    return []int{minimum, last - first}
}
