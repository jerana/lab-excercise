package main

import "fmt"

//import "sync"
//import "errors"

func findPathNode(root *tree, key int, path []*tree) (bool, []*tree) {
	var l, r bool
	//Base condition
	if root == nil {
		return false, path
	}
	path = append(path, root)
	if root.key == key {
		return true, path
	}
	//Check if Given Node exist
	l, path = findPathNode(root.left, key, path)
	if l != true {
		r, path = findPathNode(root.right, key, path)
	}
	if l != true && r != true {
		//Remove node from list which is added before call
		path = path[:len(path)-1]
		return false, path
	}
	return true, path
}

type tree struct {
	key   int
	left  *tree
	right *tree
}

func treeNodeAlloc(key int) *tree {
	n := new(tree)
	n.key = key
	n.left, n.right = nil, nil
	return n
}

func findLCABinaryTree(root *tree, n1, n2 int) *tree {
	var path1, path2 []*tree

	//Find path from root to n1, and root to n2 , if either of node doesn't
	//find path in tree, return false
	left, path1 := findPathNode(root, n1, path1)
	right, path2 := findPathNode(root, n2, path2)
	if !left || !right {
		return nil
	}
	//Walk to both vector and break when find first different nodes  during walk

	var i int

	for i = 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] != path2[i] {
			break
		}
	}
	return path1[i-1]
}

func FindLCAnode(root, n1, n2 *tree) *tree {
	if root == nil {
		return nil
	}
	//Case 1: if root node is equal to one of given nodes therefore root is LCA
	if root == n1 || root == n2 {
		return root
	} else {
		lsub := FindLCAnode(root.left, n1, n2)
		rsub := FindLCAnode(root.right, n1, n2)
		//Case 2 : if both nodes is non-null then given root is LCA
		if lsub != nil && rsub != nil {
			return root
		}
		//Case 3 : if right sub tree result is null while left not ; then check for LCA on left subtree
		if lsub != nil {
			return lsub
		} else {
			return rsub
		}
	}
}

/*
Logic require 2 steps
1.Steps first, From target nodes ,walk to its subtree which root as target and print k node,
2. Then walk it ancestor nodes and print  remaining k-d nodes on subtrees w
target nodes subtree doesn't exist , if acestor node is d disatance apart then search k-d nodes  on subtree which is root as target ancestor
*/

func printKnodeFromTargetBinary(root, target *tree, k int) int {
	if root == nil {
		return -1
	}
	if root.key == target.key {
		printKnodeFromRoot(root, k)
		return 0
	}
	ld := printKnodeFromTargetBinary(root.left, target, k)
	//Walk from root node to target node in both left and right subtree // and print node
	if ld != -1 {
		if ld+1 == k { //Means root node distance from target node {
			fmt.Println(root.key)
		} else {
			//Walk its right Subtree
			printKnodeFromRoot(root.right, k-ld-2)
		}
		return ld + 1
	}
	lr := printKnodeFromTargetBinary(root.right, target, k)
	if lr != -1 {
		if lr+1 == k { //Means root node distance from target node {

			fmt.Println(root.key)
		} else {
			//Walk its right Subtree
			printKnodeFromRoot(root.left, k-lr-2)

		}
		return lr + 1
	}
	return -1
}

func printKnodeFromRoot(root *tree, k int) {
	//Base case
	if root == nil || k < 0 {
		return
	}
	fmt.Println("Node: ", root.key, k)
	if k == 0 {
		fmt.Println(root.key)
		return
	}
	printKnodeFromRoot(root.left, k-1)
	printKnodeFromRoot(root.right, k-1)
	return
}

/*
Follow up from Basic Calculator II.

The expression string contains only non-negative integers, +, -, *, / ,(,)operators and empty spaces . The integer division should truncate toward zero.

You may assume that the given expression is always valid.

Some examples:

"(3+2)*2" = 10
" 3/(2+6) " = 0
" (3*(4 + 3)+5) / 2 " = 13
//Will Use Rverse polish notation,
//Use 2 stack DS , 1 for digist and other for operators

type stack struct {
	lock sync.Mutex //For thread safe
	s    []int
}

func (s stack) push(val int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s, val)
}
func (s stack) pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty stack")
	}
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
func (s stack) empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.s) == 0 {
		return true
	}
	return false
}
func (s stack) peek() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.s[0]
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
func calc(op1, op2 int, c byte) int {
	var rslt int
	if c == '+' {
		rslt = op1 + op2
	} else if c == '-' {
		rslt = op1 - op2
	} else if c == '/' {
		rslt = op1 / op2
	} else if c == '*' {
		rslt = op1 / op2
	}
	return rslt
}
func higherPriority(op1, op2 byte) bool {
	if op1 == '*' || op2 == '/' {
		return true
	} else if op1 == '+' || op2 == '-' {
		return true
	}
	return false
}

func simpleCalculator(t []byte) int {
	var st stack
	var op stack
	for i := 0; i < len(t); i++ {
		num := 0
		for isDigit(t[i]) { //it will either digit or operator
			num = num*10 + int(t[i])
			i++
		}
		ch := t[i]
		st.push(num)
		if ch == '(' {
			op.push(int(ch))
		} else if ch == ')' { //expr with ()
			for op.peek() != '(' {
				t1, _ := st.pop()
				t2, _ := st.pop()
				op, _ := op.pop()
				st.push(calc(t1, t2, op))
			}
			op.pop()
		} else {
			if !op.empty() && op.peek() != '(' && higherPriority(op.peek(), ch) {
				t1, _ := st.pop()
				t2, _ := st.pop()
				op, _ := op.pop()
				st.push(t1, t2, op)
			}
			op.push()
		}
	}
	for !op.Empty() {
		t1, _ := st.pop()
		t2, _ := st.pop()
		op, _ := op.pop()
		st.push(t1, t2, op)
	}
	return st.peek()
}
*/
