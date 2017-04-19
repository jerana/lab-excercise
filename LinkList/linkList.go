package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type llist struct {
	val  int
	next llist
}

func reverseList(head *llist) {
}
