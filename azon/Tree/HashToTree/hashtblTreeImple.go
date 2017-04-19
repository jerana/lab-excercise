package main

import "fmt"

func main() {
	setNode(20, 30) //key -20 , Val 30
	setNode(22, 23) //key -20 , Val 30
	setNode(21, 24) //key -20 , Val 30
	setNode(10, 10) //key -20 , Val 30
	setNode(8, 11)  //key -20 , Val 30
	setNode(17, 22) //key -20 , Val 30
	setNode(18, 23) //key -20 , Val 30
	setNode(19, 33) //key -20 , Val 30
	setNode(28, 54) //key -20 , Val 30
	setNode(29, 55) //key -20 , Val 30
	setNode(25, 58) //key -20 , Val 30
	printTree(root)
	//	fmt.Println("Found node:", GetNode(17))
	//	fmt.Println("Del node key :")
	//	var k int
	//	fmt.Scanf("%d", &k)
	//	delNode(root, k)
	//	fmt.Println("After del key: ", k)
	//	printTree(root)
	nums := []int{2, 5, 3, 4, 6, 1}
	list := countGreater(nums)
	fmt.Println("list:", list)

}

//Build Hash Table functinality using BST tree
//Write Get,Set ,lookup,del function Using BST

//Data Structure
type TreeHashNode struct {
	key   int
	Val   int
	left  *TreeHashNode
	right *TreeHashNode
}

var root *TreeHashNode

func printTree(cur *TreeHashNode, left bool) {
	if cur == nil {
		return
	}
	fmt.Println(cur.key)
	printTree(cur.left)
	printTree(cur.right)
	return
}

func findParentNode(root *TreeHashNode, key int) *TreeHashNode {
	//Following function try to return previous node
	if root == nil {
		return nil
	}
	cur := root
	var prev *TreeHashNode

	for cur != nil {
		if cur.key == key {
			return prev
		} else if key < cur.key {
			prev = cur
			cur = cur.left
		} else {
			prev = cur
			cur = cur.right
		}
	}
	return prev
}
func GetNode(key int) int {
	if root == nil {
		return -1
	}
	parent := findParentNode(root, key)
	if parent == nil {
		//Root would be match node
		return root.Val
	}
	//Case Node not found ,
	if parent.left != nil && parent.left.key == key {
		return parent.left.Val
	}
	if parent.right != nil && parent.right.key == key {
		return parent.right.Val
	}
	return -1
}
func setNode(key, val int) {
	if root == nil {
		root = new(TreeHashNode)
		root.Val = val
		root.key = key
		return
	}
	parent := findParentNode(root, key)
	//Three Two cases, Node Already exist then just update node
	//else add node into its parent Subtree
	tnodeL := parent.left
	if tnodeL != nil && tnodeL.key == key {
		tnodeL.Val = val
		return
	}
	tnodeR := parent.right
	if tnodeR != nil && tnodeR.key == key {
		tnodeR.Val = val
		return
	}
	if key < parent.key {
		n := new(TreeHashNode)
		n.key = key
		n.Val = val
		parent.left = n
	} else {
		n := new(TreeHashNode)
		n.key = key
		n.Val = val
		parent.right = n
	}
	return
}
func lookup(key int) *TreeHashNode {
	p := findParentNode(root, key)
	if p.left != nil && p.left.key == key {
		return p.left
	}
	if p.right != nil && p.right.key == key {
		return p.right
	}
	return nil
}
func getNodeFromParent(node *TreeHashNode, key int) *TreeHashNode {
	n := node.left
	if n != nil && n.key == key {
		return n
	}
	n = node.right
	if n != nil && n.key == key {
		return n
	}
	return nil
}
func delNode(root *TreeHashNode, key int) bool {
	p := findParentNode(root, key)
	delN := getNodeFromParent(p, key)
	if delN != nil {
		//Case 1 : No Subtree of request node
		if delN.left == nil && delN.right == nil {
			if p.left.key == delN.key {
				p.left = nil
			} else {
				p.right = nil
			}
			return true
		}
		//Case 2: 1 Subtree exist
		if delN.left == nil && delN.right != nil {
			if p.left.key == delN.key {
				p.left = delN.right
			} else {
				p.right = delN.right
			}
			return true
		}
		//Case 3: 1 Subtree exist
		if delN.left != nil && delN.right == nil {
			if p.left.key == delN.key {
				p.left = delN.left
			} else {
				p.right = delN.left
			}
			return true
		}
		//Case 4: 2 Subtree exist
		//first Move 1 node to right and then walk left most node
		// And swap key,value of left most node with current node,
		//Remove left most node
		n := delN.right
		//Find left most node
		for n.left != nil {
			n = n.left
		}
		//copy key,val with left most  with current node
		delN.key = n.key
		delN.Val = n.Val

		//Delete left most node
		delNode(n, key)
		return true
	}
	return false
}
