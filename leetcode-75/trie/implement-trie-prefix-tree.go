package trie

type Trie struct {
	chars [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	var trie Trie
	return trie
}

func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		char := word[i] - 'a'
		if this.chars[char] == nil {
			this.chars[char] = new(Trie)
		}
		this = this.chars[char]
	}
	this.isEnd = true
}

func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		char := word[i] - 'a'
		if this.chars[char] == nil {
			return false
		}
		this = this.chars[char]
	}
	if this.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		char := prefix[i] - 'a'
		if this.chars[char] == nil {
			return false
		} else {
			this = this.chars[char]
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
