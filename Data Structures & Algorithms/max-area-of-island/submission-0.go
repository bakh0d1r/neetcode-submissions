func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0

		return 1 +
			dfs(i-1, j) +
			dfs(i+1, j) +
			dfs(i, j-1) +
			dfs(i, j+1)
	}

	mx := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				mx = max(mx, area)
			}
		}
	}

	return mx
}