func minCostConnectPoints(points [][]int) int {
    n := len(points)
	adj := make(map[int][][]int)
	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dist := int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
			adj[i] = append(adj[i], []int{dist, j})
			adj[j] = append(adj[j], []int{dist, i})
		}
	}
	res := 0
	visit := make(map[int]bool)
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, []int{0, 0})

	for len(visit) < n {
		if pq.Len() == 0 {
			break
		}
		item := heap.Pop(pq).([]int)
		cost, point := item[0], item[1]
		if visit[point] {
			continue
		}
		res += cost
		visit[point] = true

		for _, edge := range adj[point] {
			neiCost, neiPoint := edge[0], edge[1]
			if !visit[neiPoint] {
				heap.Push(pq, []int{neiCost, neiPoint})
			}
		}
	}
	return res
}


type MinHeap [][]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}