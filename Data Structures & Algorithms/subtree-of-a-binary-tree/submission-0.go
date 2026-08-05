
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isEqualTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqualTree(root *TreeNode, subRoot *TreeNode) bool {
	if (root == nil) || (subRoot == nil) {
		return root == subRoot
	}
	return (root.Val == subRoot.Val) && isEqualTree(root.Left, subRoot.Left) && isEqualTree(root.Right, subRoot.Right)
}
