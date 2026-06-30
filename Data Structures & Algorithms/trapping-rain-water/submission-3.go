func trap(height []int) int {
	lm,rm:=0,0
	t:=0
	l,r:= 0,len(height)-1
	for l<r{
		if height[l] < height[r] {
			lm=max(lm,height[l])
			t+=lm-height[l]
			l++
		}else{
			rm=max(rm,height[r])
			t+=rm-height[r]
			r--
		}
	}
	return t
}
