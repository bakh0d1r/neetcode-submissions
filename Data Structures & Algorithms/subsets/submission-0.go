func subsets(nums []int) [][]int {
	arr := [][]int{}
	arr = backTracking(nums, 0, arr, []int{})
	return arr
}

func backTracking(nums []int, i int, arr [][]int, path []int) [][]int {
	if i == len(nums) {
		subset := append([]int{}, path...)
		arr = append(arr, subset)
		return arr
	}
	arr=backTracking(nums, i+1, arr, path)
	path = append(path, nums[i])
	arr=backTracking(nums, i+1, arr, path)

	path = path[:len(path)-1]
	return arr
}