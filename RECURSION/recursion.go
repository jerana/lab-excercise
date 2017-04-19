package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//Tower of Hanoi problam

//If First move top n-1 into temp tower
//then move last ring from src tower to dest tower
//After this move all element from temp tower to dest tower

type list []int

var tower []list

func towerofHanoi() {
	var arr = []int{1, 2, 3, 4, 5}
	var num int
	computeTowerofHanoi(num, tower, tower[0], tower[1], tower[2])
}

func computeTowerofHanoi(numTower, from, to, use []tower) {
	if numTower {
		computeTowerofHanoi(numTower-1, from, use, to)
		to[0] = from[0]
		computeTowerofHanoi(numTower-1, use, to, from)
	}
}
