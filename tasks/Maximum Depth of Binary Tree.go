package main

func main() {

}

//func max(x, y int) int {
//	if y > x {
//		return y
//	}
//	return x
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
