func partition(s string) [][]string {
	var res [][]string
	var path []string
	backtrack(s, 0, path, &res)
	return res
}

func backtrack(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for end := start; end < len(s); end++ {
		if isPalindrome(s, start, end) {
			path = append(path, s[start:end+1])
			backtrack(s, end+1, path, res)
			path = path[:len(path)-1] 
		}
	}
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}