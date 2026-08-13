func removeSubfolders(folder []string) []string {
	trie := Constructor()
	for i := range folder {
		trie.Insert(folder[i])
	}
	arr := []string{}
	Find(trie.root, "", &arr)
	return arr
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	Val      string
	Children map[string]*TrieNode
	IsEnd    bool
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			Children: map[string]*TrieNode{},
		},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	folders := strings.Split(word, "/")
	for i := range folders {
		folder := folders[i]
		if folder == "" {
			continue
		}
		child := curr.Children[folder]

		if child == nil {
			child = &TrieNode{
				Val:      folder,
				Children: map[string]*TrieNode{},
			}
			curr.Children[folder] = child
		}
		curr = child
	}
	curr.IsEnd = true
}

func Find(curr *TrieNode, path string, arr *[]string) {
	if curr == nil {
		return
	}
	if curr.Val != "" {
		path = path + "/" + curr.Val
	}
	if curr.IsEnd {
		*arr = append(*arr, path)
		return
	}
	for k := range curr.Children {
		Find(curr.Children[k], path, arr)
	}
}
