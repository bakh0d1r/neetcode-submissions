func search(nums []int, target int) bool {
	l:=0
	h:=len(nums)-1
	for l<=h{
		m:=l+(h-l)/2
		if nums[m] == target {
			return true
		}
		if nums[m] > nums[l] {
			if nums[l] <= target && target <= nums[m] {
				h = m -1
			}else {
				l = m+1
			}
		} else if nums[m] < nums[l] {
			if nums[m] <= target && target <= nums[h] {
				l = m + 1
			}else {
				h = m - 1
			}
		} else {
			l++
		}
	}
	return false
}
