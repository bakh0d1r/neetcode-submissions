type NestedIterator struct {
	stack []NestedInteger
}

func Constructor(nestedList []NestedInteger) *NestedIterator {
	it := &NestedIterator{}
	for i := len(nestedList) - 1; i >= 0; i-- {
		it.stack = append(it.stack, nestedList[i])
	}
	return it
}

func (this *NestedIterator) Next() int {
	// Ensure the top of the stack is an integer
	if this.HasNext() {
		curr := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1] // Pop from stack
		return curr.GetInteger()
	}
	return 0
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		curr := this.stack[len(this.stack)-1]
		
		// If the top element is an integer, we are ready
		if curr.IsInteger() {
			return true
		}
		
		// If it's a list, pop it and push its contents in reverse order
		this.stack = this.stack[:len(this.stack)-1]
		list := curr.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			this.stack = append(this.stack, list[i])
		}
	}
	return false
}