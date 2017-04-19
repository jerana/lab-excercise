package main

import "fmt"

/*

Mooshak the mouse has been placed in a maze.There is a huge chunk of cheese somewhere in the maze.
The maze is represented as a two-dimensional array of integers, where o represents walls, 1 represents paths where Mooshak can move, and 9 represents the huge chunk of cheese.Mooshak starts in the top-left corner at 0,0.

Write a method isPath of class Maze Path to determine if Mooshak can reach the huge chunk of cheese. The input to isPath consists of a two dimensional array grid for the maze matrix.

The method should return 1 if there is a path from Mooshak to the cheese, and 0 if not.
Mooshak is not allowed to leave the maze or climb on walls/

Example 8x8 maze where Mooshak can get the cheese.

1 0 1 1 1 0 0 1

1 0 0 0 1 1 1 1

1 0 0 0 0 0 0 0

1 0 1 0 9 0 1 1

1 1 1 0 1 0 0 1

1 0 1 0 1 1 0 1

1 0 0 0 0 1 0 1

1 1 1 1 1 1 1 1

*/
var maze = [][]int{
	{1, 0, 1, 1, 1, 0, 0, 1},
	{1, 0, 0, 0, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0},
	{1, 0, 1, 0, 9, 0, 1, 1},
	{1, 1, 1, 1, 0, 0, 0, 1},
	{1, 0, 1, 0, 1, 1, 0, 1},
	{1, 0, 0, 0, 0, 1, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

func main() {

	fmt.Println("Is MAZE:", isPathHelper(maze, 0, 0))
	m := len(maze)
	n := len(maze[0])
	if isPathHelper(maze, 0, 0) {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Println(maze[i], " ")
			}
			fmt.Println("\n")
		}
	}
}

func isPathHelper(matrix [][]int, x, y int) bool {
	m := len(matrix)
	n := len(matrix[0])
	//first do sanity check
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	//Base condition
	if matrix[x][y] == 9 {
		return true
	}
	if matrix[x][y] != 1 {
		return false
	}
	//Mark x,y is part of solution path
	if matrix[x][y] == 1 {
		matrix[x][y] = 3
	}
	if isPathHelper(matrix, x-1, y) {
		return true
	}
	if isPathHelper(matrix, x, y-1) {
		return true
	}
	if isPathHelper(matrix, x+1, y) {
		return true
	}
	if isPathHelper(matrix, x, y+1) {
		return true
	}
	matrix[x][y] = 0
	return false
}
