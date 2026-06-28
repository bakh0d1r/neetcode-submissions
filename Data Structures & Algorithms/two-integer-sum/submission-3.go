func twoSum(nums []int, target int) []int {
   set:=map[int]int{}
   for i :=range nums{
	 set[nums[i]]=i
   }
   for i:= range nums{
	n:=target - nums[i]
	if j,ok := set[n]; ok && j!=i {
		return []int{i,j}
	}
   }
   return []int{}
}
