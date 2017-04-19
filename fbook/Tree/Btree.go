package main

import "fmt"
import "math/rand"

var Broot *tree

func main() {
	var head, prev *tree
	fmt.Println("vim-go")
	BuildBSTTree()
	printBtree(Broot)
	convertBSTtoDoubleList(Broot, &prev, &head)
	printConvBSTtoDoubleList(head)
}

type tree struct {
	key   int
	left  *tree
	right *tree
}

//Build Mirror Binaray tree
func buildBTreeMirror(root *tree) {
	if root == nil {
		return
	}
	swapNodePtr(root)
	buildBTreeMirror(root.left)
	buildBTreeMirror(root.right)
}
func swapNodePtr(root *tree) {
	tmpPtr := root.left
	root.left = root.right
	root.right = tmpPtr
}

//Non-recursive solution
func buildMirror(root *tree) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		return
	}
	var queue []*tree
	queue = append(queue, root)
	for len(queue) > 0 { //As long as queue in not emptry , take  node from queue and swap its pointer

		n := queue[0] // peek the node
		swapNodePtr(n)
		queue = queue[1:] //Pop the Node
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}

//HashMap using Value to Array Index mapping
// Random Value; Selected by Random Index selection in Array which  would be decided by Array size

var hashTbl = make(map[int]int) //Key is Value and Value is Index in Dynamic Array
var list []int                  //Variable Size Array

func Insert(val int) {
	_, ok := hashTbl[val]
	if !ok {
		list = append(list, val)
		hashTbl[val] = len(list) - 1
	}
}
func remove(val int) {
	_, ok := hashTbl[val]
	if !ok {
		lastIndex := len(list) - 1
		lastValue := list[lastIndex]
		curIndex := hashTbl[val]
		//curValue := list[curIndex]
		//		swap(list, curIndex, lastIndex)
		list[curIndex], list[lastIndex] = list[lastIndex], list[curIndex]
		list = list[0 : lastIndex-1]
		hashTbl[lastValue] = curIndex
		delete(hashTbl, val)
	}
}
func getRandom() int {
	if len(hashTbl) == 0 {
		return -1
	}
	randomIndx := rand.Intn(len(list))
	return list[randomIndx]
}
func DelRandom() int {
	if len(hashTbl) == 0 {
		return -1
	}
	randomIndx := rand.Intn(len(list))
	val := list[randomIndx]
	remove(val)
	return val
}

//Build Convert BST into Double-llink list
//Use Inorder Walk in BST and convert into Double link list
func convertBSTtoDoubleList(n *tree, prev, head **tree) {
	if n == nil {
		return
	}
	//Walk to left most node till find nil
	convertBSTtoDoubleList(n.left, prev, head)
	n.left = *prev
	if *prev != nil {
		(*prev).right = n
	} else {
		*head = n
	}
	*prev = n
	convertBSTtoDoubleList(n.right, prev, head)
}

//print Convert BST into LinkList
func printConvBSTtoDoubleList(head *tree) {
	fmt.Println("Converted Btree to Double List")
	for head.right != nil {
		fmt.Println(":->", head.key)
		head = head.right
	}
}
func printBtree(root *tree) {
	if root == nil {
		return
	}
	printBtree(root.left)
	fmt.Println(":<->", root.key)
	printBtree(root.right)
}
func bstInsert(root *tree, key int) *tree {
	if root == nil {
		root = new(tree)
		root.key = key
		return root
	}
	if root.key > key { //Walk to left subTree
		root.left = bstInsert(root.left, key)
	} else {
		root.right = bstInsert(root.right, key)
	}
	return root
}

func BuildBSTTree() {
	Broot = bstInsert(Broot, 6)
	Broot = bstInsert(Broot, 3)
	Broot = bstInsert(Broot, 2)
	Broot = bstInsert(Broot, 4)
	Broot = bstInsert(Broot, 8)
	Broot = bstInsert(Broot, 7)
	Broot = bstInsert(Broot, 9)
}
