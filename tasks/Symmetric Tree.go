package main

func main() {

}

func help(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) ||
		(p != nil && q != nil && p.Val == q.Val && help(p.Right, q.Left) && help(p.Left, q.Right))
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || help(root.Left, root.Right)
}
