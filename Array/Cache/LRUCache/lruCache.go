package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type node struct {
	key int
	s   string
}
type LruCache struct {
	maxSize uint64
	size    uint64
	list    []node
	hashMap [int]int
}

var cache LruCache

func init() {
	cache.maxSize = 1000
	cache.list = make([]node, 0)
	cache.hashMap = make(map[int]int, 0)
}
func lookup(key int) (string, error) {
	n, ok := cache.hashMap[key]
	if !ok {
		return "", errors.New("Key doesn't found")
	}

	s := cache.list[n].s
	moveNodeToHead(n)
	return s, nil
}
func put(k int, str string) {
	n, ok := cache.hashMap[k]
	if ok {
		moveNodeToHead(n)
		return
	}
	cache.list = append([]node{key: k, s: str}, cache.list...)

}
