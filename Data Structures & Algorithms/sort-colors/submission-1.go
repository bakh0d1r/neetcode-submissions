func sortColors(nums []int) {
    l,m,r:=0,0,len(nums)-1
	for m <= r {
		if nums[m]==0 {
			nums[l],nums[m]=nums[m],nums[l]
			l++
			m++
		} else if nums[m] == 2 {
			nums[r],nums[m]=nums[m],nums[r]
			r--
		} else if nums[m] == 1{
			m++
		}
	}
}
