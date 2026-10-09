func minPathSum(grid [][]int) int {
    m:=len(grid)
    n:=len(grid[0])
	tab := make([][]int,m+1)
	for i:=range tab {
		tab[i] = make([]int,n+1)
		for j :=range tab[i] {
			tab[i][j] = 1 << 30
		}
	}
	tab[m-1][n]=0
	for i := m-1 ; i>=0 ;i-- {
		for j:=n -1 ; j>=0;j-- {
            tab[i][j] = grid[i][j] + min(tab[i+1][j], tab[i][j+1])
		}
	}
	return tab[0][0]
}