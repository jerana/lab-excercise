package main

import "fmt"

func main() {
	var root *tree = treeNodeAlloc(2)
	root.left = treeNodeAlloc(3)
	root.right = treeNodeAlloc(5)
	root.left.left = treeNodeAlloc(4)
	root.left.right = treeNodeAlloc(8)
	root.right.left = treeNodeAlloc(6)
	root.right.right = treeNodeAlloc(-2)
	root.right.right.right = treeNodeAlloc(2)
	findAllPathEqualSum(root, 7)

	//printKnodeFromTargetBinary(root, root.left.right, 4)
	//fmt.Println("Find LCA(4,5):", findLCABinaryTree(root, 2, 6))
	//fmt.Println("New LCA(4,5):", FindLCAnode(root, root.left.left, root.left.right))

}

/*Find LCA of Given 2 Nodes*/

func findLCAUtil(root *tree, n1, n2 int, d1, d2 *int, level int, dist *int) *tree {
	if root == nil {
		return nil
	}
	if root.key == n1 {
		*d1 = level
		return root
	}
	if root.key == n2 {
		*d2 = level
		return root
	}
	lca_left := findLCAUtil(root.left, n1, n2, d1, d2, level+1, dist)
	lca_right := findLCAUtil(root.right, n1, n2, d1, d2, level+1, dist)
	if lca_left != nil && lca_right != nil { //both are not null , mean current node is LCA
		*dist = *d1 + *d2 - 2*level
		return root
	}
	if lca_left != nil { //Need to check LCA in left subtree
		return lca_left
	} else {
		return lca_right //Check LCA in right subtree
	}

}
func findDistanceBetweenNodes(root *tree, n1, n2 int) int {
	var d1, d2 int
	d1, d2 = -1, -1
	var dist int
	var lca *tree

	lca = findLCAUtil(root, n1, n2, &d1, &d2, 0, &dist)
	if d1 != -1 && d2 != -1 {
		return dist
	}
	if d1 != -1 {
		dist = findLevel(lca, d2, 0)
	} else if d2 != -1 {
		dist = findLevel(lca, d1, 0)
	}
	return -1
}
func findLevel(root *tree, key, level int) int {
	if root == nil {
		return -1
	}
	if root.key == key {
		return level
	}
	var l = findLevel(root.left, key, level+1)
	if l != -1 {
		return l
	} else {
		return findLevel(root.right, key, level+1)
	}

}

func findAllPathEqualSum(root *tree, sum int) {
	path := make([]int, 0)
	result := make([][]int, 0)
	fmt.Println("Print all path which sum is :")
	fmt.Scanf("%d", &tSum)
	findPathEqualToSum(root, 0, path, result)
	if len(result) > 0 {
		for _, r := range result {
			fmt.Println("list:", r)
		}
	}

}

var tSum int

func findPathEqualToSum(root *tree, cSum int, path []int, result [][]int) {
	if root == nil {
		return
	}
	path = append(path, root.key)
	cSum += root.key
	if cSum == tSum { //Means given path exist and save into result list
		fmt.Println("Adding path:", path)
		result = append(result, path)
	}
	findPathEqualToSum(root.left, cSum, path, result)
	findPathEqualToSum(root.right, cSum, path, result)
	path = path[:len(path)-1]
	pathN := make([]int, 0)
	findPathEqualToSum(root.left, 0, pathN, result)
	findPathEqualToSum(root.right, 0, pathN, result)

	return
}

//Mirror of Tree
func MirrioTree(root *tree, m *tree) {

	if root == nil {
		return root
	}
	m.key = root.key
	MirrorTree(root.right, m.left)
	MirrorTree(root.left, m.right)

}
