func longestCommonSubsequence(text1 string, text2 string) int {
    m,n:=len(text1),len(text2)
	tab := make([][]int, m+1)
    for i := range tab {
        tab[i] = make([]int, n+1)
    }
	for i:= m -1 ; i >=0;i-- {
		for j:=n-1;j>=0;j--{
			if text1[i] == text2[j]{ 
				tab[i][j] = 1 + tab[i+1][j+1]
			}else{
				tab[i][j]=max(tab[i+1][j],tab[i][j+1])
			}
		}
	}
	return tab[0][0]
}
