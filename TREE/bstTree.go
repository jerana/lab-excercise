package main

import "fmt"
import "log"

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

type maxSumNode struct {
	val  int
	node *tree
}

func maxSumSubtree(root *tree, max *maxSumNode) int {
	if root == nil {
		return 0
	}
	lsum := maxSumSubtree(root.left, max)
	rsum := maxSumSubtree(root.right, max)
	csum := root.key + lsum + rsum
	if csum > max.val {
		max.node = root
		max.val = csum
	}
	return csum
}

func istreeBSTUsingBFS(root *tree) bool {
	//first Insert Root node in Queue with MIn and Max value
	queue = append(queue, queueEntry{root, intMin, intMax})
	for len(queue) > 0 {
		n := queue[0]
		if n.node.key < n.lower || n.node.key < n.upper {
			return false
		}
		queue = append(queue, queueEntry{n.node.left, n.lower, n.node.key})
		queue = append(queue, queueEntry{n.node.right, n.lower, n.node.key})
		queue = queue[0:]
	}
	return true
}

//Find Next largest element of given element in Inoder traversal
func findNextInorderElemToGiven(t *tree, k int) *tree {
	var seen_so_far *tree = nil
	for t != nil {
		if t.key > k {
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
	if t == nil {
		return -1
	}
	if t.left != nil {
		return getMin(t.left)
	}
	return t.key
}
func getMax(t *tree) int {
	if t == nil {
		return -1
	}
	if t.right != nil {
		return getMax(t.right)
	}
	return t.key
}

func isTreeBST(t *tree) bool {
	var max, min int
	if t == nil {
		return true
	}
	lSubtree := t.left
	rSubtree := t.right
	if lSubtree != nil {
		max = getMax(lSubtree)
		if t.key < max {
			return false
		}
	}
	if rSubtree != nil {
		min = getMin(rSubtree)
		if t.key > min {
			return false
		}
	}
	return (isTreeBST(t.left) && isTreeBST(t.right))
}

//Another Approach: Instead calculting , min and max at every node. it better cache min and max value of its subtree and compare it with range

func isTreeBST2(t *tree) bool {
	var maxInt = (1<<31 - 1)
	var minInt = -maxInt - 1
	return (isTreeBSTSecond(t, minInt, maxInt))
}
func isTreeBSTSecond(t *tree, min, max int) bool {
	if t == nil {
		return true
	}
	if t.key < min || t.key > max {
		return false
	}
	return isTreeBSTSecond(t.left, min, t.key) && isTreeBSTSecond(t.right, t.key, max)
}

//Find the Node which is equal to given Node and appear first in inorder traversal
func findNodeEqualInInorder(t *tree, key int) *tree {
	//Use BST logic : Go to Subtree which would contain given key
	if t == nil {
		return t
	}
	if t.key < key {
		return findNodeEqualInInorder(t.left, key)
	} else if t.key == key {
		return t
	} else {
		return findNodeEqualInInorder(t.right, key)
	}

}

//Find K largest Nodes in given BST tree
func findKLargestNode(t *tree, k int) []*tree {
	var listKnode []*tree
	findKLargestHelper(t, &k, listKnode)
	fmt.Println(listKnode, len(listKnode))
	return listKnode
}

func findKLargestInBST(t *tree, k int) []*tree {
	findKLargestInBSTHelper(t, k)
	fmt.Println(list, len(list))
	return list
}

var list []*tree

func findKLargestInBSTHelper(t *tree, k int) {

	if t != nil && len(list) < k {
		if t.right != nil {
			findKLargestInBSTHelper(t.right, k)
		}
		//Add Node into list
		if len(list) < k {
			list = append(list, t)
			fmt.Println("Node add:", t.key)
		}
		if t.left != nil {
			findKLargestInBSTHelper(t.left, k)
		}
	}
}

func findKLargestHelper(t *tree, k *int, l []*tree) {
	if t == nil {
		return
	}
	if t.right != nil {
		findKLargestHelper(t.right, k, l)
	}
	if *k >= 0 {
		l = append(l, t)
		*k--
	}
	if t.left != nil {
		findKLargestHelper(t.left, k, l)
	}
	return
}
func treeNode(k int) *tree {
	t := new(tree)
	t.key = k
	t.left = nil
	t.right = nil
	return t
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
	return root
}
func insertBST(root *tree, key int) *tree {
	if root == nil {
		root = treeNode(key)
		return root
	}
	if key < root.key {
		root.left = insertBST(root.left, key)
	} else {
		root.right = insertBST(root.right, key)
	}
	return root
}
func inorderTraverse(root *tree) {
	if root == nil {
		return
	}
	if root.left != nil {
		inorderTraverse(root.left)
	}
	fmt.Printf(":->%d \t", root.key)
	if root.right != nil {
		inorderTraverse(root.right)
	}
}
func main() {
	fmt.Println("BUild BST ")
	var k int
	root := BuildBST()
	inorderTraverse(root)
	fmt.Printf("How many nodes want to find\n")
	fmt.Scanf("%d", &k)
	fmt.Printf("Find %d largest nodes in BST\n", k)
	list := findKLargestNode(root, k)
	list1 := findKLargestInBST(root, k)
	for _, n := range list {
		fmt.Printf("%d nodes in tree\n", k, n.key)
	}
	fmt.Printf("Tree %d largest Node:", k)
	for _, n := range list1 {
		fmt.Println(":->", n.key)
	}
	var n1, n2 int
	fmt.Println("Check if Tree is BST", isTreeBST(root), isTreeBST2(root))
	n1 = 31
	n2 = 53
	for {
		fmt.Println("Enter the Nodes which LCA needs")
		if _, err := fmt.Scan(&n1, &n2); err != nil {
			log.Print("  Scan for n1,n2 failed, due to ", err)
			return
		}
		n := findLCAinBST(root, n1, n2)
		if n != nil {
			fmt.Println("LCA node ", n.key, "for nodes", n1, n2)
		} else {
			fmt.Println("No LCA node", n1, n2, "nodes")
		}
	}
	sortedArray := []int{10, 20, 40, 45, 50, 100, 120, 200, 300, 400, 430, 450, 500}
	tree := buildMinimumHeightBSTFromSortedArray(sortedArray)
	fmt.Println("Tree", tree, " from sortedArray", sortedArray)
	printBST(tree)

}
func printBST(root *tree) {
	if root == nil {
		fmt.Println("")
	}
	if root.left != nil {
		printBST(root.left)
	}
	fmt.Println(root.key)
	if root.right != nil {
		printBST(root.right)
	}
	return
}

func findLCAinBST(root *tree, n1, n2 int) *tree {
	for root != nil {
		if (root.key > n1 && root.key < n2) ||
			(root.key > n2 && root.key < n1) {
			return root
		}
		if root.key < n1 && root.key < n2 {
			root = root.right
		}
		if root.key > n1 && root.key > n2 {
			root = root.left
		} else {
			return nil
		}
	}
	return nil
}

//Build BST from preorder Sequence
func buildBSTfromPreorder(list []int) *tree {
	var root_idx = 0
	var upper = (1 << 31) - 1
	var lower = -upper - 1
	return buildBSTfromPreorderHandler(list, lower, upper, &root_idx)
}
func buildBSTfromPreorderHandler(list []int, lower, upper int, root_idx *int) *tree {
	var rootIdx = *root_idx
	if rootIdx > len(list) {
		return nil
	}
	if list[rootIdx] < lower || list[rootIdx] > upper {
		return nil
	}
	*root_idx++
	lsubTree := buildBSTfromPreorderHandler(list, lower, list[rootIdx], root_idx)
	rsubTree := buildBSTfromPreorderHandler(list, lower, list[rootIdx], root_idx)
	n := new(tree)
	n.key = list[rootIdx]
	n.left = lsubTree
	n.right = rsubTree
	return n

}

//Hash Tbl build up from log file
var logHashTbl = make(map[int]*tree)
var root *tree

func readFileAndUpdateHashTree(ids int) {
	t, ok := logHashTbl[ids]
	if ok {
		//ids Exist , Need to update BST tree Node for its freq
		//First Delete Tree Node from BST
		bstDelete(root, t)
		//Update t node
		t.key++
		bstAdd(root, t)
	} else {
		t = new(tree)
		t.key++
		bstAdd(root, t)
		logHashTbl[ids] = t
	}
}

//There are multiple case on BST delete
//1. Delete node is leaf node , Don't need to do much , just free node and update// its' parent points to this node
//2. Delete node has 1 Subtree .
func bstDelete(root, t *tree) *tree {
	if root == nil {
		return nil
	}
	if root.key < tree.key {
		root.right = bstDelete(root.right, tree)
	} else if root.key > tree.key {
		root.right = bstDelete(root.right, tree)
	} else {
		//Key matched
		//If Found nodes just have 1 sub tree
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			//Both nodes exist
			tmp := findMinNode(root.right)
			root.key = tmp.key
			root.right = bstDelete(root.right, tmp.key)
		}

	}
	return root
}
func bstAdd(root *tree, t *tree) *tree {
	if root == nil {
		root := new(tree)
		root.key = key
		return root
	}
	if root.key > key {
		root.left = bstAdd(root.left, key)
	} else if root.key <= key {
		root.right = bstAdd(root.right, key)
	}
	return root
}

func buildMinimumHeightBSTFromSortedArray(slist []int) *tree {
	if len(slist) == 0 {
		return nil
	}
	return buildMinimumHeightBSTFromSortedArrayHelper(slist, 0, len(slist)-1)
}
func buildMinimumHeightBSTFromSortedArrayHelper(slist []int, start, end int) *tree {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	n := new(tree)
	n.key = mid
	n.left = buildMinimumHeightBSTFromSortedArrayHelper(slist, start, mid-1)
	n.right = buildMinimumHeightBSTFromSortedArrayHelper(slist, mid+1, end)
	return n
}

//Find Min depth of Tree

//Using BFS logic,
//track rightMost node and increase depth when see rightmost node
// When found first leaf node, stop search and return depth

func minDepth(t *tree) {
	if t == nil {
		return 0
	}
	var queue []*tree
	var rightMost *tree
	queue = queue(t)
	rightMost = t
	var depth = 1
	for len(queue) > 0 { //Till queue is not empty
		n := queue[0]
		if n.left == nil && n.right == nil {
			break
		}
		if n.left != nil {
			queue[len(queue)] = n.left
		}
		if n.right != nil {
			queue[len(queue)] = n.right
		}
		if rigthMost == n {
			depth++
			if n.right {
				rightMost = n.right
			} else {
				rightMost = n.left

			}
		}

	}
	return depth

}

//IsBalance BTree
//o(n2) worest case
func isBalanced(root *tree) bool {
	if root == nil {
		return true
	}
	result := (abs(maxDepth(root.left)-maxDepth(root.right)) <= 1) &&
		isBalanced(root.left) && sBalanced(root.right)
}

//O(n) bottom up recursion
func isBalanced(root *tree) bool {
	return maxDepth(root) != -1
}
func maxDepth(root *tree) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.left)
	if l == -1 {
		return -1
	}
	r := maxDepth(root.right)
	if r == -1 {
		return -1
	}

	if abs(l-r) >= 1 {
		return -1
	}
	return (max(l, r) + 1)
}
