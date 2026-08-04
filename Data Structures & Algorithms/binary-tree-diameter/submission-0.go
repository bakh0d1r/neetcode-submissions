/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	traversal(root, &diameter)
	return diameter
}

func traversal(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := traversal(root.Left, diameter)
	right := traversal(root.Right, diameter)
	*diameter = max(*diameter, left+right)

	return 1 + max(left, right)
}