func numOfSubarrays(arr []int, k int, threshold int) int {
	s := 0
	c := 0
	for r := 0; r < k; r++ {
		s += arr[r]
	}
	if s >= threshold*k  {
		c++
	}
	for r := k; r < len(arr); r++ {
		s += arr[r]
		s -= arr[r-k]
		if s >= threshold*k {
			c++
		}
	}
	return c
}