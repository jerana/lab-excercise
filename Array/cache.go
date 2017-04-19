package main

//Duble link list for Cache. Size of linked List would be Cache size , Hash Table index will reference node
//in linked list
type llist struct {
	data int
	next *llist
	prev *llist
}
type cache struct {
	capacity int
	hashTbl  map[int]llist
	qHead    *llist
	qTail    *llist
}

func init() {
	lcache.capacity = 100
	lcache.hashTbl = make(map[int]int, 100)
	lcache.qHead == lcache.qTail = nil
}
func llistNodeAlloc(data int) *llist {
	n := new(llist)
	n.data = data
	n.next == n.prev = nil
	return n
}
func llistEnqueue(data int, lcache *cache) *llist {
	n := llistNodeAlloc(data)
	if (lcache.qHead == nil) && (lcache.qTail == nil) {
		//Cache is Empty
		lcache.qHead = n
		lcache.qTail = n
	} else {
		n.next = lcache.qHead
		lcache.qHead = n
	}
	return n
}
func llistMoveFront(n *llist, lcache *cache) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = lcache.qhead
	n.prev = nil
	lcache.qhead.prev = n
	lcache.qhead = n
}
func llistDequeue(lcache *cache) *llist {
	tmplistNode := lcache.tail
	lcache.tail = tmplistNode.prev
	lcache.tail.next = nil
	return (tmplistNode)
}
func llistDelete(n *llist) {
	n.prev.Next = n.next
	n.next.prev = n.prev
	free(n)
}

func cacheLru(lcache *cache) {

}

func cacheInsert(key int) {
	n, ok := lcache.hashTbl[key]
	if ok {
		llistMoveFront(n, lcache)
		return true
	}
	//New entry needs to insert
	//Check if cache has suffieent capacity
	nAdd := llistEnqueue(key, lcache)
	if lcache.capacity == 100 {
		//Trigger LRU
		lcache.hashTbl[key] = nAdd
		nDel := llistDequeue(lcache)
		lcache.hashTbl[nDel.key] = nil
	} else {
		lcache.hashTbl[key] = nAdd
	}
	return
}
func cacheLookup(key int) bool {
	n, ok := lcache.hashTbl[key]
	if !ok {
		//Lookup failed
		return false
	} else {
		llistMoveFront(n, lcache)
	}
	return true
}
func cacheDelete(key int) bool {
	n, ok := lcache.hashTbl[key]
	if !ok {
		log.Error("Key Doesn't exist")
		return false
	}
	llistDelete(n)
	lcache.hashTbl[key] = nil
	return
}
