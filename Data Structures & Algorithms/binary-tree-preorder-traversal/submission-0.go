func preorderTraversal(root *TreeNode) []int {
	arr := []int{}
	POT(root, &arr)
	return arr
}

func POT(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	*arr = append(*arr, node.Val)
	POT(node.Left, arr)
	POT(node.Right, arr)
}