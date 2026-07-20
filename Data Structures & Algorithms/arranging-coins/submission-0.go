func arrangeCoins(n int) int {
	l := 0
	r := n
	for l <= r {
		m := l + (r-l)/2
		s := (1 + m) * m / 2
		if s > n {
			r = m - 1
		} else  {
			l = m + 1
		}
	}
	return r
}