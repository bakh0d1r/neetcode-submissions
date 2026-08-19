func leastInterval(tasks []byte, n int) int {
	ut := []Task{}
	ht := map[byte]int{}
	for i := range tasks {
		ht[tasks[i]]++
	}
	for k, v := range ht {
		ut = append(ut, Task{kind: k, count: v})
	}

	th := (*TaskHeap)(&ut)
	heap.Init(th)
	c := 0
	for th.Len() > 0 {
		cycle := n + 1
		curr := 0
		temp := []Task{}
		for cycle > 0 && th.Len() > 0 {
			t := heap.Pop(th).(Task)
			if t.count > 0 {
				t.count = t.count - 1
				cycle--
				curr++
				if t.count > 0 {
					temp = append(temp, t)
				}
			}
		}
		for i := range temp {
			heap.Push(th, temp[i])
		}
		if th.Len() > 0 {
			c += n + 1
		} else {
			c += curr
		}
	}
	return c
}

type Task struct {
	kind  byte
	count int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}