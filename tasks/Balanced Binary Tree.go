package main

func main() {

}

//func max(a, b int) int {
//	if b > a {
//		return b
//	}
//	return a
//}

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getMaxDepth(root.Left), getMaxDepth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	dif := getMaxDepth(root.Left) - getMaxDepth(root.Right)

	if dif > 1 || dif < -1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
