func createBinaryTree(data [][]int) *TreeNode {
	nodes := map[int]*TreeNode{}
	children := map[int]bool{}
	for i:=range data {
		p:=data[i][0]
		c:=data[i][1]
		isLeft := data[i][2] == 1
		if nodes[p] == nil {
			nodes[p] = &TreeNode{Val:p}
		}
		if nodes[c] == nil {
			nodes[c] = &TreeNode{Val: c}
		}
		if isLeft {
			nodes[p].Left=nodes[c]
		}else{
			nodes[p].Right = nodes[c]
		}
		children[c]=true
	}
	for v, node := range nodes {
		if !children[v] {
			return node
		}
	}
	return nil
}
