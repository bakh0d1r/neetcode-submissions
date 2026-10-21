func jump(nums []int) int {
    c:=0
	l,r:=0,0
	for r < len(nums)-1 {
		mx:=0
		for i:=l; i<=r ; i++{
			mx = max(mx,i+nums[i])
		}
		l=r+1
		r = mx
		c++
	}
	return c
}
