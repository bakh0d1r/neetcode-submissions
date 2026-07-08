func minimumDifference(nums []int, k int) int {
m:=math.MaxInt64
sort.Ints(nums)
for i:=k-1;i<len(nums);i++{
	m=min(m,nums[i]-nums[i-k+1])
}
return m
}
