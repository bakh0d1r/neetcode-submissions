type Node struct {
    url  string
    prev *Node
    next *Node
}

type BrowserHistory struct {
    curr *Node
}

func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        curr: &Node{url: homepage},
    }
}

func (this *BrowserHistory) Visit(url string) {
    newNode := &Node{
        url:  url,
        prev: this.curr,
    }
    this.curr.next = newNode
    this.curr = newNode
}

func (this *BrowserHistory) Back(steps int) string {
    for this.curr.prev != nil && steps > 0 {
        this.curr = this.curr.prev
        steps--
    }
    return this.curr.url
}

func (this *BrowserHistory) Forward(steps int) string {
    for this.curr.next != nil && steps > 0 {
        this.curr = this.curr.next
        steps--
    }
    return this.curr.url
}