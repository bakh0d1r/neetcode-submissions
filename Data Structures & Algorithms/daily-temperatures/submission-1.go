func dailyTemperatures(temperatures []int) []int {
    s:=NewStack[int]()
    arr:=make([]int,len(temperatures))
	for i:=range temperatures{
		top,existTop:=s.Top()
		for !s.IsEmpty() && existTop && temperatures[top] < temperatures[i] {
			arr[top]=i-top
			s.Pop()
			top,existTop=s.Top()
		}
		s.Push(i)
	}
	return arr
}



// Stack is a generic stack implementation.
type Stack[T any] struct {
	stack []T
}

// NewStack initializes a new stack with optional initial values.
// Note: In Go, constructors are typically standalone functions, not methods.
func NewStack[T any](values ...T) *Stack[T] {
	s := &Stack[T]{
		stack: make([]T, 0, len(values)),
	}
	for _, v := range values {
		s.Push(v)
	}
	return s
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.stack = append(s.stack, v)
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.stack)
}

// IsEmpty returns true if the stack contains no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

// Last returns the top element of the stack without removing it.
// It also returns a boolean indicating if the operation was successful.
func (s *Stack[T]) Top() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	return s.stack[s.LastIdx()], true
}

// LastIdx returns the index of the top element.
func (s *Stack[T]) LastIdx() int {
	return s.Len() - 1
}

// Pop removes and returns the top element of the stack.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}

	// Get the last element
	pop := s.stack[s.LastIdx()]
	// Slice off the last element
	s.stack = s.stack[:s.LastIdx()]

	return pop, true
}

// List returns the stack.
func (s *Stack[T]) List() ([]T) {
	return s.stack
}
