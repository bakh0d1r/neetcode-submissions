func longestConsecutive(nums []int) int {
set:= map[int]bool{}
for i:=range nums {
    set[nums[i]]=true
}
m:=0
for k := range set {
    c:=0
  if ok:= set[k-1] ; !ok {
    c++
   for {
    if set[k+c] {
        c++
        continue
    }
    m=max(m,c)
    break
  }
  }
}
return m
}
