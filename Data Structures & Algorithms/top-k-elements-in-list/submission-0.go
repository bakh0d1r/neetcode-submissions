func topKFrequent(nums []int, k int) []int {
    res := []int{}
    set:=map[int]int{}
    for i:=range nums {
       set[nums[i]]++
    }
    arr:=[][2]int{}
    for k,v :=range set {
        arr = append(arr,[2]int{v,k})
    }
    sort.Slice(arr,func(i,j int)bool {
        return arr[i][0] > arr[j][0]
    })
    for i:=0; i < k; i++ {
        res=append(res,arr[i][1])
    }
    return res
}
