package main

func main() {

}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return ((p != nil || q == nil) && (p == nil || q != nil)) &&
		((p == nil && q == nil) || (p.Val == q.Val && isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)))
}
