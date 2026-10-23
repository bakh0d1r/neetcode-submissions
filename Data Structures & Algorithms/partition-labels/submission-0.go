func partitionLabels(s string) []int {
    arr:= []int{}
    o := make([]int,26)
    for i := range s {
        o[s[i]-'a']=i
    }
    l,r := 0,0
    for i := range s {
        r = max(o[s[i]-'a'],r)
        if r == i {
            arr=append(arr,r-l+1)
            l=i+1
        }
    }
    return arr
}