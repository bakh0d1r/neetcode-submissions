func minEatingSpeed(piles []int, h int) int {
	r := -1
	for i := range piles {
		r = max(r, piles[i])
	}
	l := 1
	for l <= r {
		m := l + (r-l)/2
		t := 0
		for i := range piles {
			t += (piles[i] + m - 1) / m
		}
		if t > h {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}