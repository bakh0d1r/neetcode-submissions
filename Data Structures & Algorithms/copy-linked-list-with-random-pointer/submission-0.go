/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	set := map[*Node]*Node{}
	dummy := &Node{}
	new := dummy
	for curr := head; curr != nil; curr = curr.Next {
		set[curr] = &Node{
			Val: curr.Val,
		}
	}
	for curr := head; curr != nil; curr = curr.Next {
		node := set[curr]
		node.Random = set[curr.Random]
		new.Next = node
		new = new.Next
	}
	return dummy.Next
}