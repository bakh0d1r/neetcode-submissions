type MyHashSet struct {
    Store [1_000_001]bool
}


func Constructor() MyHashSet {
    return MyHashSet{}
}


func (this *MyHashSet) Add(key int)  {
    this.Store[key]=true
}


func (this *MyHashSet) Remove(key int)  {
      this.Store[key]=false
}


func (this *MyHashSet) Contains(key int) bool {
    return  this.Store[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */