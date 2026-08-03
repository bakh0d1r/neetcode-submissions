func inorderTraversal(root *TreeNode) []int {
	arr := []int{}
	IOT(root, &arr)
	return arr
}

func IOT(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	IOT(node.Left, arr)
	*arr = append(*arr, node.Val)
	IOT(node.Right, arr)
}