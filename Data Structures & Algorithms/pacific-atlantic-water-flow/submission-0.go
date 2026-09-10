func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	arr := [][]int{}

	sp := make([][]bool, m)
	sa := make([][]bool, m)

	for i := 0; i < m; i++ {
		sp[i] = make([]bool, n)
		sa[i] = make([]bool, n)
	}

	for i := range m {
		for j := range n {
			if i == 0 || j == 0 {
				dfs(i, j, heights[i][j], heights, sp)
			}
			if i == m-1 || j == n-1 {
				dfs(i, j, heights[i][j], heights, sa)
			}
		}
	}
	for i := range m {
		for j := range n {
			if sp[i][j] == sa[i][j] && sp[i][j] == true {
				arr = append(arr, []int{i, j})
			}
		}
	}
	return arr
}

func dfs(i, j, height int, heights [][]int, set [][]bool) {
	m := len(heights)
	n := len(heights[0])
	if i < 0 || j < 0 || i >= m || j >= n || set[i][j] || heights[i][j] < height {
		return
	}

	set[i][j] = true
	dfs(i-1, j, heights[i][j], heights, set) // up
	dfs(i+1, j, heights[i][j], heights, set) // down
	dfs(i, j-1, heights[i][j], heights, set) // right
	dfs(i, j+1, heights[i][j], heights, set) // left
}
