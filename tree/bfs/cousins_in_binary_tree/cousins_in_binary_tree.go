// 在 html 中, 操作 doc 文档的有个 api 是查找堂兄弟节点, 那么这个堂兄弟节点怎么查找呢?
// 首先需要明确什么叫做兄弟结点, 父结点
//  **父亲节点**或**父节点**：若一个节点含有子节点，则这个节点称为其子节点的父节点；
// **孩子节点**或**子节点**：一个节点含有的子树的根节点称为该节点的子节点；
// **堂兄弟节点**：父节点在同一层的节点互为堂兄弟；
// 所以, 也就是说, 堂兄弟节点在二叉树的同一层, 所以这是要给 bfs 遍历二叉树的题目

// 提取出解题需要关注的点: 按层次遍历二叉树, 记录节点所在的层次
//
package main

import (
	"fmt"
	"reflect"
)

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	data  int
}

// 保存节点所在的层次和结点的信息
type HelperBinaryTree struct {
	node  *BinaryTree //当前的节点
	level int         // 节点所在的层次
}

// 查找节点所在的层次
func helperNodeLevel(root, find *BinaryTree, level *int, curr int) {
	if root == nil {
		return
	}

	// 找到需要查找的节点, 赋值给 level
	if find == root {
		*level = curr
		return
	}

	// 遍历左子树
	if root.Left != nil {

		helperNodeLevel(root.Left, find, level, curr+1)
	}

	// 遍历右子树
	if root.Right != nil {

		helperNodeLevel(root.Right, find, level, curr+1)
	}

}

// 查找堂兄弟节点的步骤是: bfs 遍历, 如果节点在同一层, 遍历
func cousinInBinaryTree(root, find *BinaryTree) (ans []*BinaryTree) {
	if root == nil || find == nil {
		return ans
	}

	level := 0

	helperNodeLevel(root, find, &level, 0)

	que := make([]HelperBinaryTree, 0)
	front := HelperBinaryTree{node: root, level: 0}

	que = append(que, front)

	for front.level <= level && len(que) > 0 {
		front = que[0]
		// 如果当前的层次和 目标层次一致, 保存, 将节点信息入队列
		if front.level == level {
			ans = append(ans, front.node)
		}
		que = que[1:]
		// 如果左子树不为空
		if front.node.Left != nil {
			que = append(que, HelperBinaryTree{
				node:  front.node.Left,
				level: front.level + 1,
			})
		}

		// 如果右子树不为空
		if front.node.Right != nil {
			que = append(que, HelperBinaryTree{
				node:  front.node.Right,
				level: front.level + 1,
			})
		}
	}

	return ans
}
func main() {

	a := BinaryTree{data: 0}
	b := BinaryTree{data: 1}
	c := BinaryTree{data: 2}
	d := BinaryTree{data: 3}
	e := BinaryTree{data: 4}
	f := BinaryTree{data: 5}
	g := BinaryTree{data: 6}

	root := BinaryTree{data: -1}
	root.Left = &a
	a.Left = &b
	b.Left = &c
	c.Left = &d
	d.Left = &e
	e.Left = &f
	e.Right = &g

	h := BinaryTree{
		Left:  nil,
		Right: nil,
		data:  8,
	}
	i := BinaryTree{data: 9}
	j := BinaryTree{data: 10}
	k := BinaryTree{data: 11}

	l := BinaryTree{data: 12}
	m := BinaryTree{data: 13}
	n := BinaryTree{data: 14}

	root.Right = &h
	h.Right = &i
	i.Right = &j
	j.Right = &k
	k.Right = &l
	l.Left = &m
	l.Right = &n

	/**
		         -1
		        / \
		       0   8           // a
		     /      \
		    1        9        // b
		   /          \
	      2            10      // c
		 /              \
		3                11    // d
		/                  \   //
		4                  12 // e
		/ \                / \
		5  6  // f , g    13  14



	*/

	ans := cousinInBinaryTree(&root, &f)
	expected := []*BinaryTree{&f, &g, &m, &n}
	fmt.Println(ans, reflect.DeepEqual(ans, expected)) // output two node

}
