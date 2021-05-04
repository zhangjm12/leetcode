package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var f = make(map[*TreeNode]int)
var g = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	dfs(root)

	return (Max(f[root], g[root]))
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)

	f[root] = root.Val + g[root.Left] + g[root.Right]
	g[root] = Max(f[root.Left], g[root.Left]) + Max(f[root.Right], g[root.Right])
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 1}

	fmt.Printf("%v\n", root)
	fmt.Println(rob(root))
}
