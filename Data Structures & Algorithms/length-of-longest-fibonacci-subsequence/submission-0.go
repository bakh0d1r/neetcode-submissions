func lenLongestFibSubseq(arr []int) int {
	n, mx := len(arr), 0
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 2
		}
	}

	for k := 2; k < n; k++ {
		left, right := 0, k-1
		for left < right {
			total := arr[left] + arr[right]
			if total < arr[k] {
				left++
			} else if total > arr[k] {
				right--
			} else {
				dp[right][k] = dp[left][right] + 1
				mx = max(mx, dp[right][k])
				left++
				right--
			}
		}
	}
	return mx
}
