type KthLargest struct {
    k int
	heap []int
}


func Constructor(k int, nums []int) KthLargest {
    kl:= KthLargest{
		k:k,
		heap:make([]int,0,k+1),
	}
	for i:=range nums{
		kl.Add(nums[i])
	}
	return kl
}


func (this *KthLargest) Add(val int) int {
  this.heap = append(this.heap,val)
  this.siftUp(len(this.heap) - 1)

  if len(this.heap) > this.k {
	this.heap[0] = this.heap[len(this.heap) - 1]
	this.heap = this.heap[:len(this.heap)-1]
	this.siftDown(0)
  }

  return this.heap[0]
}
func (this *KthLargest) siftUp(i int) {
	parent := (i-1)/2
	for i>0 && this.heap[i] < this.heap[parent] {
		this.heap[i],this.heap[parent] = this.heap[parent],this.heap[i]
		i=parent
		parent = (i-1)/2
	}
}
func (this *KthLargest) siftDown(i int) {
	n:=len(this.heap)
	for {
		smallest:=i
		left:=2*i+1
		right:=2*i+2
		if left < n && this.heap[left] < this.heap[smallest] {
			smallest = left
		}
		if right < n && this.heap[right] < this.heap[smallest] {
			smallest = right
		}
		if smallest != i {	this.heap[i],this.heap[smallest]=this.heap[smallest],this.heap[i]
			i=smallest
		}else{
			break
		}
	}
}