package main

import "fmt"

/*
Given list of nodes of a tree, find the root of the tree. Nodes in the list are not in any particular order. If some of the nodes in the tree are not given, return null
Eg:

        A
	   B        C
	   E         F
	   Input : [B, C, A, E, F] Output : A
*/

//Logic : Only root nodes will be reference by any nodes left or right pointer
//therefore, make a copy of node list and walk of first list and remove its left or righ//nodes from second list if its is referenced
//Option 2: put all nodes into map and remove it during first list walk
type tree struct {
	key         int
	left, right *tree
}

var treeNodeList []*tree

func returnRootNodesfromlist(tlist []*tree) *tree {
	//Built Map ;
	copyLMap := make(map[*tree]bool, 0)
	//Call copy in-built function
	for _, n := range tlist {
		copyLMap[n] = true
	}
	//till hear , built-up Map
	//Now Walk to list and remove node right and left node from MAp
	for _, n := range tlist {
		if n.left != nil {
			delete(copyLMap, n.left)
		}
		if n.right != nil {
			delete(copyLMap, n.right)
		}
	}
	if len(copyLMap) > 0 {
		for n, _ := range copyLMap {
			return n
		}
	}
	return nil
}

/*
Given an array A of size N with numbers from 1 to N in random order, return an array B such that B[i] is equal to the number of elements greater than A[i] to its right in A.

This question is same as this Leetcode quesiton.

Except the fact that here we know that array is permutation of [1..N]. Can we leverage that fact and optimize this? Say to time O(n)?


*/
var Nums = []int{3, 2, 5, 6, 1, 4}

//Use binary Search
func countGreater(nums []int) []int {
	//Result Count list
	var rslt []int
	var sortList []int
	fmt.Println("list:", Nums)

	for i := len(nums) - 1; i >= 0; i-- {
		index := findIndex(sortList, nums[i])
		fmt.Println("Get Index:", index)
		rslt = append(rslt, index)
		var tmpList1 []int
		if index != 0 {
			index -= len(sortList)
		}
		if index == 0 {
			tmpList1 = make([]int, 0)
		} else {
			tmpList1 = sortList[:index]
		}
		tmpList2 := sortList[index:]
		tmpList1 = append(tmpList1, nums[i])
		fmt.Println("index", index, "tmpList1:", tmpList1, "tmpList2:", tmpList2)
		sortList = append(tmpList1, tmpList2...)
		fmt.Println("sortList:", sortList)
		//sorted.add(index, nums[i])
	}
	return rslt

}

func findIndex(sortList []int, target int) int {
	if len(sortList) == 0 {
		return 0
	}
	start := 0
	end := len(sortList) - 1
	if target < sortList[start] {
		return end + 1
	}
	if target > sortList[end] {
		return 0
	}
	for start+1 > end {
		mid := start + (end-start)/2
		if target > sortList[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if target >= sortList[start] {
		return start

	}
	return end
}
