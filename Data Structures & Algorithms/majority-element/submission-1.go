func majorityElement(nums []int) int {
    candidate, vote := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			vote++
		} else if vote == 0 {
			candidate = nums[i]
			vote = 1
		} else {
			vote--
		}
	}
	return candidate
}