package main

import (
	"fmt"
	"strconv"
)

var arr = []int{4, 2, 5, 5, 6, 1, 4}

//takebathandcome
var dics = map[string]bool{
	"take": true,
	"bat":  true,
	"bath": true,
	"and":  true,
	"come": true,
}

func main() {
	fmt.Println("vim-go")
	//	fmt.Println("Arrary:", arr)
	//fmt.Println("Number of words string:", isStrCanBreakIntoWords("takebathand", dics))
	buildAllListNode()
	printList(head)
	head = removeAllEntriesWhichhaveduplicate(head)
	printList(head)
	//	var root *tree = nil
	//	t := BuildBSTfromSortedArray(arr, 0, len(arr)-1)
	//	for _, e := range arr {
	//		root = bstInsert(root, e)
	//	}
	//PrintBSTInorder(t)
	//	PrintBSTInorder(root)
	//	fmt.Println(bstBFS(root))
}

type tree struct {
	key   int
	freq  int
	left  *tree
	right *tree
}

func NodeAlloc(key int) *tree {
	n := new(tree)
	n.key = key
	n.freq = 1
	return n
}

func BuildBSTfromSortedArray(arr []int, start, end int) *tree {
	var root *tree
	if start == end {
		root = NodeAlloc(arr[start])
		return root
	}
	if start < end {
		mid := start + (end-start)/2
		root = NodeAlloc(arr[mid])
		root.left = BuildBSTfromSortedArray(arr, start, mid-1)
		root.right = BuildBSTfromSortedArray(arr, mid+1, end)
	}
	return root
}

func PrintBSTInorder(t *tree) {
	if t == nil {
		return
	}
	if t.left != nil {
		PrintBSTInorder(t.left)
	}
	fmt.Println("%d:", t.key)
	if t.right != nil {
		PrintBSTInorder(t.right)
	}
	return

}
func bstInsert(root *tree, key int) *tree {
	if root == nil {
		root := NodeAlloc(key)
		return root
	}
	if root.key > key {
		root.left = bstInsert(root.left, key)
	} else if root.key < key {
		root.right = bstInsert(root.right, key)
	} else {
		root.freq++
	}
	return root
}

func bstBFS(root *tree) []string {
	var queue = []*tree{}
	var s []string

	//Append Root nodes
	queue = append(queue, root)

	for len(queue) > 0 {
		n := queue[0]
		if n == nil {
			s = append(s, ",")
		} else {
			tmp := strconv.Itoa(n.key) + ":" + strconv.Itoa(n.freq)
			s = append(s, tmp)
			queue = append(queue, n.left)
			queue = append(queue, n.right)
		}
		queue[0] = nil
		queue = queue[1:]

		fmt.Println("queue len", len(queue))
	}
	return s
}

//Given String and Dictonary, Find if String can be broken down into Valid Words
//Recursive Solution

func isStringCanBreakIntoWords(s string, dics map[string]bool) int {
	count := 0
	if len(s) == 0 {
		return 1
	}
	for i := 1; i < len(s); i++ {
		subStr := s[0:i]
		rmStr := s[i:len(s)]
		if dics[subStr] && isStringCanBreakIntoWords(rmStr, dics) > 0 {
			count += count + isStringCanBreakIntoWords(rmStr, dics)
		}
	}
	return count
}

//DP solution
//takebathand
func isStrCanBreakIntoWords(s string, dics map[string]bool) int {
	var i int
	fmt.Println("String len:", len(s))
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i = 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			var subStr string
			if i == len(s)-1 {
				subStr = s[j : i+1]
			} else {
				subStr = s[j:i]
			}
			_, ok := dics[subStr]
			fmt.Println("i", i, "j", j, "subStr", subStr)
			if ok && dp[j] > 0 {
				dp[i] += dp[j]
				fmt.Println("dp", i, j, dp[i], dp[j], "subStr", subStr)
			}
		}
	}
	return dp[len(s)-1]
}

type llink struct {
	val  int
	next *llink
}

func skipDuplicateNode(n *llink) *llink {
	for n != nil && n.next != nil {
		if n.val == n.next.val {
			n = n.next
			fmt.Println("next Node", n.val)
		} else {
			break
		}

	}
	fmt.Println("remove duplicates ", n.next)
	return n.next
}

var head *llink

func buildAllListNode() {
	buildLinklist(7, &head)
	buildLinklist(6, &head)
	buildLinklist(5, &head)
	buildLinklist(5, &head)
	buildLinklist(4, &head)
	buildLinklist(3, &head)
	buildLinklist(3, &head)
	buildLinklist(3, &head)
}

func buildLinklist(val int, head **llink) {
	n := new(llink)
	n.val = val
	if *head == nil {
		*head = n
	} else {
		n.next = *head
		*head = n
	}
	return
}
func printList(head *llink) {
	for head != nil {
		fmt.Println(":->", head.val)
		head = head.next
	}
}

func removeAllEntriesWhichhaveduplicate(head *llink) *llink {
	var prev *llink
	var cur *llink
	var newHead *llink
	prev = nil
	cur = head

	for cur != nil && cur.next != nil {
		if cur.val != cur.next.val {
			prev = cur
			cur = cur.next
		} else {
			if prev != nil {
				prev.next = skipDuplicateNode(cur)
				cur = prev.next
			} else {
				prev = skipDuplicateNode(cur)
				cur = prev
				newHead = prev
			}
		}

	}
	return newHead
}

//Recursive Method to delete all nodes which have duplicate
func rmDuplicateNodes(head *llink) *llink {
	if head == nil || head.next == nil {
		return head
	}
	var pNode *llink
	pNode = head.next
	var Val int
	Val = head.val
	if Val != pNode.Val {
		head.next = rmDuplicateNodes(pNode)
	} else {
		for pNode != nil && Val == pNode.Val {
			pNode = pNode.next
		}
		return rmDuplicateNodes(pNode)
	}
	return head
}

type listInterval struct {
	start int
	end   int
}

//Merge to non-Overlap list and final list should be non-overlap too
func mergeNonOverlapList(l1, l2 []listInterval) {
	var result listInterval
	for _, c := range l1 { //Walk to first List
		if c.start >= l2.end || c.end <= l2.start { // interval is not overlapped
			result = append(result, c)
		} else {
			//Merge overlap interval
			l2.start = min(l2.start, c.start)
			l2.end = max(l2.end, c.end)
		}
	}
	result = append(result, l2)
	return result
}

//
//Given a maze with cells being: gates, walls or empty spaces.

//INPUT maze:

//_ W G _
//_ _ _ W
//_ W _ W
//G W _ _

//RESULT should be:

//3 W G 1
//2 2 1 W
//1 W 2 W
//G W 3 4
//Fill the empty spaces with the number of steps to the closest gate.
//Allowed steps: UP, RIGHT, LEFT & DOWN

func findShortestToGate(matrix [][]int) {
	m := len(matrxi)
	n := len(matrix[0])
	list := make([]pointList, 1)
	for i := 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if matrix[i][j] == 'G' {
				updateShortestPathToGate(matrix, i, j, 0, list)
			}
		}
	}

}

type point struct {
	x int
	y int
}
type pointList []point

func (*pointList) Contain(index point) bool {
	for i = 0; i < len(pointList); i++ {
		if pointList[i].x == index.x || pointList.y == index.y {
			return true
		}
	}
	return false
}

func updateShortestPathToGate(matrix [][]int, x, y int, pDistance int, list pointList) {
	if x < 0 || x > len(matrix) || y < 0 || y > len(matrix[0]) {
		return
	}
	if matrix[x][y] == 'W' {
		return
	}
	if matrix[x][y] == '-' {
		matrix[x][y] = '0' + pDistance + 1
	} else {
		min = min((matrix[x][y] - '0'), pDistance+1)
		matrix[x][y] = '0' + min
	}
	var index point
	index.x = x
	index.y = y
	list = append(list, index)
	x1 := []int{1, -1, 0, 0}
	y1 := []int{0, 0, 1, -1}
	for k = 0; k < len(x1); k++ {
		newPoint = new(point)
		new.x = x + k
		new.y = y + k
		if !list.Contain(new) {
			updateShortestPathToGate(matrix, new.x, new.y, pDistance+1, point)
		}
	}
	list = list[0 : len(list)-1]
}

//We have discussed Backtracking and Knight’s tour problem in Set 1. Let us discuss Rat in a Maze as another example problem that can be solved using Backtracking.

/*A Maze is given as N*N binary matrix of blocks where source block is the upper left most block i.e., maze[0][0] and destination block is lower rightmost block i.e., maze[N-1][N-1]. A rat starts from source and has to reach destination. The rat can move only in two directions: forward and down.
In the maze matrix, 0 means the block is dead end and 1 means the block can be used in the path from source to destination. Note that this is a simple version of the typical Maze problem. For example, a more complex version can be that the rat can move in 4 directions and a more compl
*/
/* Algorithm :
If destination is reached
    print the solution matrix
Else
   a) Mark current cell in solution matrix as 1.
   b) Move forward in horizontal direction and recursively check if this
       move leads to a solution.
   c) If the move chosen in the above step doesn't lead to a solution
       then move down and check if  this move leads to a solution.
   d) If none of the above solutions work then unmark this cell as 0
       (BACKTRACK) and return false.
*/
func isSafe(matrix [][]int, x, y int) bool {
	if x > 0 && x <= len(matrix)-1 && y > 0 && y <= len(matrix[0])-1 && matrix[x][y] == 1 {
		return true
	}
	return false
}
func MarkMatrixWithBestPath(matrix [][]int, x, y int, sol [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if x == m-1 && y == n-1 {
		sol[x][y] = 1
		return true
	}
	if isSafe(matrix, x, y) {
		solo[x][y] = 1
		//Check if left side path is  possible
		if MarkMatrixWithBestPath(matrix, x+1, y, solo) == true {
			return true
		}
		//if left is not possible than check down side
		if MarkMatrixWithBestPath(matrix, x, y+1, solo) == true {
			return true
		}
		//if both direction failed to than make this point as blocked
		solo[x][y] = 0
	}
	return false

}
