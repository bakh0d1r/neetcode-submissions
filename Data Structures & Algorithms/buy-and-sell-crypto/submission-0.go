func maxProfit(prices []int) int {
    l,r := 0,1
    maxProfit:=0
    for r < len(prices) {
        if prices[l] < prices[r] {
            profit := prices[r]-prices[l]
            maxProfit=max(profit,maxProfit)
        }else{
            l = r
        }
        r++
    }
    return maxProfit
}
