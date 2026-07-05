func fourSum(nums []int, target int) [][]int {
sort.Ints(nums)
arr:=[][]int{}
for i:=0;i < len(nums)-2; i++{
	if i>0 && nums[i] == nums[i-1] {
		continue
	}
	for j:=i+1 ; j<len(nums)-1; j++ {
		if j > i+1 && nums[j] == nums[j-1] {
			continue
		}
		t:= target - nums[i] - nums[j]
		l:= j+1
		r:= len(nums)-1
		for l < r {
			s:= nums[l] + nums[r]
			if s > t {
				r--
			} else if t > s {
				l++
			} else {
				arr=append(arr,[]int{nums[i],nums[j],nums[l],nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
}
return arr 
}