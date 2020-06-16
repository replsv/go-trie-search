package main

import "sync"

// Trie data structure
type Trie struct {
	l        rune
	mutex    sync.Mutex
	children []*Trie
	leaf     bool
	dataBag  map[string]interface{}
}

// NewTrie Create a new Trie data structure
func NewTrie() *Trie {
	trie := &Trie{}
	trie.children = []*Trie{}
	trie.dataBag = make(map[string]interface{})
	return trie
}

func (t *Trie) addChild(l rune) *Trie {
	t.mutex.Lock()
	n := NewTrie()
	n.l = l
	t.children = append(t.children, n)
	t.mutex.Unlock()
	return n
}

func (t *Trie) has(l rune) (bool, *Trie) {
	for _, child := range t.children {
		if child.l == l {
			return true, child
		}
	}
	return false, nil
}

// Count count number of words in a trie
func (t *Trie) Count() int {
	c := 0
	for _, child := range t.children {
		if child.leaf == true {
			c++
		}
		c += child.Count()
	}
	return c
}

// FindNode - the note if it's a node or not
func (t *Trie) FindNode(word string) *Trie {
	letters := []rune(word)
	node := t
	i := 0
	n := len(letters)
	for i < n {
		if has, current := node.has(letters[i]); has {
			node = current
		} else {
			return nil
		}
		i++
	}
	return node
}

// Find a node pointing to a word
func (t *Trie) Find(word string) *Trie {
	node := t.FindNode(word)
	if node == nil {
		return nil
	}
	if node.leaf == true {
		return node
	}
	return nil
}

// Add word to a trie
func (t *Trie) Add(word string) *Trie {
	letters := []rune(word)
	node := t
	i := 0
	n := len(letters)
	for i < n {
		if has, current := node.has(letters[i]); has {
			node = current
		} else {
			node = node.addChild(letters[i])
		}
		i++

		if i == n {
			t.mutex.Lock()
			node.leaf = true
			t.mutex.Unlock()
		}
	}
	return node
}

// Remove a word from a trie
func (t *Trie) Remove(word string) {
	find := t.Find(word)
	if find != nil {
		t.mutex.Lock()
		find.leaf = false
		t.mutex.Unlock()
	}
}

// SetData for trie
func (t *Trie) SetData(key string, value interface{}) {
	t.dataBag[key] = value
}

// GetData from trie
func (t *Trie) GetData(key string) interface{} {
	if t == nil {
		return nil
	}

	return t.dataBag[key]
}
