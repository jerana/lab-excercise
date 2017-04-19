package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Coordinate struct {
	x, y int
}

var maze = [3][3]int{}

func isFessiableCorrdinate(cur Coordinate, m maze) bool {
	if cur.x >= 0 && cur.x <= len(maze) &&
		(cur.y >= 0 && cur.y <= len(maze[0])) && maze[cur.x][cur.y] == 1 {
		return true
	}
	return false
}

func searchMaze(maze [][]int, s, t Coordinate) []SearchPath {
	var sPath []Coordinate
	maze[s.x][s.y] = 0
	sPath = append(sPath, s)
	if !SearchMazePath(s, t, sPath) {
		sPath = sPath[0:len(sPath)]
	}
	return sPath
}
func searchMazeBFS(maze [][]int, s, t Coordinate) []path {
	var sPath []Coordinate
	var kfactor = []Coordinate{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	var queue = []Coordinate{}
	//Add Starting Coordinate location into Search path
	queue = append(queue, s)
	for len(queue) != 0 {
		//Get top element don't remove it
		n := queue[0]
		if n.x == t.x && n.y == t.y {
			sPath = append(sPath, n)
			return sPath
		}
		//Find all Valid Neigbhor which is having White
		for _, k := range kfactor {
			var new Coordinate
			new.x = n.x + k.x
			new.y = n.y + k.y
			if isFessiableCorrdinate(new, maze) {
				sPath = append(sPath, n)
				queue = append(queue, new)
			}
		}
		//Discard Top element
		queue = queue[1:]
	}
	return nil
}
func SearchMazePath(s, t Coordinate, sPath []Coordinate) bool {
	var kfactor = []Coordinate{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for _, k := range kfactor {
		var n Coordinate
		n.x = s.x + k.x
		n.y = s.y + k.y
		if isFessiableCorrdinate(n, maze) {
			maze[n.x][n.y] = 0
			sPath = append(sPath, n)
			if SearchMazePath(n, t, sPath) {
				return true
			}
			sPath = sPath[0:len(sPath)]
		}
	}
	return false
}

func updateMatrix(m [][]int, point Coordinate) {
	var delta = []Coordinate{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	x := m[point.x][point.y]
	if x != 1 {
		m[point.x][point.y] = 1
		for _, d := range delta {
			var new Coordinate
			new.x = point.x + d.x
			new.y = point.y + d.y
			updateMatrix(m, new)
		}
	}
}
func isPossiblePaint(m *[4][4]int, point int) bool {
	if point.x >= 0 && point.x <= len(m) &&
		point.y >= 0 && point.y <= len(m[4]) &&
		m[x][y] != 1 {
		return true
	} else {
		return false
	}
}
func isSamePathPoint(m [][]int, point int, color int) bool {
	if point.x >= 0 && point.x <= len(m) &&
		point.y >= 0 && point.y <= len(m[0]) &&
		m[point.x][point.y] == color {
		return true
	}
	return false
}

//Using BFS logic, Consider matrix index as verties and Adjacencts points are
//its edges
func matrixFlipColor(matrix [][]int, x, y int) {
	var kfactor = []Coordinate{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	var vQueue []Coordinate
	var color = matrix[x][y]
	matrix[x][y] = !color
	for len(vQueue) != 0 { //Till Queue is not getting Empty
		var new Coordinate
		n := vQueue[0]
		for _, k := range kfactor {
			new.x = n.x + k.x
			new.y = n.y + k.y
			if isSamePathPoint(matrix, new, color) {
				matrix[new.x][new.y] = !color
				vQueue = append(vQueue, new)
			}
		}
		//Pop the element
		vQueue = vQueue[1:]
	}
}

//Clone Grpah
type GraphVertex struct {
	lable int
	edges []GraphVertex
}

func clone(G *GraphVertex) *GraphVertex {
	var gQueue []GraphVertex
	var gHashtbl = make([*GraphVertex]*GraphVertex, 0)
	if G == nil {
		return nil
	}
	cGprah := new(GraphVertex)
	cGrpah.lable = G.lable
	ghashtbl[G] = cGraph
	for len(gQueue) != 0 {
		v := gQueue[0]
		gQueue = gQueue[1:]
		for _, e := range v.edges {
			_, ok := gHashTbl[e]
			if !ok {
				//Allocate new clone Graph node
				n := new(GraphVertex)
				gHashTbl[e] = n
				gQueue = append(gQueue, n)
			}
			vEdges := v.edges
			//Add edges from orinial to clone Graph node
			vEdges = append(vEdges, e)
		}

	}
	return cGraph
}

//Find String S can be Transform into T with sequence S0,S1...Sn-1 to T . Providded D as Disctionary
type stringDistance struct {
	s string
	d int
}

func findSTransformedToT(s, t string, m map[string]string) bool {

	var q []stringDistance
	q = append(s)
	erase(m, s)
	for len(q) != 0 {
		n := q[0]
		if n.s == t {
			return d
		}

		for i := 0; i < len(s); i++ {
			for j := 0; j < 26; j++ {
				s1 = string(byte(s[i] + 'a' +j
			}
		}

	}

}
type chFreq struct {
	c byte 
	f int
}


void Heapify(



