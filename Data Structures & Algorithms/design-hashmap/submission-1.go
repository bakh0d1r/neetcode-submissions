type MyHashMap struct {
	Store []int
}

func Constructor() MyHashMap {
    store := make([]int, 1_000_001)
	for i:=range store {
		store[i]=-1
	}
	return MyHashMap{
		Store:store,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.Store[key]=value
}

func (this *MyHashMap) Get(key int) int {
	return this.Store[key]
}

func (this *MyHashMap) Remove(key int) {
    this.Store[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */