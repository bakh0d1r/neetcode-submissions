/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(traversal(root.Right, 0), traversal(root.Left, 0))
}

func traversal(root *TreeNode, md int) int {
	if root == nil {
		return md
	}
	return maxDepth(root)
}
