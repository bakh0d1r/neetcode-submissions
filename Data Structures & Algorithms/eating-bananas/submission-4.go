func minEatingSpeed(piles []int, h int) int {
    l,r:=1,0
    for i := range piles {
        r = max(piles[i],r)
    }
    res := r
    for l <= r {
        k:=  (r + l)/2
        tt:=0
        for i :=range piles {
            tt+=int(math.Ceil(float64(piles[i])/float64(k)))
        }
        if tt <= h {
            res = k
            r = k - 1
        }else{
            l = k + 1
        }
    }
    return res 
}
