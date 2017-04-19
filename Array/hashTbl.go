package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func FindNearestRepetition(paraGraph []string) int {
	var nearestRepetitionWord int = 1<<31 - 1
	//Map key is word which appear in paragraph and val is last seen index
	// in paraGraph
	paraGraphMap := make(map[string]int)
	var diff int

	for i, s := range paraGraph {
		lastSeen, ok := paraGraphMap[s]
		if ok {
			diff = i - lastSeen
			nearestRepetitionWord = min(nearestRepetitionWord, diff)
		}
		paraGraphMap[s] = i //Update with latest seen index
	}
	return nearestRepetitionWord
}
func main() {
	var s []string = []string{"all", "work", "and", "no", "play", "makes", "for", "no", "work", "no", "fun", "and", "no", "result"}
	var n = FindNearestRepetition(s)
	fmt.Printf("Nearst Pair :%d\n", n)
	var s1 = []string{"all", "work", "and", "no", "play", "makes", "for", "no", "work", "fun", "and", "no", "fun", "fun"}
	n = FindNearestRepetition(s1)
	fmt.Printf("Part 2: Nearst Pair :%d\n", n)
}

type llist struct {
	prev        *llist
	next        *llist
	lastSeenIdx int
}
type qlist struct {
	head *llist
	tail *llist
	len  int
}
type diff struct {
	first int
	last  int
}

func findSmallestSubarryCoveringSet(s, t []string) []int {

	var smallestSet = []int{-1, -1}
	var queue *qlist
	var subSetSz = len(t)
	//Use Hash Table and double link-list to track lowest subset which covering all string in queury
	//Hash Map will key = query substring and value is its reference in double-link-list
	qHashTbl = make(map[string]*llist)
	for _, id := range t {
		qHashTbl[id] = nil
	}
	for idx, x := range s {
		old, ok := qHashTbl[x]
		if ok {
			//Sub string matched with queury sub string
			//Update double list with its latest seen index
			qListUpdate(idx, old)
			if subSetSz == qListSize { //If double list size equal to query set size , check if found smallest subset
				var subsetdiff diff
				subsetdiff := qListfirstLastIndexDiff()
				if smallest[0] == -1 && smallest[1] == -1 {
					smallest[0] = subsetdiff.first
					smallest[1] = subsetdiff.last
				} else if (subsetdiff.last - subsetdiff.first) < (smallest[1] - smallest[0]) {
					smallest[1] = subsetdiff.last
					smallest[0] = subsetdiff.first
				}

			}
		}

	}
	return smallestSet

}

func findLognestSubstringMatch(t []string) []int {
	var rslt = []int{-1, -1}
	var nonDupIdx int
	m := make(map[string]int) //Map for string as key and its location in subset is value
	for i, s := range t {
		idx, ok = m[s]
		if ok {
			//String got repeated , save its old location ,
			//Get subset strings diff
			// Clear old location
			//Moved
			if rslt[0] == -1 && rslt[1] == -1 {
				rslt[0] = 0
				rslt[1] = i
				nonDupIdx++
				m[0] = -1
			} else {
				if (i - nonDupIdx) > (rslt[1] - rslt[0]) {
					rslt[0] = nonDupIdx
					reslt[1] = i
					m[nonDupIdx] = -1
					nonDupIdx++
				}

			}

		} else {
			m[s] = i
		}
	}
	return rslt
}
