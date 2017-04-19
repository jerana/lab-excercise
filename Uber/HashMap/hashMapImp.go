package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("vim-go")
	set(20, 1)
	set(40, 11)
	set(50, 12)
	set(60, 13)
	set(70, 14)
	set(80, 15)
	set(90, 101)
	set(100, 01)
	//	fmt.Println("Get at index 50:", get(50))
	i := 2
	for {
		//v := getRandonKey()
		//		fmt.Println("random key:", v, "val:", get(v))
		if i > 2 {
			break
		}
		i++

	}

	list := countGreater(Nums)
	fmt.Println("list:", list)

}
func printList() {
	fmt.Println("Len of list:", len(availKey))
	for idx, r := range availKey {
		fmt.Println("", idx, ":", r.key, r.val)
	}
}

/*
Design a key-value store to support methods like get, put and getRandom.

getRandom -> Should generate a random number and return the key-value stored at this random number

*/
//Will use HashMap to store key-val, But wrapper on top of value to store index  which is pointing to index of list which store key/val pair

type keyValPair struct {
	key int
	val int
}

//Use this list index to hashMap as Value
var availKey []keyValPair
var hashMap = make(map[int]int, 0)

func set(key, val int) {
	v, ok := hashMap[key]
	if !ok {
		//Store List Index into HashMap
		hashMap[key] = len(availKey)
		//n := new(keyValPair)
		var n keyValPair
		n.key = key
		n.val = val
		availKey = append(availKey, n)
	} else {
		//Just Update the entry
		availKey[v].val = val
	}
}
func get(key int) int {
	v, ok := hashMap[key]
	if !ok {
		return -1
	}
	//V: is index of list
	return availKey[v].val
}
func getRandonKey() int {
	return availKey[rand.Intn(len(availKey))].key
}
