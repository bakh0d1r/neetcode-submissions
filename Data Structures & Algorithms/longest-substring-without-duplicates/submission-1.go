func lengthOfLongestSubstring(s string) int {
	set:=map[byte]int{}
	l:=0
	m:=0
	for r:=range s {
		if j,ok:=set[s[r]] ; ok && j >= l {
			l=j+1
		}
		set[s[r]]=r
		m=max(m,r-l+1)
	}
	return m
}
