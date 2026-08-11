/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	return dfsMove(root, arr, 0)
}

func dfsMove(root *TreeNode, arr []int, depth int) []int {
	if root == nil {
		return arr
	}
	if depth == len(arr) {
		arr = append(arr, root.Val)
	}
	arr = dfsMove(root.Right, arr, depth+1)
	arr = dfsMove(root.Left, arr, depth+1)
	return arr
}