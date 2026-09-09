func numIslands(grid [][]byte) int {
	c := 0
	m := len(grid)
	n := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				c++
				dfs(i, j)
			}
		}
	}
	return c
}

