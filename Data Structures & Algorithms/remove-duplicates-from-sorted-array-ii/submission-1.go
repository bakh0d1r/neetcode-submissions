func removeDuplicates(nums []int) int {
    l,r := 0,0
	for r < len(nums) {
		c:=1
		for r+1 < len(nums) && nums[r] == nums[r+1] {
			r++
			c++
		}
		for i:=0 ; i < min(2,c) ; i++ {
			nums[l]=nums[r]
			l++
		}
		r++
	}
	return l
}