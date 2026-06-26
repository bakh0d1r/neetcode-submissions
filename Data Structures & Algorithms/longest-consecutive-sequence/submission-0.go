func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    minNum,maxNum:=math.MaxInt,math.MinInt
    set := map[int]bool{}
    for i:= range nums{
        minNum=min(minNum,nums[i])
        maxNum=max(maxNum,nums[i])
        set[nums[i]]=true
    }
    count:=0
    maxCount:=math.MinInt
    for i :=minNum;i<=maxNum;i++ {
        if set[i] && set[i+1] {
            count++
        }else{
            maxCount=max(maxCount,count)
            count=0
        }
    }
    maxCount++
    return maxCount
}
