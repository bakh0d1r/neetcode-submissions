

type Element struct {
	sum   int
	index int
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Element)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rangeSum(nums []int, n int, left int, right int) int {
	mod := 1000000007
	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		heap.Push(h, Element{sum: nums[i], index: i})
	}

	result := 0
	for i := 1; i <= right; i++ {
		curr := heap.Pop(h).(Element)

		if i >= left {
			result = (result + curr.sum) % mod
		}

		if curr.index+1 < n {
			heap.Push(h, Element{
				sum:   curr.sum + nums[curr.index+1],
				index: curr.index + 1,
			})
		}
	}

	return result
}