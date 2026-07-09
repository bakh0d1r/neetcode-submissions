func minSubArrayLen(target int, nums []int) int {
	s:=0
	l:=0
	ml:=math.MaxInt64
	for r:=range nums{
		s+=nums[r]
		for s >= target {
			ml=min(ml,r-l+1)
			s-=nums[l]
			l++
		}
	}
	if ml == math.MaxInt64 {
		return 0
	}
	return ml
}
