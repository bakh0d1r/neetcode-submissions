/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    arr := []int{}
    cur := head

    for cur != nil {
        arr = append(arr, cur.Val)
        cur = cur.Next
    }

    sort.Ints(arr)
    cur = head
    for _, val := range arr {
        cur.Val = val
        cur = cur.Next
    }

    return head
}