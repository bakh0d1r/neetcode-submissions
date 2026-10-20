func minMovesToSeat(seats []int, students []int) int {
	maxSeat, maxStudent := 0, 0
    for _, s := range seats {
        if s > maxSeat {
            maxSeat = s
        }
    }
    for _, s := range students {
        if s > maxStudent {
            maxStudent = s
        }
    }

    countSeats := make([]int, maxSeat+1)
    countStudents := make([]int, maxStudent+1)

    for _, s := range seats {
        countSeats[s]++
    }
    for _, s := range students {
        countStudents[s]++
    }

    remain, i, j, res := len(seats), 0, 0, 0
	for remain > 0 {
		if countSeats[i] == 0 {
			i++ 
			continue
		}
		if countStudents[j] == 0 {
			j++ 
			continue
		}
		tmp:=min(countSeats[i],countStudents[j])
		diff:=max(j,i)-min(i,j)
		res+=diff*tmp
		countSeats[i]-=tmp
		countStudents[j]-=tmp
		remain -=tmp
	}
	return res
}
