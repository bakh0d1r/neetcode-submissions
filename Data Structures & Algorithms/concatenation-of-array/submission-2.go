func getConcatenation(nums []int) []int {
    lenNums:=len(nums)
    ans:=make([]int,2*lenNums,2*lenNums)
    for i,v:=range nums{
        ans[i]=v
        ans[i+lenNums]=v
    }
    return ans
}
