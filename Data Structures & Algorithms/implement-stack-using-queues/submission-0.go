type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

// push(3) da:
// 1. q2 ga yangi element qo'sh: q2=[3]
// 2. q1 dagi barcha elementlarni q2 ga ko'chir: q2=[3,1,2]
// 3. q1 va q2 ni almashtir (swap): q1=[3,1,2], q2=[]
func (this *MyStack) Push(x int) {
	q1 := this.q1
	q2 := this.q2
	q2 = append(q2, x)
	q2 = append(q2, q1...)
	this.q1 = q2
	this.q2 = []int{}
}

func (this *MyStack) Pop() int {
	q1 := this.q1
	if this.Empty() {
		return 0
	}
	top := q1[0]
	this.q1 = q1[1:]
	return top
}

func (this *MyStack) Top() int {
	q1 := this.q1
	if this.Empty() {
		return 0
	}
	top := q1[0]
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */