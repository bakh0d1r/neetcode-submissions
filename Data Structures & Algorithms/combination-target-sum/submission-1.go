func combinationSum(candidates []int, target int) [][]int {
	var arr [][]int

	arr = backtracking(candidates, 0, []int{}, target, target, arr)
	return arr
}

func backtracking(nums []int, start int, path []int, remain int, target int, arr [][]int) [][]int {
	if remain == 0 {
		new := make([]int, len(path))
		copy(new, path)
		arr = append(arr, new)
		return arr
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		remain := remain - nums[i]

		if remain < 0 {
			path = path[:len(path)-1]
            continue
		}

		arr = backtracking(nums, i, path, remain, target, arr)

		path = path[:len(path)-1]
	}

	return arr
}