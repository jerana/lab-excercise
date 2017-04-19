package main

import "fmt"

var root *Trie

func main() {
	insertTrieTree("abcd")
	insertTrieTree("abgl")
	insertTrieTree("amml")
	insertTrieTree("tmml")
	fmt.Println("Search :", "ab", startWith("ab"))
	fmt.Println("Search :", "abcd", search("abcd"))

}

//Implement a trie with insert, search, and startsWith methods.

type Trie struct {
	hash   map[rune]*Trie
	isLeaf bool
}

func trieNodeAlloc() *Trie {
	n := new(Trie)
	n.hash = make(map[rune]*Trie)
	return n
}

func insertTrieTree(s string) {
	if root == nil {
		root = trieNodeAlloc()
	}
	cur := root
	for i := 0; i < len(s)-1; i++ { //Walk into Trie
		r := rune(s[i])
		_, ok := cur.hash[r]
		if !ok {
			cur.hash[r] = trieNodeAlloc()
		}
		cur = cur.hash[r]
	}
	//Last  element of string
	cur.hash[rune(s[len(s)-1])] = trieNodeAlloc()
	cur.isLeaf = true
}
func find(s string) (P, C *Trie) {
	if root == nil {
		return nil
	}
	var p, c *Trie
	p, c = root, root

	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		_, ok := c.hash[r]
		if ok {
			p = c
			c = c.hash[r]

		} else {
			return nil
		}
	}
	return p, c
}
func search(word string) bool {
	t, _ := find(word)
	if t != nil && t.isLeaf {
		return true
	}
	return false
}
func startWith(pfx string) bool {
	t, c := find(pfx)
	if t != nil {
		return true
	}
	return false
}
func delete(word string) bool {
	var p, c *Trie
	if root == nil {
		return false
	}
	for i := len(s); i >= 0; i-- {
		p, c := find(word[:i])
		if p != nil {
			delete(p.hash[word[i]])
		}
	}

}

/*PrintSuggestion*/
func printSuggestion(root *Trie, prefix string) {
	if root == nil || prefix == nil || len(root.hash) == 0 {
		return
	}
	parent := root
	for i := 0; i < len(prefix); i++ {
		c = prefix[i]
		next := parent.hash[c]
		if n == nil {
			return
		}
		parent = next
	}
	printSorted(parent, prefix)
}
func printSorted(root *Trie, s string) {
	if root.isLeaf {
		fmt.Println(s)
	}
	for key, node := range root.hash {
		printSorted(node, s+key)

	}

}
