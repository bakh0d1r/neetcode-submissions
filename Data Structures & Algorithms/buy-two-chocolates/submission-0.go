func buyChoco(prices []int, money int) int {
    min1, min2 := math.MaxInt32, math.MaxInt32

    for _, p := range prices {
        if p < min1 {
            min2 = min1
            min1 = p
        } else if p < min2 {
            min2 = p
        }
    }

    leftover := money - min1 - min2
    if leftover >= 0 {
        return leftover
    }
    return money
}