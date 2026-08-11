/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}
	i := 0
	for range inorder {
		if inorder[i] == rootVal {
			break
		}
        i++
	}
    root.Left=buildTree(preorder[1:i+1],inorder[:i+1])
    root.Right=buildTree(preorder[i+1:],inorder[i+1:])
    return root
}