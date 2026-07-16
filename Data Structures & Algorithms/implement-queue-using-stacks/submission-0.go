type MyQueue struct {
	buffer []int // New elements are pushed here
	queue  []int // Stack that keeps the queue front ready
}

func Constructor() MyQueue {
	return MyQueue{}
}

// Add an element to the queue.
// Stack push = append()
func (q *MyQueue) Push(x int) {
	q.buffer = append(q.buffer, x)
}

// Move elements from buffer to queue stack.
func (q *MyQueue) move() {
	// Transfer only when queue stack is empty.
	if len(q.queue) == 0 {

		for len(q.buffer) > 0 {

			// ----- Stack Pop -----
			n := len(q.buffer)
			x := q.buffer[n-1]      // top element
			q.buffer = q.buffer[:n-1] // pop element
			// ---------------------

			// ----- Stack Push -----
			q.queue = append(q.queue, x)
			// ----------------------
		}
	}
}

// Remove and return the front element.
func (q *MyQueue) Pop() int {
	q.move()

	// queue stack top == queue front

	// ----- Stack Pop -----
	n := len(q.queue)
	x := q.queue[n-1]
	q.queue = q.queue[:n-1]
	// ---------------------

	return x
}

// Return the front element without removing it.
func (q *MyQueue) Peek() int {
	q.move()

	// Stack top
	return q.queue[len(q.queue)-1]
}

// Check whether the queue is empty.
func (q *MyQueue) Empty() bool {
	return len(q.buffer) == 0 &&
		len(q.queue) == 0
}