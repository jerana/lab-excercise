package main

import "fmt"

func main() {
	//arr := []int{10, 30, 1, 2, 43, 12, 33, 50, 23, 13, 43, 21, 60, 19, 18}
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//reArrange(arr, 6)
	circularCounter(a, 3)
	//fmt.Println(arr)
}

//Rearrange array element given index
func reArrange(arr []int, i int) {
	pivot := arr[i]
	fmt.Println("Pivot value :", pivot, i)
	right := len(arr) - 1
	left := 0
	//Swap the value
	arr[right], arr[i] = arr[i], arr[right]
	right = right - 1
	fmt.Println("After pivort:", arr, left, right)
	for {
		for left < right && arr[left] < pivot {
			left++
		}
		for right > left && arr[right] >= pivot {
			right--
		}
		if left == right {
			arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left]
			return
		} else {
			fmt.Println("left:", left, arr[left], "right", right, arr[right])
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
			if left > right {
				arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left]
				break
			}
		}
	}

}

func listremove(list []int, k int) []int {
	list = append(list[:k], list[k:]...)
	return list
}

func circularCounter(list []int, skip int) {
	skip = skip - 1 //index is 0
	idx := 0
	for len(list) > 0 {
		idx = (idx + skip) % len(list)
		fmt.Printf("%d ", list[idx])
		list = append(list[:idx], list[idx+1:]...)
	}
	return
}
