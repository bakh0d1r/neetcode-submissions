func majorityElement(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	arr:=[]int{}
	v1,v2:=nums[0],nums[1]
	c1,c2:=0,0
	g:=len(nums)/3
	for i:=range nums{
		if v1 == nums[i] {
			c1++
		} else if v2 == nums[i] {
			c2++
		} else if c1 == 0 {
			c1 =1
			v1=nums[i]
		} else if c2 ==0 {
			c2=1
			v2=nums[i]
		} else {
			c1--
			c2--
		}
	}
	c1,c2 = 0,0
	for i:=range nums {
		if v1 == nums[i] {
			c1++
	}
		if v2 == nums[i] {
			c2++
		}
	}
	if c1 > g {
		arr=append(arr,v1)
	}
	if c2 > g && !(c1 > g && v1 == v2) {
		arr=append(arr,v2)
	}
	return arr
}
