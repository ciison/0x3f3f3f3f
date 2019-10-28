package main

import (
	"fmt"
	"math"
)

// 使用单链表实现栈, 不用头结点
type Node struct {
	Val int
	next *Node
}

func (c * Node) Push(val int ) *Node{
	tmp:= new(Node)
	tmp.Val = val
	tmp.next = c
	c = tmp
	return c
}
func (c *Node) Pop() *Node{
	if c == nil {
		return nil
	}
	c = c.next
	return c
}
func (c *Node) Top() int {
	if c == nil {
		return math.MinInt32
	}
	return c.Val
}
func main() {
	node := new(Node)
	node.Val = 9
	node = node.Push(8)
	fmt.Println(node, node.next)
	fmt.Println("top", node.Top())
	node = node.Pop()
	fmt.Println(node.Top())
}
