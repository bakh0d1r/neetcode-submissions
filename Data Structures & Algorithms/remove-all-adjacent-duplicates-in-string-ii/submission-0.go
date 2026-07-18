func removeDuplicates(s string, k int) string {
	type str struct {
		char  byte
		count int
	}
	ss := []str{}

	for i := 0; i < len(s); i++ {
		count := 1
		if len(ss) > 0 && ss[len(ss)-1].char == s[i] {
			count += ss[len(ss)-1].count
		}
		ss = append(ss, str{
			char:  s[i],
			count: count,
		})
		if count == k {
			for range k {
				ss = ss[:len(ss)-1]
			}
		}
	}
	r := ""
	for i := range ss {
		r += string(ss[i].char)
	}
	return r
}
