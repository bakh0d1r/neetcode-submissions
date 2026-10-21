func canJump(nums []int) bool {
    mx := 0
    n := len(nums)-1
    for i := 0; i <= mx; i++ {
        mx = max(mx, i+nums[i])
        if mx >= n {
            return true
        }
    }
    return false
}
