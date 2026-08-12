type TrieNode struct {
	Val      byte
	Children [26]*TrieNode
	Count    int
	IsEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			Children: [26]*TrieNode{},
		},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root

	for i := range word {
		char := word[i]

		idx := char - 'a'
		if curr.Children[idx] == nil {
			curr.Children[idx] = &TrieNode{
				Val:      char,
				Children: [26]*TrieNode{},
			}
		}
		curr = curr.Children[idx]
		curr.Count = curr.Count + 1
	}
	curr.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := range word {
		char := word[i]
		idx := char - 'a'
		if curr.Children[idx] == nil {
			return false
		}
		curr = curr.Children[idx]
	}
	return curr.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root

	for i := range prefix {
		char := prefix[i]

		idx := char - 'a'
		if curr.Children[idx] == nil {
			return false
		}
		curr = curr.Children[idx]
	}
	return true
}