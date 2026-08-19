func findKthLargest(nums []int, k int) int {
	maxH := (*IntMaxHeap)(&nums)
	heap.Init(maxH)
	n := 0
	for k > 0 && maxH.Len() > 0 {
		n = heap.Pop(maxH).(int)
		k--
	}

	return n
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}