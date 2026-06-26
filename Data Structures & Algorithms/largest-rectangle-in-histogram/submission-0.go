func largestRectangleArea(heights []int) int {
    n := len(heights)
    ma:=0
    for i:=0;i<n;i++ {
        h:=heights[i]
        rm:=i+1
        for rm < n && heights[rm]>=h {
            rm++
        }
        lm:=i
        for lm >= 0 && heights[lm] >= h{
            lm--
        }
        rm--
        lm++
        a:=h*(rm-lm+1)
        ma=max(ma,a)
    }
    return ma
}
