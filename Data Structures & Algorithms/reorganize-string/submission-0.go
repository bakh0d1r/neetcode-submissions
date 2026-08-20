func reorganizeString(s string) string {
	ut := []Str{}
	ht := map[byte]int{}
	mx := 0
	for i := range s {
		ht[s[i]]++
		mx = max(mx, ht[s[i]])
	}
	if mx > (len(s)+1)/2 {
		return ""
	}
	for k, v := range ht {
		ut = append(ut, Str{char: k, count: v})
	}

	sh := (*StrHeap)(&ut)
	heap.Init(sh)
	str := []byte{}
		var p *Str
		for sh.Len() > 0 {
			t := heap.Pop(sh).(Str)
			if t.count > 0 {
				t.count = t.count - 1
				str = append(str, t.char)
				if p != nil {
					heap.Push(sh, *p)
				}
				p = &t
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