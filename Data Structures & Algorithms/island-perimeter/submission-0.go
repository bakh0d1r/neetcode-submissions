func islandPerimeter(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	c := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				//up
				if i-1 == -1 || grid[i-1][j] == 0 {
					c++
				}

				//down
				if i+1 == m || grid[i+1][j] == 0 {
					c++
				}

				//left
				if j-1 == -1 || grid[i][j-1] == 0 {
					c++
				}

				//right
				if j+1 == n || grid[i][j+1] == 0 {
					c++
				}
			}
		}
	}
	return c
}