import (
	"cmp"
	"slices"
)
func carPooling(trips [][]int, capacity int) bool {
	slices.SortFunc(trips, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	ot := []Order{}
	oh := (*OrderHeap)(&ot)
	heap.Init(oh)
	for _, trip := range trips {
		count, from, to := trip[0], trip[1], trip[2]
		for oh.Len() > 0 && (*oh)[0].to <= from {
			o := heap.Pop(oh).(Order)
			capacity += o.count
		}
		capacity -= count
		if capacity < 0 {
			return false
		}
		heap.Push(oh, Order{count: count, from: from, to: to})
	}
	return true
}

type Order struct {
	count int
	from  int
	to int
}

type OrderHeap []Order

func (h OrderHeap) Len() int           { return len(h) }
func (h OrderHeap) Less(i, j int) bool { return h[i].to < h[j].to }
func (h OrderHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *OrderHeap) Push(x any) {
	*h = append(*h, x.(Order))
}

func (h *OrderHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}