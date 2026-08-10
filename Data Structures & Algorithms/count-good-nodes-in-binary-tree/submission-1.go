/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	mx := root.Val
	return bfsMove(root, mx)
}

func bfsMove(root *TreeNode, mx int) int {
	if root == nil {
		return 0
	}

	if mx <= root.Val {
        mx=root.Val
        return bfsMove(root.Left,mx) + bfsMove(root.Right,mx)  + 1
	}
	
	return bfsMove(root.Left,mx) + bfsMove(root.Right,mx)
}