/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    var helper func(*TreeNode) (int, int)
    helper = func(root *TreeNode) (int, int){
        if root == nil {
            return 0,0
        }

        withRootLeft, withoutRootLeft := helper(root.Left)
        withRootRight, withoutRootRight := helper(root.Right)

        withRoot := root.Val + withoutRootLeft + withoutRootRight
        withoutRoot := max(withRootLeft, withoutRootLeft)+max(withRootRight, withoutRootRight)

        return withRoot, withoutRoot
    }

    withRoot, withoutRoot := helper(root)
    return max(withRoot, withoutRoot)
}