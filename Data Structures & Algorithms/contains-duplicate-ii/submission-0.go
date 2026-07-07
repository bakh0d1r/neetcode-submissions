func containsNearbyDuplicate(nums []int, k int) bool {
  set:=map[int]bool{}
  l:=0
  for i:=range nums {
	if set[nums[i]] {
		return true
	}
	set[nums[i]]=true
	if i-l>=k{
		delete(set,nums[l])
		l++
	}
  }
  return false
}
