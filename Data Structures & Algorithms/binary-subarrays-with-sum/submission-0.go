func numSubarraysWithSum(nums []int, goal int) int {
    count := func(x int) int {
        if x < 0 {
            return 0
        }
        res, l, cur := 0, 0, 0
        for r := 0; r < len(nums); r++ {
            cur += nums[r]
            for cur > x {
                cur -= nums[l]
                l++
            }
            res += r - l + 1
        }
        return res
    }

    return count(goal) - count(goal-1)
}