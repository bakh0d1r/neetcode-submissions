func longestCommonSubsequence(text1 string, text2 string) int {
    m,n:=len(text1),len(text2)
	memo:=make([][]int,m+1)
	for i:=range memo {
		memo[i] = make([]int,n+1)
		for j:=range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i,j int) int 
	dfs = func(i,j int) int {
		if i == m || j == n {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if text1[i] == text2[j] {
			memo[i][j] = 1 + dfs(i+1,j+1)
		} else {
			memo[i][j] = max(dfs(i+1,j),dfs(i,j+1))
		}
		return memo[i][j]
	}
	return dfs(0,0)
}
