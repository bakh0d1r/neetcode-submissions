func totalFruit(fruits []int) int {
	set := map[int]int{}
	l := 0
	r := 0
	mx := 0
	for r < len(fruits) {
		set[fruits[r]]++
		for len(set) > 2 {
			set[fruits[l]]--
			if set[fruits[l]] == 0 {
				delete(set, fruits[l])
			}
            l++
		}
		mx = max(mx, r-l+1)
		r++
	}
	return mx
}