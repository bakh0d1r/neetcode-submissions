
type TrieNode struct {
	Val      byte
	Children [26]*TrieNode
	Count    int
	IsEnd    bool
}

func insert(root *TrieNode, s string, isReverse bool) {
	curr := root
	n := len(s)

	for i := 0; i < n; i++ {
		char := s[i]
		if isReverse {
			char = s[n-1-i]
		}

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

func search(root *TrieNode, s string, isReverse bool) int {
	curr := root
	n := len(s)

	for i := 0; i < n; i++ {
		char := s[i]
		if isReverse {
			char = s[n-1-i]
		}

		idx := char - 'a'
		if curr.Children[idx] == nil {
			return 0
		}
		curr = curr.Children[idx]
	}
	return curr.Count
}

func prefixCount(words []string, pref string) int {

	c := 0
	t := &TrieNode{Children: [26]*TrieNode{}}

	for i := range words {
		insert(t, words[i], false)
	}

	c += search(t, pref, false)
	return c
}
