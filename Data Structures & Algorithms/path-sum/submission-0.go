func hasPathSum(root *TreeNode, targetSum int) bool {
    sum:=0
    return dfs(root,sum,targetSum)
}

func dfs(root *TreeNode,sum int,targetSum int) bool{
    if root == nil {
        return false
    }
    sum+=root.Val
    if sum == targetSum && root.Left == nil && root.Right == nil{
        return true
    }
    return dfs(root.Left,sum,targetSum) || dfs(root.Right,sum,targetSum)
}
