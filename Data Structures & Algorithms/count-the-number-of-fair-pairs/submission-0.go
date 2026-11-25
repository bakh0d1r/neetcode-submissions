func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	countLessEqual := func(target int) int64 {
		var count int64
		l, r := 0, len(nums)-1
		for l < r {
			if nums[l]+nums[r] <= target {
				count += int64(r - l)
				l++
			} else {
				r--
			}
		}
		return count
	}

	return countLessEqual(upper) - countLessEqual(lower-1)
}