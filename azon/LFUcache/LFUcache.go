package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
It requires three data structures. One is a hash table which is used to cache the key/values so that given a key we can retrieve the cache entry at O(1). Second one is a double linked list for each frequency of access. The max frequency is capped at the cache size to avoid creating more and more frequency list entries. If we have a cache of max size 4 then we will end up with 4 different frequencies. Each frequency will have a double linked list to keep track of the cache entries belonging to that particular frequency.
The third data structure would be to somehow link these frequencies lists. It can be either an array or another linked list so that on accessing a cache entry it can be easily promoted to the next frequency list in time O(1). In our article it is based on array as traversing would be faster than linked list.

*/
type entry struct {
	freq int
	key  int
	val  int
	next *entry
	prev *entry
}
type freqList []freqQ
type lfuCache_t struct {
	factor    float32
	capacity  int
	hashTable map[int]*entry
	freq      []freqList //list of list , index by freq cnt
}
type freqQ struct {
	head *entry
	tail *entry
}

var lfuCache lfuCache_t

func cache_init(capacity int) {
	lfuCache.capacity = capacity
	lfuCache.hashTable = make(map[int]*entry, 0)
	lfuCache.freq = make([]freqList, 0)
}

func allocEntry(key, val int) *entry {
	n := new(entry)
	n.key = key
	n.val = val
	n.freq = 1
	return n
}

func addNodeIntoFreqList(n *entry, freqNum int) {
	prevH := lfuCache.freq[freqNum-1]
	entry.next = prevH
	prevH.prev = entry
	lfuCache.freq[freqNum-1] = entry
}
func rmNodeFromFreqList(n *entry) {
	f := n.freq
	tPrev = n.prev
	tNext = n.next
	if tPrev {
		tPrev.next = tNext
	}
	if tNext {
		tNext.prev = tPrev
	}
	if tPrev && tNext {
		//It was only single node
		lfuCache.freq[f-1] = nil
	}
}

func cacheAdd(val int) {
	node, ok := lfuCache.hashTable[val]
	if !ok { //if Entries doesn't exist
		if len(lfuCache.hashTable) == lfuCache.capacity { //Reach Cache limit, remove entries
			rmEntries := int(float32(capacity) * factor)
			//Remove rmEntries from cache before adding new one
			removeEntries(rmEntries)
		}
		n := allocEntry(key, val)
		addNodeIntoFreqList(n, n.freq)
	} else { //entries exist
		fmt.Println("Node existing freq:", node.freq)
		//Move the entries into New freq bucket
		moveIntoHigherFreq(node)
	}
	return
}
func moveIntoHigherFreq(node *entry) {
	f := node.freq
	rmNodeFromFreqList(node, f)
	node.freq++
	if node.freq > lfuCache.capacity {
		//Put capsize
		node.freq = lfuCache.capacity
	}
	addNodeIntoFreqList(node, node.freq)
	return
}
func getVal(val int) bool {
	node, ok := lfuCache.hashTable[val]
	if !ok {
		return false
	} else {
		//Move the Node into higher freq

	}
}
func removeEtries(rmCnt int) {
	//Walk freq list
	for i := 0; i < len(lfuCache.freq); i++ {
		for rmCnt {
			n := lfuCache.freq[i]
			n.tail = n.tail.prev
			rmCnt--
			if n.tail == nil {
				//all Node from list is deleted
				break
			} else {
				n.tail.next = nil
			}
		}
	}
}
