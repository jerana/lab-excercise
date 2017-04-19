package main

import (
	"fmt"
	"sort"
)

func isPowerOfTwo(n uint64) bool {
	if n <= 0 {
		return false
	}
	return ((n & (n - 1)) == 0)
}
func findLargestInclusiveInterval(arr []int) []int {
	var intVal []int = []int{-1, -1}
	for i := 0; i < len(arr)-1; i++ {
		intVal[0] = i
		for j := 0; i < len(arr)-1; j++ {
			if (arr[j] + 1) == arr[j+1] {
				intVal[1] = j + 1
			} else {
				break
			}
		}
	}
	return intVal
}
func findLargestInterval(arr []int) []int {
	m := make(map[int]bool)
	for _, idx := range arr {
		m[idx] = true
	}
	var maxIntVal = []int{0, 0}
	var y int
	for i := 0; i < len(arr)-1; i++ {
		x := arr[i]
		for {
			_, ok := m[x+1]
			if ok {
				x = x + 1
			} else {
				break
			}
		}
		y = x
		x = arr[i]
		if (maxIntVal[1] - maxIntVal[0]) < (y - x) {
			maxIntVal[0] = x
			maxIntVal[1] = y

		}
	}
	return maxIntVal
}

type interval struct {
	start int
	end   int
}
type timeSlot []interval

func (p timeSlot) Len() int {
	return len(p)
}
func (p timeSlot) Less(x, y int) bool {
	if p[x].end < p[y].end {
		return true
	}
	return false
}
func (p timeSlot) Swap(x, y int) {
	p[x].start, p[y].start = p[y].start, p[x].start
	p[x].end, p[y].end = p[y].end, p[x].end
}
func main() {
	fmt.Println("vim-go")
	var calender = []interval{{-4, -1}, {0, 2}, {7, 9}, {3, 6}, {11, 12}, {14, 17}}
	fmt.Println(calender)
	sort.Sort(timeSlot(calender))
	fmt.Println("After sort:", calender)
	var nSlot = interval{1, 8}
	fmt.Println("Before Insertion Slot:", calender)
	newTmSlot := addNewTimeSlot(calender, nSlot)
	fmt.Printf("After New Slot:%t, Slot:%p", nSlot, newTmSlot)
	fmt.Println("After Insertion Slot:", newTmSlot)
	/*
		arr := []int{1, 3, 5, 7, 4, 6, 10, 11, 12, 13, 14, 15, 16, 20, 21, 23}
		maxIntVal := findLargestInterval(arr)
		fmt.Println("Max Intval:", maxIntVal)
	*/
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func addNewTimeSlot(oldSlot []interval, newSlot interval) []interval {
	//First add all time slot which end comes before given new add slot start time
	var newTmSlot []interval
	var unionSlot interval
	var i, l int

	l = len(oldSlot) - 1

	for i = 0; i < l; i++ {
		if oldSlot[i].end < newSlot.start {
			newTmSlot = append(newTmSlot, oldSlot[i])
		} else {
			break
		}
	}
	//At this stage , all slot which comes before given slots has added
	unionSlot.start, unionSlot.end = oldSlot[i].start, oldSlot[i].end
	for i < l && newSlot.end >= oldSlot[i].start {

		unionSlot.start = min(min(oldSlot[i].start, newSlot.start), unionSlot.start)
		unionSlot.end = max(max(oldSlot[i].end, newSlot.end), unionSlot.end)
		i++
	}
	newTmSlot = append(newTmSlot, unionSlot)

	for i <= l {
		newTmSlot = append(newTmSlot, oldSlot[i])
		i++
	}
	return newTmSlot
}

type llist struct {
	data int
	next *llist
}

func appendNode(tail, n, **llist) {
	llist * tmp = *n.next
	*tail.next = *n
	*n.next = nil
	*tail = *n
	*n = tmp
}
func mergeSortTwoSortedList(L1, L2 *llist) *list {
	dummy := new(llist)
	tail := dummy

	for l1 && l2 {
		if l1.data < l2.data {
			appendNode(&tail, &l1)
		} else {
			appendNode(&tail, &l2)
		}
	}
	if l1 {
		tail.next = l1
	} else if l2 {
		tail.next = l2
	}
}
func splitlinklist(l *llist) {
	var pre_slow, slow, fast *llist
	var slist *llist
	if l == nil || l.next == nil {
		return
	}
	pre_slow = nil
	slow == fast = nil
	slow == fast = l
	for fast != nil && fast.next != nil {
		pre_slow = slow
		slow = slow.next
		fast = fast.next.next
	}
	pre_slow.next = nil
	slist = slow
	mergeSortTwoSortedList(splitlinklist(l), splitlinklist(slist))
}

type tree struct {
	left  *tree
	right *tree
	key   int
}
type queueEntry struct {
	node  *tree
	lower int
	upper int
}

var queue []queueEntry

var intMax int = 1<<31 - 1
var intMin int = -(intMax - 1)

func istreeBSTUsingBFS(root *tree) bool {
	//first Insert Root node in Queue with MIn and Max value
	queue = append(queue, queueEntry{root, intMin, intMax})
	for len(queue) {
		n := queue[0]
		if n.data < n.lower || n.data < n.upper {
			return false
		}
		queue = append(queue, queueEntry{n.left, n.lower, n.data})
		queue = append(queue, queueEntry{n.right, n.lower, n.data})
		queue = queue[0:]
	}
	return true
}

//Find Next largest element of given element in Inoder traversal
func findNextInorderElemToGiven(t *tree, k int) {
	var seen_so_far *tree = nil
	for t {
		if t.data > k {
			seen_so_far = t
			t = t.left
		} else {
			t = t.right
		}

	}
	return seen_so_far
}

//Check is Given Tree is BST or not
//BST properties; each and every node left Subtree key should be equal or less then node key and its right subtree keys should be equal or greater then node key
// In sort, all node of BST should follow this properties .
//Another properties: Inorder Walk of binrary tree gives us ascending sort .

//First Algo : Walk each and every node and check its  max of its left subtree is less then or equal to node and min of right subtree is greater than or equal to node key
// This approch will make walk every node almost n times where n is number of Tree node therefore time complexity would be O(n2)

func getMin(t *tree) int {
	var min int
	for t {
		min = t.data
		t = t.left
	}
	return min
}
func getMax(t *tree) int {
	var max int
	for t {
		max = t.data
		t = t.right
	}
	return max
}

func isTreeBST(t *tree) bool {
	if t == nil {
		return true
	}
	lSubtree = t.left
	rSubtree = t.right
	max = getMax(lSubtree)
	min = getMin(rSubtree)
	if t.data < max || t.data > min {
		return false
	}
	return (isTreeBST(t.left) && isTreeBST(t.right))
}

//Another Approach: Instead calculting , min and max at every node. it better cache min and max value of its subtree and compare it with range

func isTreeBSTSecond(t *tree) bool {
	var maxInt = (1<<31 - 1)
	var minInt = -maxInt - 1
	return (isTreeBSTSecond(t, minInt, maxInt))
}
func isTreeBSTSecond(t *tree, min, max int) bool {
	if t == nil {
		return true
	}
	if t.data < min || t.data > max {
		return false
	}
	return isTreeBSTSecond(t.left, min, t.data) && isTreeBSTSecond(t.right, t.data, max)
}

//Find the Node which is equal to given Node and appear first in inorder traversal
func findNodeEqualInInorder(t *tree, key int) *tree {
	//Use BST logic : Go to Subtree which would contain given key
	if t == nil {
		return t
	}
	if t.data < key {
		return findNodeEqualInInorder(t.left, key)
	} else if t.data == key {
		return t
	} else {
		return findNodeEqualInInorder(t.right, key)
	}

}

//Find K largest Nodes in given BST tree
func findKLargestNode(t *tree, k *int) []*tree {
	var listKnode []*tree
	if t == nil {
		return nil
	}
	if t.right {
		findKLargestNode(t.right, k)
	}
	if k >= 0 {
		listkNode = append(listKnode, t)
		k--
	}
	if t.left {
		findKLargestNode(t.left, k)
	}
	return listKNode
}
func treeNode(k int) *tree {
	t := new(tree)
	t.data = k
	t.left = nil
	t.right = nil
}
func BuildBST() *tree {
	root := treeNode(19)
	insertBST(root, 7)
	insertBST(root, 3)
	insertBST(root, 2)
	insertBST(root, 5)
	insertBST(root, 11)
	insertBST(root, 17)
	insertBST(root, 13)
	insertBST(root, 43)
	insertBST(root, 23)
	insertBST(root, 37)
	insertBST(root, 29)
	insertBST(root, 31)
	insertBST(root, 41)
	insertBST(root, 47)
	insertBST(root, 53)
}
func insertBST(root *tree, key int) *tree {
	if root == nil {
		root = treeNode(key)
		return root
	}
	if key < root.data {
		root.left = insertBST(root.left, key)
	} else {
		root.right = insertBST(root.right, key)
	}
	return root
}

//Find median  number on running number of stream , if numbers read so far is even than even number is middle of average of 2 middle number
// while in odd case, it should be average of its .

//Solution , use 2 heap , one as minheap where all higher element store second is maxheap where all smaller element store,
//Make sure both heap should remain balance execpt size of data read so far is odd ;
//At odd number case, max heap should be  1 element more than max heap

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	//Push and Pop use Pointer receiver becoz they modify slice;s
	//not just contents
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	//Push and Pop use Pointer receiver becoz they modify slice;s
	//not just contents
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
