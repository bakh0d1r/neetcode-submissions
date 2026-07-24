func searchRange(nums []int, target int) []int {
	l := 0
	h := len(nums) - 1
	j, k := -1, -1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] > target {
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			j = m
			h = m - 1
		}
	}
	if j == -1 {
		return []int{j, k}
	}
	l = 0
	h = len(nums) - 1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] > target {
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			k = m
			l = m + 1
		}
	}
	return []int{j, k}
}