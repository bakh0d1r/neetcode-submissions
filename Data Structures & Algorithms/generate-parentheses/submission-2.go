func generateParenthesis(n int) []string {
	var res []string
	backtrack(n, 0, 0, []byte{}, &res)
	return res
}

func backtrack(n, open, close int, path []byte, res *[]string) {
	if open == n && close == n {
		*res = append(*res, string(path))
		return
	}

	if open < n {
		backtrack(n, open+1, close, append(path, '('), res)
	}

	if close < open {
		backtrack(n, open, close+1, append(path, ')'), res)
	}
}