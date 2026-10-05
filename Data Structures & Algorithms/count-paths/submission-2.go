func uniquePaths(m int, n int) int {
	tab := make([][]int,m+1)
	for i:=range tab {
		tab[i] = make([]int,n+1)
	}
	tab[m-1][n-1] = 1
	for i := m-1 ; i>=0 ;i-- {
		for j:=n -1 ; j>=0;j-- {
			tab[i][j]+=tab[i+1][j]+tab[i][j+1]
		}
	}
	return tab[0][0]
}
