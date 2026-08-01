type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	table    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	dummyHead := &Node{}
	dummyTail := &Node{}
	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead
	return LRUCache{
		capacity: capacity,
		head:     dummyHead,
		tail:     dummyTail,
		table:    map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.table[key]
	if node == nil {
		return -1
	}
	this.remove(node)
	this.addTail(node)
	this.table[key] = node
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node := this.table[key]
	if node != nil {
        node.val=value
		this.remove(node)
		this.addTail(node)
		this.table[key] = node
		return
	}
	node = &Node{
		key: key,
		val: value,
	}
	if this.capacity == len(this.table) {
		this.remove(this.head.next)
	}
	this.addTail(node)
	this.table[key] = node
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(this.table, node.key)
}

func (this *LRUCache) addTail(node *Node) {
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev.next = node
	this.tail.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */