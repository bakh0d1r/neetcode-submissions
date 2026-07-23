func singleNonDuplicate(nums []int) int {
	l := 0
	h := len(nums)-1
	for l < h {
		m := l+(h-l)/2
		if m%2 == 1 {
			m--
		}
		if nums[m] == nums[m+1] {
			l = m+2
		}else {
			h = m 
		}
	}
	return nums[l]
}
