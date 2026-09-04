func subsetsWithDup(nums []int) [][]int {
    arr := [][]int{}
    sort.Ints(nums)
    return backTracking(nums, 0, arr, []int{})
}

func backTracking(nums []int, i int, arr [][]int, path []int) [][]int {
    if i == len(nums) {
        subset := append([]int{}, path...)
        return append(arr, subset)
    }

    path = append(path, nums[i])
    arr = backTracking(nums, i+1, arr, path)
    path = path[:len(path)-1]

    for i+1 < len(nums) && nums[i] == nums[i+1] {
        i++
    }

    arr = backTracking(nums, i+1, arr, path)

    return arr
}