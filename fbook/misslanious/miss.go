package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	s := "MMXIV"
	s1 := "CXVI"
	fmt.Println("convert Value:", convertRomantoDecimal(s))
	fmt.Println("convert Value:", convertRomantoDecimal(s1))

}
func convertRomantoDecimal(s string) int {
	hashMap := make(map[byte]int)
	hashMap['M'] = 1000
	hashMap['D'] = 500
	hashMap['C'] = 100
	hashMap['L'] = 50
	hashMap['X'] = 10
	hashMap['V'] = 5
	hashMap['I'] = 1

	i := len(s) - 1
	var result int

	for i >= 0 {
		fmt.Println("i :", i, "hashMap:", hashMap[s[i]])
		if i+1 < len(s) && hashMap[s[i+1]] > hashMap[s[i]] {
			fmt.Println("Substract ", hashMap[s[i]], "from result", result)
			result -= hashMap[s[i]]
		} else {
			fmt.Println("Add ", hashMap[s[i]], "from result", result)
			result += hashMap[s[i]]
		}

		i--
	}
	return result
}

/*Given an array A of integers, find the index of values that satisfy A + B =C + D, where A,B,C & D are integers values in the array. Find all combinations of quadruples.

output all indexes of quadruple into format List<List<Integer>> indexofQuadruples

similar to leetcode 4sum, or has better solution ?
*/

type point struct {
	x int
	y int
}

func equivalentPairsOf4(arr []int) [][]int {
	hMap := make(map[int][]point)
	var pair point

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			_, ok := hMap[sum]
			if !ok {
				hMap[sum] = make([]int, 0)
			}
			pair.x = arr[i]
			pair.y = arr[j]
			hMap[sum] = append(hMap[sum], pair)
		}
	}
	var result []pair
	for _, list := range hMap {
		if len(list) > 1 { //If more than 1 entry then adds its to arry
			for _, p := range list {
				//Add pair index into result list
				result = appen(result, p)
			}

		}
	}
	return result
}

//Remove comments from a given C/C++ program
/*
The idea is to maintain two flag variables, one to indicate that a single line comment is started, another to indicate that a multiline comment is started. When a flag is set, we look for the end of comment and ignore all characters between start and end.
*/

func removeComments(pram string) string {

	var result []rune
	size = len(param) - 1
	var i int
	var s_cmt bool
	var m_cmt bool
	for i := 0; i < size; i++ {
		if s_cmt && pram[i] == '\n' { //Check for end of line
			s_cmt = false
		} else if m_cmt && pram[i] == '*' && pram[i+1] == '/' { //Check for end of comment
			m_cmt = false
			i++
		} else if s_cmt || m_cmt {
			continue
		} else if param[i] == '/' && param[i+1] == '/' { //Detect single line comment
			s_cmt = true
			i++
		} else if param[i] == '/' && param[i+1] == '*' { //multiple line comment
			m_cmt = true
			i++
		} else {
			result = append(result, param[i])
		}

	}
	return result

}
