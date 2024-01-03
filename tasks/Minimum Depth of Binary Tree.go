package main

func main() {

}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	path := make(map[*TreeNode]*TreeNode)
	path[root] = nil

	var node *TreeNode

	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		if node.Left == nil && node.Right == nil {
			break
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			path[node.Left] = node
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			path[node.Right] = node
		}
	}

	depth := 0

	for node != nil {
		depth++
		node = path[node]
	}

	return depth
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := 1<<31 - 1
	if root.Left != nil {
		depth = minDepth(root.Left)
	}
	if root.Right != nil {
		depth = min(depth, minDepth(root.Right))
	}
	return depth + 1
}
