type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size {
		return -1
	}
	return this.getNode(index).Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.insertBetween(val, this.head, this.head.Next)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.insertBetween(val, this.tail.Prev, this.tail)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	var succ *Node
	if index == this.size {
		succ = this.tail
	} else {
		succ = this.getNode(index)
	}

	this.insertBetween(val, succ.Prev, succ)
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}

	target := this.getNode(index)
	target.Prev.Next = target.Next
	target.Next.Prev = target.Prev
	this.size--
}

// --- Helper Functions ---

func (this *MyLinkedList) insertBetween(val int, prev *Node, next *Node) {
	newNode := &Node{
		Val:  val,
		Prev: prev,
		Next: next,
	}
	prev.Next = newNode
	next.Prev = newNode
	this.size++
}

func (this *MyLinkedList) getNode(index int) *Node {
	if index < this.size/2 {
		curr := this.head.Next
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		return curr
	}

	curr := this.tail.Prev
	for i := this.size - 1; i > index; i-- {
		curr = curr.Prev
	}
	return curr
}