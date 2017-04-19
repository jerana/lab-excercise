package main

//Define Process node as Veterix node
type GraphVertex struct {
	color int
	edges []*GraphVertex
}

const (
	white = iota
	gray
	black
)

//This design will detect  dealock in system

func isDeadLock(G []*GraphVertex) bool {
	for _, r := range G {
		if r.color == white && hasCycle(r, nil) {
			return true
		}
	}
	return false
}
func hasCycle(cur *GraphVertex, pre *GraphVertex) bool {
	if cur.color == gray {
		return false
	}
	cur.color = gray
	//Walk into its neigh vertex
	for _, next := range cur.edges {
		if next != pre && cur.color != black {
			if hasCycle(next, cur) {
				return true
			}
		}
	}
	cur.color = black
	return false

}

//Clone Graph
type GraphNode struct {
	label int
	neig  []*GraphNode
}

func copyNode(G *GraphNode) *GraphNode {
	n := new(GraphNode)
	n.label = G.label
	n.neig = make([]*GraphNode, len(G.neig))
}

var hashMap = make(map[*GraphNode]bool, 0)

func GrpahClone(G *GraphNode) *GraphNode {
	if G == nil {
		return nil
	}
	var C *GraphNode = copyNode(GraphNode)
	GraphCloneHelper(G, C)
	return C
}
func GraphNodeHelper(G, C *GraphNode) {

	if G == nil || hashMap[G] {
		return
	}
	hashMap[G] = true
	for _, r := range G.neig {
		t1 := copyNode(r)
		C.neig = append(C.neig, t1)
		GraphCloneHelper(r, t1)
	}
	return

}
