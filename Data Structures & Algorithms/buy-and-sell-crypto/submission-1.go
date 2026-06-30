func maxProfit(prices []int) int {
	mp:=math.MaxInt64
	p:=0
	for i:=range prices {
		mp=min(prices[i],mp)
		p=max(p,prices[i]-mp)
	}
	return p
}
