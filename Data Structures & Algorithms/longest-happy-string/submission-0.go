func longestDiverseString(a int, b int, c int) string {
	ut := []Str{}
	if a > 0 {
		ut = append(ut, Str{char: 'a', count: a})
	}
	if b > 0 {
		ut = append(ut, Str{char: 'b', count: b})
	}
	if c > 0 {
		ut = append(ut, Str{char: 'c', count: c})
	}

	sh := (*StrHeap)(&ut)
	heap.Init(sh)
	str := []byte{}
	var p *Str
	for sh.Len() > 0 {
        t := heap.Pop(sh).(Str)

        str = append(str, t.char)
        t.count--

        if p != nil {
            if p.count > 0 {
                heap.Push(sh, *p)
            }
            p = nil
        }

       
        n := len(str)
        if n >= 2 && str[n-1] == str[n-2] {
            if t.count > 0 {
                p = &t
            }
        } else {
            if t.count > 0 {
                heap.Push(sh, t)
            }
        }
    }

	return string(str)
}

type Str struct {
	char  byte
	count int
}

type StrHeap []Str

func (h StrHeap) Len() int           { return len(h) }
func (h StrHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h StrHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StrHeap) Push(x any) {
	*h = append(*h, x.(Str))
}

func (h *StrHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}