func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	memo := make([][2]int, len(nums))
    for i := range memo {
        memo[i][0] = -1
        memo[i][1] = -1
    }

	var dp func(i int,flag int ) int
	dp = func(i int,flag int) int {
		if i >= len(nums) || (flag == 1 && i == len(nums)-1){
			return 0
		}
		if memo[i][flag] != -1 {
			return memo[i][flag]
		}
		t:=flag
		if i == 0 {
			t = 1
		}
		memo[i][flag] = max(nums[i]+dp(i+2,flag), dp(i+1,t))
		return memo[i][flag]
	}
	return max(dp(0,1),dp(1,0))
}

