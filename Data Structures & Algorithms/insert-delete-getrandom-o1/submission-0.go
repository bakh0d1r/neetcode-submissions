type RandomizedSet struct {
    numMap map[int]int
    nums   []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        numMap: make(map[int]int),
        nums:   []int{},
    }
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.numMap[val]; exists {
        return false
    }
    this.numMap[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    if _, exists := this.numMap[val]; !exists {
        return false
    }
    idx := this.numMap[val]
    last := this.nums[len(this.nums)-1]
    this.nums[idx] = last
    this.numMap[last] = idx
    this.nums = this.nums[:len(this.nums)-1]
    delete(this.numMap, val)
    return true
}

func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))]
}