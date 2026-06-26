type Node struct {
    Val int
    Prev *Node
    Min int
}
type MinStack struct {
    Head *Node
}

func Constructor() MinStack {
    return MinStack{
        Head: &Node{
            Min : math.MaxInt64,
        },
    }
}

func (this *MinStack) Push(val int) {
    mn := min(this.Head.Min,val)
    this.Head = &Node{
        Val: val,
        Min: mn,
        Prev:this.Head,
    }
}

func (this *MinStack) Pop() {
    this.Head = this.Head.Prev
}

func (this *MinStack) Top() int {
    return this.Head.Val
}

func (this *MinStack) GetMin() int {
    return this.Head.Min
}
