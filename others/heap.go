package main

import (
	"container/heap"
	"fmt"
)

//Find median  number on running number of stream , if numbers read so far is even than even number is middle of average of 2 middle number
// while in odd case, it should be average of its .

//Solution , use 2 heap , one as minheap where all higher element store second is maxheap where all smaller element store,
//Make sure both heap should remain balance execpt size of data read so far is odd ;
//At odd number case, max heap should be  1 element more than max heap

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap)Top() interface{} {
	return h[0]
}

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
func (h maxHeap)Top() interface{} {
	return h[0]
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type hi minHeap //Min Heap
type lo maxHeap // Max heap

func addNum(num int) {
	lo.push(num) //Add to Max heap
	hi.push(lo.Pop()) //Balancing step
	if lo.Len() < hi.Len() { //maintain size propety 
		lo.Push(hi.pop())
	}
}
func findMedian() float32 {
	if lo.Len() > hi.Len() {
		return float32(lo.Top() 
	} 
	return(lo.Top() + hi.Top() )/2
}

func main() {
	h := &minHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("Minimum :%d\n", heap.Pop(*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d", heap.Pop(h))
	}
}
