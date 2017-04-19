package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Arra:", arr)
	s := convertBigNumber(arr)
	fmt.Println(s)
}

func reArragneArray(arr []int) []int {
	//Swap A[i] and A[i+1] if i is even and A[i] > A[i+1]
	//Swap A[i] and A[i+1], if i is odd and A[i] < A[i+1]
	for i = 0; i < len(arr); i++ {
		if ((i & 1) && arr[i] < arr[i+1]) || ((i&1 == 0) && arr[i] > arr[i+1]) {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}

func convertBigNumber(arr []int) string {
	var buffer bytes.Buffer
	s := comp(arr)
	sort.Sort(s)
	fmt.Println(s)
	for _, n := range s {
		buffer.WriteString(strconv.Itoa(n))
	}
	return buffer.String()
}

var arr = []int{3, 30, 34, 5, 9}

type comp []int

func (a comp) Len() int {
	return len(a)
}
func findNumDigit(x int) int {
	i := 0
	for x > 0 {
		i++
		x /= 10
	}
	return i
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func (a comp) Less(i, j int) bool {
	r1d := findNumDigit(a[i])
	r2d := findNumDigit(a[j])

	var r1, r2 int
	r1 = a[i]
	for r2d > 0 {
		r1 = r1 * 10
		r2d--
	}
	r1 = r1 + a[j]
	r2 = a[j]
	for r1d > 0 {
		r2 = r2 * 10
		r1d--
	}
	r2 = r2 + a[i]
	if r1 > r2 {
		return true
	}
	return false
}
func (a comp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
