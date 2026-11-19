func sortArrayByParity(nums []int) []int {
    j := 0
    for i, v := range nums {
        if v % 2 == 0 {
            nums[j], nums[i] = nums[i], nums[j]
            j++
        }
    }
    return nums
}