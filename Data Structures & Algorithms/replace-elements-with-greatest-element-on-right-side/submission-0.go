func replaceElements(arr []int) []int {
	m:=-1
	ra:=make([]int,len(arr))
	for r:=len(arr)-1;r>=0;r--{
		ra[r] = m
		m=max(m,arr[r])
	}
	return ra
}
