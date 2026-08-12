
type TrieNode struct {
	Val      byte
	Children [26]*TrieNode
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
	}
	curr.IsEnd = true
}

func search(root *TrieNode, s string, isReverse bool) bool {
	curr := root
	n := len(s)

	for i := 0; i < n; i++ {
		char := s[i]
		if isReverse {
			char = s[n-1-i]
		}

		idx := char - 'a'
		if curr.Children[idx] == nil {
			return false
		}
		curr = curr.Children[idx]
	}
	return true
}

func countPrefixSuffixPairs(words []string) int {
	c := 0
	n := len(words)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if len(words[i]) > len(words[j]) {
				continue
			}
			t := &TrieNode{Children: [26]*TrieNode{}}
			rt := &TrieNode{Children: [26]*TrieNode{}}
			insert(t, words[j], false)
			insert(rt, words[j], true)

			if search(t, words[i], false) && search(rt, words[i], true) {
				c++
			}
		}
	}

	return c
}