func validPalindrome(s string) bool {
	l,r:=0,len(s)-1
	for l < r {
		if s[l] != s[r] {
			return  isPalin(s, l+1,r) || isPalin(s,l,r-1)
		} 
		l++
		r--
	}
	return true
}


func isPalin(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}
