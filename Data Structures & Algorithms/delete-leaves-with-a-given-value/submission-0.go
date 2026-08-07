/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if root == nil {
        return root
    }
    return  dfs(root,target)
}

func dfs(root *TreeNode,target int) *TreeNode{
    if root == nil {
        return root
    }
    root.Left = dfs(root.Left,target)
    root.Right = dfs(root.Right,target)
    if root.Val == target && root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}