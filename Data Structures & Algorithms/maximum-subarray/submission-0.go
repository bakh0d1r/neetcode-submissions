func maxSubArray(nums []int) int {
	cmx := 0
	mx := nums[0]
	for i := range nums {
		cmx += nums[i]
		mx = max(mx, cmx)
		if cmx < 0 {
			cmx = 0
		}
	}
	return mx
}
