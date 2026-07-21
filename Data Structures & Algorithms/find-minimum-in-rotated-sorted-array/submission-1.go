func findMin(nums []int) int {
	l := 0
	h := len(nums) - 1
	for l < h {
		m := l + (h-l)/2
		if nums[m] > nums[h] {
			l = m + 1
		} else {
			h = m 
		}
	}
	return nums[l]
}
