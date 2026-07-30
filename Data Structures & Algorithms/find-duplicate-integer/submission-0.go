func findDuplicate(nums []int) int {
    slow:=0
	fast:=0
	for {
		slow=nums[slow]
		fast=nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast=0
	for {
		slow=nums[slow]
		fast=nums[fast]
		if slow == fast {
			return slow
		}
	}
	return slow
}
