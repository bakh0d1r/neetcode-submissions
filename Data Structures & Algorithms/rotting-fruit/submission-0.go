func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	queue := make([][2]int, 0)

	fresh := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]

			directions := [][2]int{
				{-1, 0}, // up
				{1, 0},  // down
				{0, -1}, // left
				{0, 1},  // right
			}

			for _, d := range directions {
				nr := r + d[0]
				nc := c + d[1]

				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}

				if grid[nr][nc] != 1 {
					continue
				}

				grid[nr][nc] = 2
				fresh--

				queue = append(queue, [2]int{nr, nc})
			}
		}
		minutes++
	}

	if fresh > 0 {
		return -1
	}

	return minutes
}