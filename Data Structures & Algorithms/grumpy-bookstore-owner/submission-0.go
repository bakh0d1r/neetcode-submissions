func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    b := 0
    e := 0
    me := 0
    l := 0

	for r:=range customers{
		if grumpy[r] == 0 {
			b+=customers[r]
		}

		if grumpy[r] == 1 {
			e+=customers[r]
		}

		if r-l+1 > minutes{
			if grumpy[l] == 1 {
                e -= customers[l]
            }
			l++
		}

		me=max(me,e)
	}

	return b+me
}
