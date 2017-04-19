package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type wordFreq struct {
	word string
	freq int
}

type heap struct {
	capacity int
	count    int
	index    []wordFreq
}

func heapInit(int capacity) *heap {
	h := new(heap)
	h.capacity = capacity
	h.index = make([]int, capacity)
}

func heapify(h heap, index int) {
	//min Heap , smallest index is passed argument
	smallest = index
	left := 2*index + 1
	right := 2*index + 2
	if h.index[smallest] > h.index[left] {
		smallest = left
	}
	if h.index[smallest] > h.index[right] {
		smallest = right
	}
	if smallest != index {
		//Need to heapify
		swap(h, smallest, index)
		heapify(h, smallest)
	}

}
func buildHeap(h heap) {
	n := h.count - 1
	for i := 0; i < n/2; i++ {
		heapify(h, i)
	}
}
func heapPoll(h heap) string {
	if h.count < 1 {
		return nil
	}
	s := h.index[0].word
	swap(h, 0, h.count)
	h.count--
	heapify(h, 0)

}

func buildHashMapfromFile(words []string) []string {
	hashMap := make(map[string]wordFreq, 0)
	for _, w := range words {
		_, k := hashMap[w]
		if !k {
			hashMap[w] = allocNode(w)
		} else {
			hashMap[w].freq++
		}
	}
	heap := heapInit(10)
	for c, r := range hashMap {

		if heap.capacity > c {
			heap.index = append(heap.index, r)
			buildHeap(heap)
		} else {
			//Heap Reach its capacity,
			//Case 1: if read word frequence is  less than top of min heap , then skip this word else remove
			// remove top entry and add new entry
			if heap.index[0].freq < r.Freq {
				hash.index[0] = r
				heapify(heap, 0)
			}
		}
	}

}
