package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*Given the list of team pair which have winning and lossing information
find out given team A to B if there is any possible sequence or not
*/
type pairTeam struct {
	winTeam  string
	lossTeam string
}
type GraphNode struct {
	team    string
	winlist []string
}

func BuildGraph(list []*pairTeam) *GraphNode {
	m := make(map[string][]string, 0)
	for _, l := range list {
		m[l.winTeam] = append(m[l.winTeam], l.lossTeam)
	}
	return m
}

var visited = make(map[string]bool, 0)

func isTeamSequence(G map[string][]string, src, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}
	_, ok := visited[src]
	if ok {
		return false
	}
	var gConn []string
	gConn, ok = G[src]
	if !ok {
		return false
	}
	visited[src] = true
	for _, n := range gConn {
		if isTeamSequence(G, n, dst, visited) {
			return true
		} else {
			return false
		}

	}
	return false
}

/*
find maze path from start to end  position in 2 dimansion array where there is some cell are blocked by number 1
*/
type point struct {
	x int
	y int
}

var neigCorodinat []point = []point{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func isValidNeig(matrix [][]int, x, y int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if (x >= 0 || x < n) && (y >= 0 || y < m) && matrix[x][y] != 1 {
		return true
	}
	return false
}

/*Track the path from start to end
Algo : Use Graph concept, consider each Matrix node white as vertex node and make reachablity path from start to target node;
During walk track each vertex into Queue if vetrix is reachable to its neigbhor without blocked
*/
func findPathFromStartToTaget(matrix [][]int, start, end point) []point {
	var path []point //Queue to track Path from start to end
	matrix[start.x][start.y] = 1
	path = append(path, start)
	if !searchPathFromStartToEnd(matrix, start, end, path) {
		path = path[:len(path)-1]
	}
	return path
}
func searchPathFromStartToEnd(matrix, start, end, path) bool {
	//Reach at Target destination
	if start == end {
		return true
	}
	matrix[start.x][start.y] = 1 //vertex is visited
	var new point
	for _, idx := range point { //check all its 4 neighors
		new.x = start + idx[0]
		new.y = start + idx[1]
		if isValidNeig(matrix, new.x, new.y) {
			//record the Visited Vertex
			path = append(path, new)
			if searchPathFromStartToEnd(matrix, new, end, path) {
				return true
			}
			//Pop the path,
			path = path[:len(path)-1]
		}
	}
	return false
}

/*
Let A be a 2D array whoes entries are either W or B . Write  a program
that replace all Ws that cannot reach the boundary with a B.
*/

func ConvertWTOB(matrix [][]int, p point) {
	m := len(matrix)
	n := len(matrix[0])
	var visited [][]bool
	visite = make([]bool, m)
	for i := 0; i < m; i++ {
		visite[i] = make([]bool, n)
	}
	//Walk First Row and Last Row and track All white from neigbors white
	//cell is reachble
	for i := 0; i < n; i++ {
		if matrix[i][0] == 'W' {
			markWhiteItsNeigh(matrix, i, 0, visited)
		}
	}
	lastRow := len(matrix) - 1
	for i := 0; i < n; i++ {
		if matrix[i][lastRow] == 'W' {
			markWhiteItsNeigh(matrix, i, lastRow, visited)
		}
	}
	//Walk first Column and Last Column and track all neighbors
	//white cells
	lastCol := len(matrix[0]) - 1
	for i := 0; i < m; i++ {
		if matrix[0][i] == 'W' {
			markWhiteItsNeigh(matrix, 0, i, visited)
		}
	}
	for i := 0; i < m; i++ {
		if matrix[lastCol][0] == 'W' {
			markWhiteItsNeigh(matrix, lastCol, i, visited)
		}
	}
	//After all above process, Walk Whole matrix , and flip all whites
	//which are not reachable to boder
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 'W' && !visited[i][j] {
				matrix[i][j] = 'B'
			}

		}
	}
}
func markWhiteItsNeig(matrix [][]int, x, y int, visited [][]int) {

	var queue []point
	var p point
	p.x = x
	p.y = y
	if x > 0 && x < len(matrix) && y > 0 && y <= len(matrix[0]) && !visited[x][y] {
		queue = append(queue, p)
		visited[x][y] = true
		for len(queue) > 0 {
			n := queue[0]
			for _, r := range neigCorodinat {
				n.x = n.x + r[0]
				n.y = n.y + r[1]
				if matrix[n.x][n.y] == 'W' {
					queue = append(queue, n)
					visited[n.x][n.y] = true
				}
			}
			queue = queue[1:]

		}
	}
}
