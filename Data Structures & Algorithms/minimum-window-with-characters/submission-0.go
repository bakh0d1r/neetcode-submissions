func minWindow(s string, t string) string {
    need := map[byte]int{}
	have := map[byte]int{}

	for i := range t {
		need[t[i]]++
	}
	required := len(need)
	formed:=0
	l:=0
	res:=""
	for r:= range s {
		have[s[r]]++
		if have[s[r]] == need[s[r]] {
			formed++
		}
		for formed == required{
			if res == "" || r-l+1 < len(res) {
				res = s[l:r+1]
			}
			have[s[l]]--
			if have[s[l]] < need[s[l]] {
				formed--
			}
			l++
		}
	}
	return res
}
