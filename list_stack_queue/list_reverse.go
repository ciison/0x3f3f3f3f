package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Contains(L1, L2 *ListNode) bool {
	// 这里的实现使用了两重的循环
	for h1 := L1.Next; h1 != nil; h1 = h1.Next {
		h2 := L2.Next
		for ; h2 != nil; h2 = h2.Next {
			if h2.Val == h1.Val {
				break
			}
		}
		// 如果 h2 == nil 表示寻找是否等于 h1.Val 的值的时候, 在 L2 中没有找到
		if h2 == nil {
			return false
		}
	}
	return true
}
func main() {
	node := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}}}
	node2 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}

	fmt.Println("contains", Contains(node, node2))
}
