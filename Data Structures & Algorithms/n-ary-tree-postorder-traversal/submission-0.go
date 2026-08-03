/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	arr := []int{}
	POT(root, &arr)
	return arr
}

func POT(node *Node, arr *[]int) {
	if node == nil {
		return
	}
	for i:=range node.Children {
		POT(node.Children[i], arr)
	}
	*arr = append(*arr, node.Val)
}