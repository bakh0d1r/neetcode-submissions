/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	gm := map[*Node]*Node{}
	return clone(node, gm)
}

func clone(curr *Node, gm map[*Node]*Node) *Node {
	if node, ok := gm[curr]; ok {
		return node
	}
	new := &Node{}
	new.Val = curr.Val
	gm[curr] = new
	for n := range curr.Neighbors {
		new.Neighbors = append(new.Neighbors, clone(curr.Neighbors[n], gm))
	}
	return new
}