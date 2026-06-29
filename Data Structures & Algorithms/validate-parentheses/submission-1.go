func isValid(s string) bool {
	stack := []byte{}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == pairs[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			stack=append(stack,s[i])
		}
	}
    if len(stack) != 0 {
        return false
    }
    return true
}