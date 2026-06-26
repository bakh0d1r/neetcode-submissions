func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    n:= len(height)
    res:=0
    l,r:=0,n-1
    lm,rm:=height[l],height[r]
    for l < r {
        if lm < rm {
            l++
            lm=max(lm,height[l])
            res+=lm-height[l]
        }else{
            r--
            rm=max(rm,height[r])
            res+=rm-height[r]
        }
    }
    return res
}
