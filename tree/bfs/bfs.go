package main

import (
	"fmt"
	"reflect"
)

// 广度优先遍历二叉树
type binaryTree struct {
	left  *binaryTree
	right *binaryTree
	data  int
}

func bfs(root *binaryTree) (ans []int) {
	// 树的广度遍历
	helper := make([]*binaryTree, 0)
	if root == nil {
		return
	}

	helper = append(helper, root)

	// 如果 长度为 0, 表示队列中需要遍历元素
	for len(helper) != 0 {
		node := helper[0]
		// 将第一个元素出队列
		helper = helper[1:]

		ans = append(ans, node.data)
		// 如果有左子树, 将左子树入队列
		if node.left != nil {
			helper = append(helper, node.left)
		}
		// 如果有右子树, 将右子树入队列
		if node.right != nil {
			helper = append(helper, node.right)
		}
	}

	return ans
}
func main() {
	a := binaryTree{data: 1}
	b := binaryTree{data: 2}
	c := binaryTree{data: 3}
	d := binaryTree{data: 4}

	root := binaryTree{data: 0}
	root.left = &a
	a.left = &b
	root.right = &c
	c.right = &d
	//
	/**
		 0
		/ \
	   1   3
	  /     \
	 2	     4
	// :output 0 1 3 2 4
	*/
	ans := bfs(&root)
	expected := []int{0, 1, 3, 2, 4}
	fmt.Println(ans)
	fmt.Println(reflect.DeepEqual(ans, expected))
}
