func calPoints(operations []string) int {
	arr := []int{}
	t := 0
	for i := range operations {
		switch operations[i] {
		case "C":
			if len(arr) > 0 {
				arr = arr[:len(arr)-1]
			}
		case "D":
			if len(arr) > 0 {
				arr = append(arr, 2*arr[len(arr)-1])
			}
		case "+":
			if len(arr) > 1 {
				arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
			}
		default:
			ball, _ := strconv.Atoi(operations[i])
			arr = append(arr, ball)
		}
	}
	for i := range arr {
		t += arr[i]
	}
	return t
}
