package main

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	tree := &TreeNode{Val: nums[mid]}
	tree.Left = sortedArrayToBST(nums[:mid])
	tree.Right = sortedArrayToBST(nums[mid+1:])

	return tree
}
