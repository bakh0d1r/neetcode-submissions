type TimeMap struct {
	store map[string][]Mood
}

type Mood struct {
	Type string
	Time int
}

func Constructor() TimeMap {
	return TimeMap{
		store: map[string][]Mood{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], Mood{Type: value, Time: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	moods, exist := this.store[key]
	mood := ""
	if exist {
		l := 0
		h := len(moods) - 1
		for l <= h {
			m := l + (h-l)/2
			if moods[m].Time <= timestamp {
				mood = moods[m].Type
				l = m + 1
			} else {
				h = m - 1
			}
		}
		return mood
	}
	return mood
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */