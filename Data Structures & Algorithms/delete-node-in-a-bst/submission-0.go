/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	return dfsMove(root, key)
}

func dfsMove(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		val := dfsMoveMin(root.Right)
		root.Val = val
		root.Right = deleteNode(root.Right, val)
		return root
	}
	if root.Val > key {
		root.Left = dfsMove(root.Left, key)
	}
	if root.Val < key {
		root.Right = dfsMove(root.Right, key)
	}
	return root
}

func dfsMoveMin(root *TreeNode) int {
	if root.Left == nil {
		return root.Val
	}
	return dfsMoveMin(root.Left)
}