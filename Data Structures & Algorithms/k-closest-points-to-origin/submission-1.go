

func kClosest(points [][]int, k int) [][]int {
	h := &MinHeap{}
	heap.Init(h)
	for i := range points {
		heap.Push(h, Item{
			dis: points[i][0]*points[i][0] + points[i][1]*points[i][1],
			x:   points[i][0],
			y:   points[i][1],
		})
	}
	arr := [][]int{}
	for k > 0 && h.Len() > 0 {
		top := heap.Pop(h).(Item)
		arr = append(arr, []int{top.x, top.y})
		k--
	}

	return arr
}

type Item struct {
	dis int
	x   int
	y   int
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].dis != h[j].dis {
		return h[i].dis < h[j].dis
	}

	if h[i].x != h[j].x {
		return h[i].x < h[j].x
	}

	return h[i].y < h[j].y
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}