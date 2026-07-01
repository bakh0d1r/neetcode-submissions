func removeElement(nums []int, val int) int {
    l,r:= 0,0
	for r<len(nums) {
		if nums[r] == val{
			r++
		}else{
			nums[l] = nums[r]
			l++
			r++
		}
	}
	return l
}
