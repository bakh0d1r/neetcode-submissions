func isPerfectSquare(num int) bool {
	l := 1
	r := num
	for l <= r {
		m := l + (r-l)/2
        if m*m > num {
            r=m-1
        } else if m*m < num {
            l = m+1
        } else {
            return true
        }
	}
    return false 
}
