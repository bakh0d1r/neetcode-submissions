/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    return move(root,&k)
}

func move(root *TreeNode,k *int) int{
	if root == nil {
		return 0
	}
	m:=move(root.Left,k)
	if *k == 0 {
		return m
	}
	(*k)--
	if *k == 0 {
		return root.Val
	}
	return move(root.Right,k)
}
