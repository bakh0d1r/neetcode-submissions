func subarraySum(nums []int, k int) int {
	c:=0
	ps:=0
	ht:=map[int]int{0:1}
	for i := range nums {
		ps+=nums[i]
		c+=ht[ps-k]
		ht[ps]++
	}
	return c
}
