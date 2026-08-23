type Tweet struct {
	id   int
	time int
	next *Tweet
}

type Twitter struct {
	timestamp int
	followers map[int]map[int]bool
	tweets    map[int]*Tweet
}

func Constructor() Twitter {
	return Twitter{
		timestamp: 0,
		followers: make(map[int]map[int]bool),
		tweets:    make(map[int]*Tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.timestamp++
	newTweet := &Tweet{
		id:   tweetId,
		time: this.timestamp,
		next: this.tweets[userId],
	}
	this.tweets[userId] = newTweet
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &TweetHeap{}
	heap.Init(h)

	if head, ok := this.tweets[userId]; ok && head != nil {
		heap.Push(h, head)
	}

	for followeeId := range this.followers[userId] {
		if followeeId != userId {
			if head, ok := this.tweets[followeeId]; ok && head != nil {
				heap.Push(h, head)
			}
		}
	}

	var res []int
	for h.Len() > 0 && len(res) < 10 {
		top := heap.Pop(h).(*Tweet)
		res = append(res, top.id)
		if top.next != nil {
			heap.Push(h, top.next)
		}
	}

	return res
}

func (this *Twitter) Follow(userId int, followeeId int) {
	if userId == followeeId {
		return
	}
	if _, ok := this.followers[userId]; !ok {
		this.followers[userId] = make(map[int]bool)
	}
	this.followers[userId][followeeId] = true
}

func (this *Twitter) Unfollow(userId int, followeeId int) {
	if _, ok := this.followers[userId]; ok {
		delete(this.followers[userId], followeeId)
	}
}


type TweetHeap []*Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(*Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}