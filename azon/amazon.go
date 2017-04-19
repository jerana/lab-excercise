package main

import "fmt"
import "strconv"
import "sort"

func main() {
	var arr = []int{4, 2, 5, 5, 6, 1, 4}
	root := BuildBSTTree(arr)
	fmt.Println("String:", BFS(root))
}

/*We have a system that records scores. We care about how many times we see the same score, and we want to maintain a rough ordering. We also want to send this information over the wire so that it can be collated with other results. As such we have decided to represent the stream of scores, and the count of how many times we see the same score, as an unbalanced binary search tree.

Your job is to write a method that will take a stream of integer scores, and put them into a tree while counting the number of times each score is seen. The first score from the given list should occupy the root node. Then you need to traverse the tree breadth-first to generate a string representation of the tree. Scores are to be inserted into the tree in the order that they are given.

For example, if you were given the stream of scores: [4, 2, 5, 5, 6, 1, 4].
*/
type TreeCount struct {
	cnt   int
	key   int
	left  *TreeCount
	right *TreeCount
}

func BuildBSTTree(arr []int) *TreeCount {

	root := new(TreeCount)
	root.key = arr[0]
	root.cnt++
	for i := 1; i < len(arr); i++ {
		SearchAndInsert(root, arr[i])
	}
	return root
}
func SearchAndInsert(n *TreeCount, key int) *TreeCount {

	if n == nil {
		n = new(TreeCount)
		n.key = key
		n.cnt++
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
func BFS(root *TreeCount) []string {
	var queue []*TreeCount
	//Insert first root element
	queue = append(queue, root)
	var str []string
	for len(queue) > 0 { //Till queue is not Empty
		//Dequeue node
		n := queue[0]
		queue = queue[1:]
		str = serialize(n, str)
		if n != nil && (n.left != nil || n.right != nil) {
			queue = append(queue, n.left)
			queue = append(queue, n.right)
		}

	}
	return str
}
func serialize(n *TreeCount, s []string) []string {
	if n == nil {
		s = append(s, ",")
	} else {
		s = append(s, (strconv.Itoa(n.key)), ":", (strconv.Itoa(n.cnt)), ",")
	}
	return s

}

/*
amazon wants to implement a new backup system, in which files are stored into data tapes. This new system must follow the following 2 rules:

Never place more than two files on the same tape.
Files cannot be split across multiple tapes.
It's guaranteed that all tapes have the same size and that they will always be able to store the largest file.

Every time this process is executed, we already know the size of each file, and the capacity of the tapes. Having that in mind, we want to design a system that is able to count how many tapes will be required to store the backup in the most efficient way.

The parameter of your function will be a structure that will contain the file sizes and the capacity of the tapes. You must return the minimum amount of tapes required to store the files.

Example:


Input: Tape Size = 100; Files: 70, 10, 20

Output: 2
*/

func findMinTapsRequirment(fSize []int, tCapacity int) int {
	//Algorithm, Take 2 index pointer , first one start from left most while second on right most ,
	// Before using 2 index pointer, sort given file size list into acending order.
	sort.Ints(fsize)
	i := 0
	j := len(fsize) - 1
	tcnt := 0
	for i <= j {
		if fsize[j] <= tCapacity/2 { // Means All element could pair from now
			tcnt += j/2 + j%2
			break
		}
		if fsize[i]+fsize[j] <= tCapacity {
			i++
			j--
		} else {
			j--
		}
		tcnt++
	}
	return tcnt
}
func UseheapMethod(fsize []int, cap int) int {
	var minH []int
	var maxH []int
	//Build Max heap and minHeeap
	for _, r := range fsize {
		if r <= cap/2 {
			minH = append(minH, r)
		} else {
			maxH = append(maxH, r)
		}
	}
	b := peek(minH)
	for i := 0; i < len(maxH); i++ {
		a = pop(maxH)
		if a+b <= cap {
			paired = true
		}
		if paired {
			if minH.Empty() {
				b = cap
			} else {
				b = pop(minH)
			}
			paired = false
		}
		count++
	}
	minelm := len(minH)
	if !paired {
		minelm += 1
	}
	count += minelm/2 + minelm%2
	return count
}

/*
Find the first word in a stream in which it is not repeated in the rest of the stream. Please note that you are being provided a stream as a source for the characters. The stream is guaranteed to eventually terminate (i.e. return false from a call to the hasNext() method), though it could be very long. You will access this stream through the provided interface methods. A call to hasNext() will return whether the stream contains any more characters to process. A call to getNext() will return the next character to be processed in the stream. It is not possible to restart the stream.

Example:

Input: The angry dog was red. And the cat was also angry.

Output: dog

In this example, the word ‘dog’ is the first word in the stream in which it is not repeated in the stream.

Use one of the following skeletons for your solution:

Java:

package questions;

public interface Stream {
    char getNext();
    boolean hasNext();
}

*/
//Use Hash Table and Queue as Data Structure to implement this logic
//Read Word from stream and put into Hash tbl,
//If word is appear first time then put same word into end of queue
//if word is appear more than one then remove word from queue
//After end of Stream , Read the Head of queue which would give us first non- repeated word of stream

func getNextWord() string {
	var word []byte
	for hasNext() {
		c := geNext()
		if c == ' ' {
			return word
		} else {
			word = append(word, c)
		}
	}
	return nil
}
func remove(l []string, item string) {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
}

func findFirstNonRepeatedWord(input []byte) string {
	var queue []string
	var hashMap = make(map[string]int)

	for {
		word := getNextWord()
		if word == nil {
			break
		} else { //Put the word into Hash and update queue if required
			_, ok = hashMap[word]
			if !ok {
				//Appear first time , put into the Queue
				hashMap[word]++
				queue = append(queue, word)
			} else { //Word is repeated , remove it from Queue
				hashMap[word]++
				remove(queue, word)
			}
		}

	}
	return queue[0]
}

/*
Given a string of lowercase ASCII characters, find all distinct continuous palindromic sub-strings of it.
*/
/*
Algo : Iterate over each character and take it as pivot  and expend its in both
direction to found if string become palindrom or not  . Use hashMap to find
unique palindrom

*/

func findSetofPalaindrom(s []byte) {
	result := make(map[string]bool, 0)

	for i := 0; i < len(s); i++ {
		//Even length palindrom
		isPalindrom(result, s, i, i+1)
		//Odd length Palaindron
		isPalindrom(result, s, i, i)
	}
	for v, _ := range result {
		fmt.Println(v)
	}
}
func isPalindrom(r map[string]bool, s []byte, i, j int) {
	for i >= 0 && j < len(s) && (s[i] == s[j]) {
		s1 := string(s[i : j+1])
		_, ok := r[s1]
		if !ok {
			r[s1] = true
		}
		i--
		j++
	}
}

/*
Print all k-sum paths in a binary tree
A binary tree and a number k are given. Print every path in the tree with sum of the nodes in the path as k.
A path can start from any node and end at any node, i.e. they need not be root node and leaf node; and negative numbers can also be there in the tree.

Examples:

Input : k = 5
        Root of below binary tree:
           1
        /     \
      3        -1
    /   \     /   \
   2     1   4     5
        /   / \     \
       1   1   2     6

Output :
3 2
3 1 1
1 3 1
4 1
1 -1 4 1
-1 4 2
5
1 -1 5

*/
func printKpath(node *tree, k int) {
	var vector []int
	printKUtilPath(node, k, vector)
}
func printAllKsumPaths(root *tree, k int, vector []int) {
	if !root {
		return nil
	}
	vector = append(vector, root.key)
	printAllKsumPaths(root.left, k, vector)
	printAllKsumPaths(root.right, k, vector)
	//check if Sum matched with tracked node sofar
	f := 0
	for i := len(vector) - 1; i >= 0; i-- {
		f += vector
		if f == k {
			printKPath(vector, i)
		}
	}
	vector = vector[1 : len(vector)-1]
}
func printKpath(vector []int, j int) {
	for i := j; i < len(vector); i++ {
		fmt.Println(vector[i])
	}
}

/*
Make largest palindrome by changing at most K-digits
Given a string containing all digits, we need to convert this string to a palindrome by changing at most K digits.
 If many solutions are possible then print lexicographically largest one.

Examples:

Input   : str = “43435”
          k = 3
Output  : "93939"
Lexicographically largest palindrome
after 3 changes is "93939"

Input :  str = “43435”
         k = 1
Output : “53435”
Lexicographically largest palindrome
after 3 changes is “53435”

Input  : str = “12345”
         k = 1
Output : "Not Possible"
It is not possible to make str palindrome
after 1 change.


We can solve this problem using two pointers method. We start from left and right and if both digits are not equal then we replace the smaller value with larger value and decrease k by 1. We stop when the left and right pointers cross each other, after they stop if value of k is negative, then it is not possible to make string palindrome using k changes. If k is positive, then we can further maximize the string by looping once again in the same manner from left and right and converting both the digits to 9 and decreasing k by 2. If k value remains to 1 and string length is odd then we make the middle character as 9 to maximize whole value.
*/
func convertToPlanidrom(s []byte, k int) bool {
	var i int
	var palin []byte
	var j = len(s) - 1

	//Check if Palindrom possible
	for i < j {
		if s[i] != s[j] {
			//Replace lowest index with highest index value
			m := max(s[i], s[j])
			palin[i], palin[j] = m, m
			k--
		}
	}
	if k < 0 {
		return false
	}

	l := 0
	r := len(s) - 1
	for l <= r {
		if l == r {
			if k > 0 {
				palin[l] = '9'
			}
		}
		if palin[l] < '9' {
			//means no change in previous loop
			if k >= 2 && palin[l] == s[l] && palin[r] == s[r] {
				//Replace both end with '9 and decrease k by 2

				palin[l], palin[r] = '9', '9'
				k -= 2
			} else if k >= 1 && palin[l] != s[l] || palin[r] != s[r] {
				palin[l], palin[r] = '9', '9'
				k--
			}
		}
	}

	return true
}
