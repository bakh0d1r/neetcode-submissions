type Item struct {
	val int
	idx int
}

type MinHeap struct {
	data []Item
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) BuildHeap() {
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *MinHeap) Push(item Item) {
	h.data = append(h.data, item)
	h.siftUp(len(h.data) - 1)
}

func (h *MinHeap) Pop() Item {
	minItem := h.data[0]
	last := len(h.data) - 1

	h.data[0] = h.data[last]
	h.data = h.data[:last]

	if h.Len() > 0 {
		h.siftDown(0)
	}

	return minItem
}

// less returns true if item at i is strictly smaller than item at j
func (h *MinHeap) less(i, j int) bool {
	if h.data[i].val == h.data[j].val {
		return h.data[i].idx < h.data[j].idx
	}
	return h.data[i].val < h.data[j].val
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if !h.less(i, parent) {
			break
		}

		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *MinHeap) siftDown(i int) {
	n := h.Len()

	for {
		left := 2*i + 1
		if left >= n {
			break
		}

		smallest := left
		right := left + 1

		if right < n && h.less(right, left) {
			smallest = right
		}

		if !h.less(smallest, i) {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func getFinalState(nums []int, k int, multiplier int) []int {
	items := make([]Item, len(nums))
	for i, v := range nums {
		items[i] = Item{val: v, idx: i}
	}

	h := &MinHeap{data: items}
	h.BuildHeap()

	for ; k > 0; k-- {
		top := h.Pop()
		nums[top.idx] *= multiplier
		top.val = nums[top.idx]
		h.Push(top)
	}

	return nums
}