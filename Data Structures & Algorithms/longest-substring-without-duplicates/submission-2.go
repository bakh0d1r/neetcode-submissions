func lengthOfLongestSubstring(s string) int {
	set:=map[byte]bool{}
	l:=0
	m:=0
	for r:=range s {
		for set[s[r]] {
			delete(set, s[l])
			l++
		}
		set[s[r]]=true
		m=max(m,r-l+1)
	}
	return m
}
