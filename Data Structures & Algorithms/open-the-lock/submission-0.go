func openLock(deadends []string, target string) int {
	initial := "0000"
	visit := map[string]bool{}
	for i := range deadends {
		visit[deadends[i]] = true
	}
	if visit[initial] {
		return -1
	}

	type Code struct {
		Number string
		Count  int
	}

	queue := []Code{{Number: initial, Count: 0}}
	for len(queue) > 0 {
		code := queue[0]
		queue = queue[1:]
		if code.Number == target {
			return code.Count
		}
		for _, child := range children(code.Number) {
			if !visit[child] {
				visit[child] = true
				queue = append(queue, Code{Number: child, Count: code.Count + 1})
			}
		}
	}
	return -1
}

func children(code string) []string {
	s := []byte(code)
	resp := []string{}
	for i := range s {
		orig := s[i]
		s[i] = '0' + (orig-'0'+1)%10
		resp = append(resp, string(s))

		s[i] = '0' + (orig-'0'-1+10)%10
		resp = append(resp, string(s))
        s[i] = orig
	}
	return resp
}


