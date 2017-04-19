package main

import "fmt"

func isPalindromPossible(s string) bool {
	var oddCharCnt int
	var charMap = make(map[byte]int)
	for _, c := range s {
		charMap[byte(c)]++
	}
	for _, val := range charMap {
		if (val % 2) != 0 {
			oddCharCnt++
			if oddCharCnt > 1 {
				return false
			}
		}
	}
	return true
}
func findCountBits(num int) *int {
	bitArray := new((num + 1) * sizeof(int))
	for i = 1; i <= n; i++ {
		bitArray[i] = bitArray[i>>1] + (i & 1)
	}
	return bitArray

}
func main() {
	var s string
	fmt.Println("Enter String :")
	fmt.Scanln(&s)
	fmt.Println("Readr String :", s)
	fmt.Println("String is  anagram ", isPalindromPossible(s))

}

func Plusone(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		if isDigit(arr[i]) < 9 {
			arr[i] = arr[i] + 1
			return arr
		} else {
			arr[i] = 0
		}
	}
	arr[0] = 0
	arr = append(arr, 1)
	return arr
}
