package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var nodes []int
	if root != nil {
		nodes = append(nodes, inorderTraversal(root.Left)...)
		nodes = append(nodes, root.Val)
		nodes = append(nodes, inorderTraversal(root.Right)...)
	}
	return nodes
}
