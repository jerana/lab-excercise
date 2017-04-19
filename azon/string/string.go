package main

import "fmt"

func main() {
	//str := "abaaa"
	//str1 := "geek"
	s := "ame is jeewan"
	fmt.Printf("First unique char: %c\n", byte(findFirstUniqueChar(s)))
	fmt.Printf("First unique char: %c\n", byte(findFirstUniqueC(s)))
	//findSetofPalaindrom([]byte(str))
	//findSetofPalaindrom([]byte(str1))
}

/*
Find first unique char in a string
*/
//Algorithm , Walk string backward and check if there is its duplicate entry exist .
func findCharInStr(s string, c rune) (bool, int) {
	for idx, r := range s {
		if r == c {
			return true, idx
		}
	}
	return false, -1
}

func findFirstUniqueChar(str string) rune {
	var result rune
	for i := len(str) - 1; i >= 0; i-- {
		r, c := findCharInStr(str, rune(str[i]))
		if r && c == i {
			result = rune(str[i])
		}
	}
	return result
}

//Another method is Use HashMap and List
//HashMap keep all char from string and list maintained all unique char
func findFirstUniqueC(str string) byte {
	hashMap := make(map[rune]int, 0)
	//If Hashmap value is index in its list , which help to remove it if its duplicate entry got found

	uniqueList := make([]rune, 0)
	for _, c := range str {
		idx, ok := hashMap[c]
		if ok && idx != 0 {
			//remove its entry from unique list
			//a = append(a[:i], a[i+1:]...)
			uniqueList = append(uniqueList[:idx], uniqueList[idx+1:]...)
			hashMap[c] = 0
		} else {
			//Add this into hashMap and unqiue list
			hashMap[c] = len(uniqueList)
			uniqueList = append(uniqueList, c)
		}

	}
	//return head of list as first unique entry
	return byte(uniqueList[0])
}

/*
Given a string of lowercase ASCII characters, find all distinct continuous palindromic sub-strings of it.
*/
/*
Algo : Iterate over each character and take it as pivot  and expend its in bothdirection to found if string become palindrom or not  . Use hashMap to find
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

type tree struct {
	key   int
	left  *tree
	right *tree
}

/*Mirror of Binary tree with Recursion*/
func mirrorBTree(n *tree) {
	if n == nil {
		return
	}
	swapNodePtr(n)
	mirrorBTree(n.left)
	mirrorBTree(n.right)
	return
}
func swapNodePtr(n *tree) {
	tmp := n.left
	n.left = n.right
	n.right = tmp
}

//Non-recursive solution
//Use Queue and push node into queue and pop node and swap its childs node

func mirrorTree(n *tree) {
	if n == nil {
		return
	}
	if n.left == nil && n.right == nil {
		return
	}
	var queue []*tree
	queue = append(queue, n)
	for len(queue) > 0 {
		m := queue[0]
		queue = queue[1:]
		swapNodePtr(m)
		queue = append(queue, m.left)
		queue = append(queue, m.right)
	}
}
