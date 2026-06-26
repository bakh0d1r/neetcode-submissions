func dailyTemperatures(temperatures []int) []int {
    arr:=make([]int,len(temperatures))
    for i:= range temperatures {
        for j:=i+1; j < len(temperatures); j++{
            if temperatures[i] < temperatures[j]{
                arr[i]=j-i
                break
            }
        }
    }
    return arr
}
