package main

import (
	"fmt"
	"math/rand"
)

//var arr = []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}

func main() {
	var arr []int = []int{10, 2, 1, 0, 12, 31, 3, 13, 11, 4, 5}
	l := len(arr)
	fmt.Println("Array elem :", arr)
	fmt.Println("how many largest element u want")
	var k int
	fmt.Scanf("%d", &k)
	fmt.Printf("Largest %d elem in arr\n", k)
	pivot := rand.Intn(l - 1)
	fmt.Println("Pivot ", pivot)
	newIndx := partitionAroundPivotIndex((arr), 0, l-1, pivot)
	fmt.Println("newIndx ", newIndx)
	fmt.Println("Array elem after partition:", arr)

	/*
			var num int
			var pair [2]int
			fmt.Printf("Enter Number which need to search\n")
			fmt.Scanf("%d", &num)
			fmt.Printf("Search Num is :%d \n", num)
			pair[0] = findFirstOccuranceIndex(arr, num, true)
			pair[1] = findFirstOccuranceIndex(arr, num, false)
			fmt.Println("Pair of num:  is : ", num, pair)
				firstO := findFirstOccuranceofElemK(arr, num)
				if firstO != -1 {
					fmt.Printf("First occurance :%d is %d \n", num, firstO)
				} else {
					fmt.Printf("Num:%d not found\n", num)
				}
		var src string
		var target string
		fmt.Printf("Enter Src string:\n")
		fmt.Scanln(&src)

		fmt.Printf("Enter target string:\n")
		fmt.Scanln(&target)
		fmt.Println(src, target)
		fmt.Printf("=====Let's Find Target:%s string into Src:%s\n ", target, src)
		fmt.Printf("Does Target:%s exist in Src:%s :%d \n", target, src, strstr(src, target))
	*/
}

func findFirstOccuranceofElemK(arr []int, k int) int {
	//Use Binary Search Algo
	var mid int
	var start, end int
	start = 0
	end = len(arr) - 1
	var result int = -1
	for start <= end {
		mid = start + (end-start)/2
		fmt.Println(arr[mid])
		if arr[mid] <= k {
			start = mid + 1
		} else {
			fmt.Println(result)
			result = mid
			end = mid - 1
		}
	}
	return result
}

//Function to  find target string exist in given string or not
//Using O(MN) algo
//walk each element and try to find target string from that offset

func strstr(s, t string) int {
	var i, j int
	for i = 0; i < len(s); i++ {
		for j = 0; j < len(t); j++ {
			//Case 1: if j reach its len(t) then target string located at index i
			if j == len(t)-1 {
				return i
			}
			//Case 2; Walk index shouldn't stop if Walk index+ target string len is greater than given string
			if i+len(t) > len(s) {
				return -1
			}
			if t[j] != s[i+j] {
				break
			}
		}
	}
	return -1
}

func findFirstOccuranceIndex(arr []int, k int, first bool) int {
	var start, end int
	var rslt = -1
	start = 0
	end = len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		fmt.Println("Mid: s: e: ", arr[mid], start, end)
		if arr[mid] == k {
			rslt = mid
			if first {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if arr[mid] < k {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return rslt
}

//For Sorted Array, Find the index i if it's value match with index value
//otherwise return -1
func findIndexEqualToVal(arr []int) int {
	var start, end, mid int
	var rslt int
	rslt = -1
	start = 0
	end = len(arr) - 1
	for start < end {
		mid = start + (end-start)/2
		if arr[mid] < mid {
			start = mid + 1
		} else if arr[mid] == mid {
			return mid
		} else {

		}
	}
	return rslt
}

//Find minimum elememt in cyclic sorted array
//Alog: Use Binary Search Algorithm , Try to find index where its neighbors elem//element are bigger than given element
// if A[M] > A[n-1] . then Smallest element must in range of (m+1,n-1)
//Conversally if A[m] < A[n-1] , then smallest element can't be betweea
//n (m+1,n-1) though m itself could be smallest element

func findLowestInCyclicSortedArray(arr []int) int {
	var start, end int
	var mid int
	start = 0
	for start < end {
		mid = start + (end-start)/2
		if arr[mid] > arr[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	//loop will end when start == end
	return start
}

//Problem : Find Element K in cyclic Sorted Array
func findElementKInCyclicSortedArray(arr []int, k int) int {

	start := 0
	end := len(arr) - 1
	mid := 0
	for start < end {
		mid = start + (end-start)/2
		if arr[mid] == k {
			return mid
		} else if arr[mid] < arr[end] { //means need to search higher value in array
			if k > arr[mid] && k < arr[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else { //higher value at mid , therefore search for lower value in array
			if k > arr[start] && k < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}

		}
	}
	return -1
}

//Find Largest K element in Array with O(n)
//Using MinHeap , gives nlog(k) time complexity
//Use devide and partition algo, we can achevie this in O(n)
//Select pivot eleement randomly, and partition array element like. all element its left is greater than pivot
// all right element to pivot is less than pivot
// Check if  after partition new pivot index is k

func partitionAroundPivotIndex(arr []int, start, end, pivot int) int {
	pVal := arr[pivot]
	newPivotIndex := start
	swap(&arr[pivot], &arr[end])
	for i := start; i < end; i++ {
		if arr[i] > pVal { //Walking element is greater than pivot
			swap(&arr[i], &arr[newPivotIndex])
			newPivotIndex++
		}

	}
	swap(&arr[end], &arr[newPivotIndex])
	return newPivotIndex
}
func swap(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func FindNearestRepetition(paraGraph []string) int {
	var nearestRepetitionWord int = 1<<31 - 1
	//Map key is word which appear in paragraph and val is last seen index
	// in paraGraph
	paraGraphMap := make(map[string]int)

	for i, s := range paraGraph {
		lastSeen, ok := paraGraphMap[s]
		diff := i - lastSeen
		paraGraphMap[s] = i
		paraGraphMap = min(paraGraphMap, diff)
	}
	return paraGraphMap
}

//Find Array Indices from target
//If array is sorted
func SumTwo(arr []int, target int) []int {
	var i, j int
	var rslt []int
	i = 0
	j = len(arr) - 1
	for i < j {
		sum := arr[i] + arr[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			rslt = append(rslt, i+1, j+1)
		}
	}
}

func filpBits(bits int) int {
	if bits == 0x0 || bits == 0x11 {
		return bits
	}
	bits = ^bits & 0x11
	return bits
}

func flipTwoAdjacentbits(n int) int {
	var finalBits = 0x0
	var ret = 0x0
	for i := 0; i < 31; i = i + 2 {
		bits = n & 0x11
		n = n >> 2
		ret = filpBits(bits)
		finalBits |= ret << i
	}
}

type freq struct {
	c byte
	f int
}

func heapify(freq []int, i int, heap_sz int) {
	l := 2*i + 1 //left child
	r := 2*i + 2 //right child
	largest = i
	if freq[l].f > freq[i].f && l < heap_sz {
		largest = l
	}
	if freq[r].f > freq[largest].f && r < heap_sz {
		largest = r
	}
	if largest != i {
		swap(&freq[largest], &freq[i])
		heapify(freq, largest, heap_sz)
	}
	return
}
func buildHeap(freq []int, heap_sz int) {
	i := (heap_sz - 1) / 2
	for i > 0 {
		heapify(freq, i, heap_sz)
		i--
	}
}
func extractMax(freq []int, heap_sz int) freq {
	e := freq[0]
	if heap_sz > 1 {
		freq[0] = freq[heap_sz-1]
		heapify(freq, 0, heap_sz-1)
	}
	return e
}
func placeCharcterAtgiven(arr []byte, k int) []byte {
	var freqNode [256]freq
	var newEntry int
	for idx, c := range arr {
		if freqNode[c] == 0 {
			freqNode[c].c = c, 1
			newEntry++
		} else {
			(freqNode[c].f)++
		}
		arr[idx] = '0'
	}
	buildHeap(freqNode, len(arr))
	for i := 0; i < newEntry; i++ {
		ch := extractMax(freqNode, 256-i)
		cnt = ch.f
		c := ch.c
		var p int = i
		for arr[p] != '0' {
			p++
		}
		for j := 0; j < cnt; j++ {
			if p+d*j >= len(arr) {
				return 0
			}
			arr[p+d*j] = c
		}

	}
}

//Merge given Overlapp set

type set struct {
	start int
	end   int
}
type setList []set

func (s *setList) Less(x, y int) {
	if x > y {
		return x
	}
	return y
}
func (s *setList) Len() int {
	return len(s)
}
func (s *setList) Swap(x, y int) {
	s.x, s.y = s.y, s.x
}

func mergeOverlapSet([]set) []set {
	sort.Sort(setList(set))
	index = 0
	for i := 0; i < len(set); i++ {
		if index != 0 && set[index].start <= set[i].end {
			set[index-1].end = max(set[index-1].end, set[i].end)
			set[index-1].start = max(set[index-1].start, set[i].start)
			index--
		} else {
			set[index] = set[i]
		}
		index++
	}
}

//Question: Given a sequence of positive integers A and an integer T, return whether there is a continuous sequence of A that sums up to exactly T
//Example
//[23, 5, 4, 7, 2, 11], 20. Return True because 7 + 2 + 11 = 20
//[1, 3, 5, 23, 2], 8. Return True because 3 + 5 = 8
//[1, 3, 5, 23, 2], 7 Return False because no sequence in this array adds up to 7

//Will Use sliding Wondow logic. Make left and rigth pointer and keep increment the right till sum is equal to target num or SUM is greater than target num. On SUM getting greater than target num
// Increment left pointer till SUM less than or equal to target SUM

func isSequenceSUMMatchwithTargetNUM(target int, arr []int) bool {
	var start = 0
	var i = 0
	var sum = 0
	for i < len(arr) {
		if sum < target {
			sum += arr[i]
		} else if sum == target {
			return true
		} else {
			sum += arr[i]

			for sum < target {
				sum -= arr[start]
				start++
			}
			if sum == target {
				return true
			}
		}
		i++
	}
	return false
}
func buildBalanceBST(arr []int, start, end int) *Tree {
	if start >= end {
		return nil
	}
	t := newNode(Tree)
	mid = start + (end-start)/2
	t.Key = mid
	t.left = buildBalanceBST(arr, start, mid-1)
	t.right = buildBalanceBST(arr, mid+1, end)
	return t

}

/*We have a system that records scores. We care about how many times we see the same score, and we want to maintain a rough ordering. We also want to send this information over the wire so that it can be collated with other results. As such we have decided to represent the stream of scores, and the count of how many times we see the same score, as an unbalanced binary search tree.

Your job is to write a method that will take a stream of integer scores, and put them into a tree while counting the number of times each score is seen. The first score from the given list should occupy the root node. Then you need to traverse the tree breadth-first to generate a string representation of the tree. Scores are to be inserted into the tree in the order that they are given.

For example, if you were given the stream of scores: [4, 2, 5, 5, 6, 1, 4].
*/
type TreeCount struct {
	cnt         int
	key         int
	left, right TreeCount
}

func BuildBSTTree(arr []int) *TreeCount {

	root := new(TreeCount)
	root.key = arr[0]
	for i := 1; i < len(arr); i++ {
		SearchAndInsert(root, arr[i])
	}
	return root
}
func SearchAndInsert(n *Tree, key int) *Tree {

	if n == nil {
		n = new(Tree)
		n.key = key
		return n

	} else if n.key == key {
		n.cnt++
		return n
	}
	if key > n.key {
		n.right = SearchAndInsert(n.right, key)
	} else {
		n.left = SearchAndInsert(n.left, key)
	}
	return n

}
