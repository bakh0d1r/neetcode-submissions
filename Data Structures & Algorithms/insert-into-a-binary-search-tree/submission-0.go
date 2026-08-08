/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return dfs(root, val)
}

func dfs(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val > val {
		root.Left = dfs(root.Left, val)
	} else if root.Val < val {
		root.Right = dfs(root.Right, val)
	} else {
		return root
	}
	return root
}