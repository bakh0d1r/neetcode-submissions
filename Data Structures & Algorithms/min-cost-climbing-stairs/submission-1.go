func minCostClimbingStairs(cost []int) int {
	memo := make([]int,len(cost))
	for i:=range memo{
		memo[i]=-1
	}
	var dp func(i int) int 
	dp = func(i int) int {
		if i >= len(cost) {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = cost[i]+min(dp(i+1),dp(i+2))
		return memo[i]
	}
	return min(dp(0),dp(1))
}
