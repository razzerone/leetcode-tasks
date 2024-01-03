package main

func main() {
	print(insertionSortList(&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}).Val)
}

//TODO

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	res := head
	head = head.Next
	res.Next.Next = nil

}
