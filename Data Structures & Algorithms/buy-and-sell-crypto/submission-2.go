func maxProfit(prices []int) int {
	p := 0
	l, r := 0, 1

	for r < len(prices) {
		if prices[r] > prices[l] {
			p = max(p, prices[r]-prices[l])
		} else {
			l = r
		}
		r++
	}
	return p
}