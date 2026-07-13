type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(value int) {
	this.stack = append(this.stack, value)
	if len(this.minStack) == 0 || this.GetMin() >= value {
		this.minStack = append(this.minStack, value)
	}
}

func (this *MinStack) Pop() {
	popElement := 0
	if len(this.stack) > 0 {
		popElement = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	if len(this.minStack) > 0 && this.GetMin() == popElement {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) > 0 {
		return this.minStack[len(this.minStack)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
