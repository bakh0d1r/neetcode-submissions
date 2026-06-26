func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    n:= len(height)
    res:=0
    for i:=0;i<n;i++{
        lm:=height[i]
        rm:=height[i]
        for j:=0;j<i;j++{
            lm=max(lm,height[j])
        }
        for j:=i+1;j<n;j++{
            rm=max(rm,height[j])
        }
        res += min(lm,rm)-height[i]
    }
    return res
}
