func getConcatenation(nums []int) []int {
    arr := make([]int,len(nums)*2,len(nums)*2)
	for i:=range nums{
		arr[i]=nums[i]
		arr[i+len(nums)]=nums[i]
	}
	return arr
}
