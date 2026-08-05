/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, balanced := travers(root)
	return balanced
}

func travers(root *TreeNode) (int, bool) {
	if root == nil {
		return -1, true
	}
	left, lb := travers(root.Left)
	right, rb := travers(root.Right)
	h := 1 + max(left, right)

	b := lb && rb && (max(right, left)-min(right, left) <= 1)
	return h, b
}