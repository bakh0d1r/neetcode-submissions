func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := range bills {
		switch bills[i] {
		case 5:
			five++
		case 10:
			if five <= 0 {
				return false
			}
			ten++
			five--
		case 20:
            if five >0 && ten >0 {
                ten--
			    five--
            }else if five >= 3 {
                five-=3
            }else{
                return false
            }
		}
	}
	return true
}