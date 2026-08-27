func combine(n int, k int) [][]int {
	arr := [][]int{}
	arr = backTracking(n, k, 1, arr, []int{})
	return arr
}

func backTracking(n, k, j int, arr [][]int, path []int) [][]int {
	if len(path) == k {
		subset := append([]int{}, path...)
		arr = append(arr, subset)
		return arr
	}
	for i := j; i <= n; i++ {
		path = append(path, i)
		arr = backTracking(n, k, i+1, arr, path)
		path = path[:len(path)-1]
	}
	return arr
}