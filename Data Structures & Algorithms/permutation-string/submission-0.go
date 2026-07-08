func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	set:=map[byte]int{}
	for i:=range s1{
		set[s1[i]]++
	}
	n:=len(s1)
	for i:=0;i<n;i++{
		set[s2[i]]--
	}
	isPermutation:=true
	for _,v:=range set{
		if v != 0 {
			isPermutation=false
		}
	}
	if isPermutation {
		return true
	}
	for l:=1;l<len(s2)-n+1;l++ {
		r:=l+n
		set[s2[l-1]]++
		set[s2[r-1]]--
		isPermutation:=true
		for _,v:=range set{
			if v != 0 {
				isPermutation=false
			}
		}
		if isPermutation {
			return true
		}
	}
	return false
}
