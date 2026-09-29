func rob(nums []int) int {
	memo := make([]int,len(nums)+1)
	for i:=range memo{
		memo[i]=-1
	}
	var dp func(i int) int 
	dp = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = max( nums[i]+dp(i+2), dp(i+1) )
		return memo[i]
	}
	return dp(0)
}

