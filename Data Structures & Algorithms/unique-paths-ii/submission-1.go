func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m :=len(obstacleGrid)
    n :=len(obstacleGrid[0])
	memo := make([][]int,m)
	for i:=range memo {
		memo[i] = make([]int, n)
		for j :=range memo[i] {
			memo[i][j]=-1
		}
	}
    var dfs func(i,j int) int
	dfs = func(i , j int) int{
		if i == m || j == n || obstacleGrid[i][j] == 1 {
			return 0
		}
		if i == m-1 && j == n-1 {
			return 1
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		memo[i][j] = dfs(i+1,j) + dfs(i,j+1)
		return memo[i][j]
	}
	return dfs(0,0)
}