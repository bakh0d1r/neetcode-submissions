func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	arr := make([]int, n)
	l, r := 0, n-1
	for i := range nums {
		if nums[i] < pivot {
			arr[l] = nums[i]
			l++
		} 
        if nums[n-1-i] > pivot {
			arr[r] = nums[n-1-i]
			r--
		}
	}
	for l <= r {
		arr[l] = pivot
		l++
	}

	return arr
}