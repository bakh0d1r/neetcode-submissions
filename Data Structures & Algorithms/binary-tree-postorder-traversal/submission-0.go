/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	arr := []int{}
	POT(root, &arr)
	return arr
}

func POT(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	POT(node.Left, arr)
	POT(node.Right, arr)
	*arr = append(*arr, node.Val)
}