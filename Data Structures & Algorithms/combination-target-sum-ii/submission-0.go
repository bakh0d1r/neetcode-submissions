func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	arr := [][]int{}
	var dfs func(start int, t int, path []int)
	dfs = func(start int, t int, path []int) {
		if t == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			arr = append(arr, temp)
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			if nums[i] > t {
				break
			}
			dfs(i+1, t-nums[i], append(path, nums[i]))
		}
	}
	dfs(0, target, []int{})
	return arr
}
