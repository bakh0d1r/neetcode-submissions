func canPartition(nums []int) bool {
    s:=0
	for i:=range nums {
		s+=nums[i]
	}
	if s%2 == 1 {
		return false
	}
	target := s/2

	memo := make([][]int, len(nums)+1)
    for i := range memo {
        memo[i] = make([]int, target+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

	var dfs func(int , int ) bool 
	dfs = func(i int , target int) bool {
		if target == 0 {
			return true
		}
		if i >= len(nums) || target < 0 {
			return false
		}
		if memo[i][target] != -1 {
			return memo[i][target] == 1
		}
		
		found := dfs(i+1,target) || dfs(i+1,target - nums[i])
		if found {
			memo[i][target] = 1
		} else {
			memo[i][target] = 0
		}
		return found
	}
	return dfs(0,target)
}
