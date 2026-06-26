func maxArea(heights []int) int {
    ma:=0
    l,r:=0,len(heights)-1
    for l < r {
        a:=min(heights[l],heights[r])*(r-l)
        ma=max(a,ma)
        if heights[l] <= heights[r]{
            l++
        } else {
            r--
        }
    }
    return ma
}

