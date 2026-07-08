func characterReplacement(s string, k int) int {
	set:=map[byte]int{}
	l:=0
	ml:=0
  	mc := 0
	for r := range s{
		set[s[r]]++
        mc=max(mc,set[s[r]])
	
		if r - l +1 - mc > k {
			set[s[l]]--
			l++
		}
		ml = max(ml,r - l +1)
	}
	return ml
}

