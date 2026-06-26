func twoSum(nums []int, target int) []int {
    for i := range nums {
        for j:=range nums{
            if nums[i]+nums[j] == target && i != j {
                return []int{i,j}
            }
        }
    }
    return []int{}
}