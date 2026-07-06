func sortArray(nums []int) []int {
    if len(nums) <= 1 {
		return nums
	}
	m:= len(nums)/2
	ln:=sortArray(nums[:m])
	rn:=sortArray(nums[m:])
	return mergeSort(ln,rn)
}


func mergeSort(ln,rn []int) []int {
	l,r:= 0,0
	arr:=make([]int, 0, len(ln)+len(rn))
	for l<len(ln) && r < len(rn) {
		if ln[l] <= rn[r] {
			arr=append(arr,ln[l])
			l++
		}else{
			arr=append(arr,rn[r])
			r++
		}
	}
	arr=append(arr,ln[l:]...)
	arr=append(arr,rn[r:]...)
	return arr
}
