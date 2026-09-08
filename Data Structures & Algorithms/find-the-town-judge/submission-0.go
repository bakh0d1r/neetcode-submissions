func findJudge(n int, trust [][]int) int {
	in := make([]int, n+1)
	out := make([]int, n+1)
	for i := range trust {
		a := trust[i][0]
		b := trust[i][1]
		in[b]++
		out[a]++
	}
	for i := 1; i < n+1; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i
		}
	}
	return -1
}