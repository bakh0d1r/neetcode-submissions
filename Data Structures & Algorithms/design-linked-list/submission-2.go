type MyLinkedList struct {
    head *Node
	tail *Node
	size int
}

type Node struct {
	Next *Node
	Prev *Node
	Val int
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

func (this *MyLinkedList) AddAtHead(val int)  {
    this.insertBetween(val,this.head,this.head.Next)
}

func (this *MyLinkedList) AddAtTail(val int)  {
    this.insertBetween(val,this.tail.Prev,this.tail)
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.size {
		return
	}
	var node *Node
	if index == this.size{
		node=this.tail
	}else{
		node = this.getNode(index)
	}
	this.insertBetween(val,node.Prev,node)
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.size {
		return
	}
	node:=this.getNode(index)
	node.Prev.Next=node.Next
	node.Next.Prev=node.Prev
	this.size--
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

func (this *MyLinkedList) insertBetween(val int, prev *Node, next *Node) {
	new:=&Node{
		Val:val,
		Prev:prev,
		Next:next,
	}
	prev.Next=new
	next.Prev=new
	this.size++
}