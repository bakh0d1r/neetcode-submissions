func pickGifts(gifts []int, k int) int64 {
	h := &MaxHeap{data: gifts}
	h.buildHeap()

	for k > 0 {
		p := h.Pop()
		h.Push(int(math.Sqrt(float64(p))))
        k--
	}
    var t int64
	for h.Len() > 0 {
		t+=int64(h.Pop())
	}

	return t
}

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) Len() int {
	return len(h.data)
}

func (h *MaxHeap) buildHeap() {
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *MaxHeap) Push(val int) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap) Pop() int {
	max := h.data[0]
	last := len(h.data) - 1

	h.data[0] = h.data[last]
	h.data = h.data[:last]

	if h.Len() > 0 {
		h.siftDown(0)
	}

	return max
}

func (h *MaxHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if h.data[i] <= h.data[parent] {
			break
		}

		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *MaxHeap) siftDown(i int) {
	n := h.Len()

	for {
		left := 2*i + 1
		if left >= n {
			break
		}

		largest := left
		right := left + 1

		if right < n && h.data[right] > h.data[left] {
			largest = right
		}

		if h.data[i] >= h.data[largest] {
			break
		}

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}