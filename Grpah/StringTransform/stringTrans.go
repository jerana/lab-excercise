package main

import "fmt"

func main() {
	start := "hit"
	end := "cog"
	fmt.Println("Start str:", start)
	fmt.Println("End str:", end)
	t, r := isSourceTransformIntoTarget([]byte(start), []byte(end), dict)
	fmt.Println("Val:", t, r)

}

/*
Given String S can be transform into t , by changing one character sequentially , intermeidate Sequence string is form dictionary
DB .

*/
//Algo; Build Graph Using , Source and Intermediate String as Vertex node, and transforming one string into other by changeing  1 character is consider is edge
/*

Given two words (start and end), and a dictionary, find the length of shortest transformation sequence from start to end, such that only one letter can be changed at a time and each intermediate word must exist in the dictionary. For example, given:

start = "hit"
end = "cog"
One shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog", the program should return its length 5.
*/

var dict map[string]bool
var production = []string{"hot", "dot", "dog", "lot", "log"}
var visit map[string]bool

func init() {
	dict = make(map[string]bool, 1)
	for _, r := range production {
		dict[r] = true
	}
	visit = make(map[string]bool, 1)
}

type GraphVertex struct {
	d int
	s []byte
}

func isSourceTransformIntoTarget(src, target []byte, dict map[string]bool) (bool, int) {

	var queue = make([]*GraphVertex, 0)
	//Insert first Src string into Queue
	root := new(GraphVertex)
	root.d = 1
	root.s = src
	queue = append(queue, root)
	//Append end string into dictnary
	dict[string(target)] = true
	for len(queue) > 0 {
		var s0 = queue[0] //Get Head of Queue
		if s0 == nil {
			fmt.Println("Nil ptr")
			return false, -1
		}
		fmt.Println(s0.d, "Walk str:", string(s0.s), s0.d, "target:", string(target))
		if string(s0.s) == string(target) {
			return true, s0.d
		}
		for i := 0; i < len(s0.s); i++ {
			tmp := s0.s[i]
			for j := 0; j < 26; j++ {
				s0.s[i] = byte('a' + j)
				_, ok := dict[string(s0.s)]
				if ok {
					n := new(GraphVertex)
					n.d = s0.d + 1
					n.s = make([]byte, len(s0.s))
					copy(n.s, s0.s)
					fmt.Println("Append node:", string(n.s))
					queue = append(queue, n)
					delete(dict, string(n.s))
				}

			}
			s0.s[i] = tmp
			fmt.Println("original ch", string(src[i]))
		}
		fmt.Println("remove str:", queue[0], string(s0.s))
		fmt.Println("Queue:", queue)
		queue = queue[1:]

	}
	return false, 0
}
