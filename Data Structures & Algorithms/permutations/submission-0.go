func permute(nums []int) [][]int {
	arr := [][]int{}
	used := make([]bool, len(nums))
	arr = backTracking(nums, used, 1, arr, []int{})
	return arr
}

func backTracking(nums []int, used []bool, i int, arr [][]int, path []int) [][]int {
	if len(path) == len(nums) {
		subset := append([]int{}, path...)
		arr = append(arr, subset)
		return arr
	}

	for i := range nums {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		arr = backTracking(nums, used, i+1, arr, path)
		used[i] = false
		path = path[:len(path)-1]
	}
	return arr
}