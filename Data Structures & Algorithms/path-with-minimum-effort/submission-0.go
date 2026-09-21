func minimumEffortPath(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = 1<<31 - 1
		}
	}
	dist[0][0] = 0
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	minHeap := &MinHeap{Item{0, 0, 0}}
	heap.Init(minHeap)
	for minHeap.Len() > 0 {
		curr := heap.Pop(minHeap).(Item)
		diff, r, c := curr.diff, curr.row, curr.column
		if r == rows-1 && c == cols-1 {
			return diff
		}
		if dist[r][c] < diff {
			continue
		}
		for i := range directions {
			dir := directions[i]
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}
			nd := max(diff, abs(heights[r][c]-heights[nr][nc]))
			if nd < dist[nr][nc] {
				dist[nr][nc] = nd
				heap.Push(minHeap, Item{nd, nr, nc})
			}
		}
	}
	return 0
}

type Item struct {
	diff   int
	row    int
	column int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].diff < h[j].diff }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}