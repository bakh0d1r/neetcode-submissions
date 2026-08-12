type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	Val      byte
	Children [26]*TrieNode
	IsEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			Children: [26]*TrieNode{},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
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
	}
	curr.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.root, word, 0)
}

func dfs(root *TrieNode, word string, i int) bool {
	if root == nil {
		return false
	}
	if i == len(word) {
		return root.IsEnd
	}
	if word[i] == '.' {
		for _, child := range root.Children {
			if child != nil &&  dfs(child, word, i+1) {
				return true
			}
		}
        return false
	}
	if root.Children[word[i]-'a'] == nil {
		return false
	}
	return dfs(root.Children[word[i]-'a'], word, i+1)
}
