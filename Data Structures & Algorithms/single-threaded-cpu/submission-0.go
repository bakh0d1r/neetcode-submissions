
type Task struct {
	enqueueTime    int
	processingTime int
	index          int
}

type TaskHeap []Task

func (h TaskHeap) Len() int { return len(h) }
func (h TaskHeap) Less(i, j int) bool {
	if h[i].processingTime == h[j].processingTime {
		return h[i].index < h[j].index
	}
	return h[i].processingTime < h[j].processingTime
}
func (h TaskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

func getOrder(tasks [][]int) []int {
	n := len(tasks)
	extendedTasks := make([]Task, n)
	for i, t := range tasks {
		extendedTasks[i] = Task{
			enqueueTime:    t[0],
			processingTime: t[1],
			index:          i,
		}
	}

	sort.Slice(extendedTasks, func(i, j int) bool {
		return extendedTasks[i].enqueueTime < extendedTasks[j].enqueueTime
	})

	h := &TaskHeap{}
	heap.Init(h)

	result := make([]int, 0, n)
	currentTime := 0
	taskIdx := 0

	for taskIdx < n || h.Len() > 0 {
		if h.Len() == 0 && currentTime < extendedTasks[taskIdx].enqueueTime {
			currentTime = extendedTasks[taskIdx].enqueueTime
		}

		for taskIdx < n && extendedTasks[taskIdx].enqueueTime <= currentTime {
			heap.Push(h, extendedTasks[taskIdx])
			taskIdx++
		}

		currentTask := heap.Pop(h).(Task)
		currentTime += currentTask.processingTime
		result = append(result, currentTask.index)
	}

	return result
}