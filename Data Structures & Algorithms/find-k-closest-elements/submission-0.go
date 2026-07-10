func findClosestElements(arr []int, k int, x int) []int {
	l:=0
	for l+k<len(arr) {
		if x-arr[l] > arr[l+k]-x {
			l++
		} else {
			break
		}
	}
	return arr[l:l+k]
}
