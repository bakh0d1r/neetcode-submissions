func rotate(nums []int, k int) {
	c:=0
	i,j:=0,0
	for c< len(nums){
		j=i
		temp:=nums[i]
		for {
			j=(j+k)%len(nums)
			temp,nums[j]=nums[j],temp
			c++
			if j==i{
				break
			}
		}
		i++
	}
}
