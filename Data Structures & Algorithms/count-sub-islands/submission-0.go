func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    rows, cols := len(grid2), len(grid2[0])
    answer := 0

    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        if r < 0 || r >= rows || c < 0 || c >= cols || grid2[r][c] == 0 {
            return true
        }

        grid2[r][c] = 0

        isSub := grid1[r][c] == 1

        up := dfs(r-1, c)
        down := dfs(r+1, c)
        left := dfs(r, c-1)
        right := dfs(r, c+1)

        return isSub && up && down && left && right
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid2[r][c] == 1 {
                if dfs(r, c) {
                    answer++
                }
            }
        }
    }

    return answer
}