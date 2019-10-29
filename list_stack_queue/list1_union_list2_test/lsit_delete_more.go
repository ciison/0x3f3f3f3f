package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 这里使用带有头节点的单链表实现
func ListDeleteMore(head *ListNode, val int) {
	curr := head

	for curr.Next != nil {
		// 如果当前的结点不满足条件, 直接把当前的结点的指针域指向下一个结点的指针域
		if curr.Next.Val >= val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}
}

func main() {
	// 这里是构造一个链表
	node := &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}}}
	ListDeleteMore(node, 4)
	for h := node.Next; h != nil; h = h.Next {
		fmt.Print(h.Val, " ")
	}
}
