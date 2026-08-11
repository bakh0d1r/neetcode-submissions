/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return dfsMove(root,math.MinInt64,math.MaxInt64)
}

func dfsMove(root *TreeNode,mn,mx int) bool {
	if root == nil {
		return true
	}

	if root.Val <= mn || root.Val >= mx {
    	return false
	}

	return dfsMove(root.Left, mn,root.Val) &&
       dfsMove(root.Right, root.Val,mx)
}
