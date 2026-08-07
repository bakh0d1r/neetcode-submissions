/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}
	if root == nil {
		return arr
	}
	bfst(root, &arr)
	return arr
}

func bfst(root *TreeNode, arr *[][]int) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelVals := []int{}
		n := len(queue)
		for range n  {
			var dequeue *TreeNode
			if queue != nil {
				dequeue = queue[0]
				queue = queue[1:]
			}
			if dequeue != nil {
				levelVals = append(levelVals, dequeue.Val)
				queue = append(queue, dequeue.Left, dequeue.Right)
			}
		}
		if len(levelVals) > 0 {
			(*arr) = append((*arr), levelVals)
		}
	}
}