type TimeMap struct {
    store map[string]map[int][]string
}

func Constructor() TimeMap {
    return TimeMap{
        store: make(map[string]map[int][]string),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    if _,exists := this.store[key];!exists{
        this.store[key]=make(map[int][]string)
    }
    this.store[key][timestamp]=append(this.store[key][timestamp],value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
    if _,exists := this.store[key];!exists{
        return ""
    }
    seen:=0
    for time:=range this.store[key]{
        if time <= timestamp {
            seen = max(seen,time)
        }
    }
    if seen == 0 {
        return ""
    }
    values := this.store[key][seen]
    return values[len(values)-1]
}

