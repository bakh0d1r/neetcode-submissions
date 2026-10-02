func lengthOfLIS(nums []int) int {
    n := len(nums)
    memo := make([]int, n)
    for i := range memo {
        memo[i] = -1
    }

    var dfs func(int) int
    dfs = func(i int) int {
        if memo[i] != -1 {
            return memo[i]
        }

        c := 1
        for j := i + 1; j < n; j++ {
            if nums[i] < nums[j] {
                c = max(c, 1+dfs(j))
            }
        }

        memo[i] = c
        return c
    }

    mx := 1
    for i := 0; i < n; i++ {
        mx = max(mx, dfs(i))
    }

    return mx
}
