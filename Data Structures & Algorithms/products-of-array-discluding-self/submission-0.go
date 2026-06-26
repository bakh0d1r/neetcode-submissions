func productExceptSelf(nums []int) []int {
    arr := make([]int,len(nums))
    for i:=range nums{
        num:=1
        for j:=range nums{
            if i != j {
                num*=nums[j]
            }
        }
        arr[i]=num
    }
    return arr
}
