package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type GraphVertex struct {
	label int
	edges []*GraphVertex
}

func copyNode(G GraphVertex) *GraphVertex {
	n := new(GraphVertex)
	n.lable = G.label
	n.edges = make([]*GraphVertex, len(G.edges))

}

//Mapping of original Vertext to clone Vertex
var hashMap = make(map[*GraphVertex]*GraphVertex)

//Use BFS traversal
func CloneGraph(G *GraphVertex) *GraphVertex {
	var queue []*GraphVertex = make([]*GraphVertex, 0)

	if G == nil {
		return nil
	}
	hashMap[G] = copyNode(G)
	queue = append(queue, G)
	//till queue is empty
	for len(queue) > 0 {
		n := queue[0]
		for _, e := range n.neig {
			v, ok := hashMap[e]
			if !ok {
				hashMap[e] = copyNode(e)
				queue = append(queue, e)
			}
			hashMap[n].neig = append(hashMap[n].neig, hashMap[e])
		}
		queue = queue[1:]
	}
	return hashMap[G]
}
