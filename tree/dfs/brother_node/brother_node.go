// 在 html 中, 操作 doc 文档的有个 api 是查找兄弟节点, 那么这个兄弟节点怎么查找呢?
// 首先需要明确什么叫做兄弟结点, 父结点
//  **父亲节点**或**父节点**：若一个节点含有子节点，则这个节点称为其子节点的父节点；
// **孩子节点**或**子节点**：一个节点含有的子树的根节点称为该节点的子节点；
// **兄弟节点**：具有相同父节点的节点互称为兄弟节点；
// **堂兄弟节点**：父节点在同一层的节点互为堂兄弟；
// **节点的祖先**：从根到该节点所经分支上的所有节点
// 所以, 也就是说, 兄弟节点具有相同的父结点
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

// 为什么需要遍历呢? 在 查找对应的节点的过程, 根据对应的 信息查找到对应的节点的父节点
// 其实空间的复杂度可以优化一下的 O(1)
func findAncestorHelper(root *BinaryTree, node *BinaryTree, find []*BinaryTree, ancestor *[]*BinaryTree) {
	if root == nil {
		return
	}

	if root == node {
		*ancestor = find[:]
		return
	}
	// 使用
	find = append(find, root)

	// 遍历左子树
	if root.Left != nil {
		findAncestorHelper(root.Left, node, find, ancestor)
	}
	// 遍历右子树
	if root.Right != nil {
		findAncestorHelper(root.Right, node, find, ancestor)

	}
	// 将当前的节点出队列
	find = find[0 : len(find)-1]
}

func findBrotherNodes(root *BinaryTree, node *BinaryTree) (ans []*BinaryTree) {
	if root == nil || node == nil {
		return ans
	}

	find := make([]*BinaryTree, 0)
	ancestor := make([]*BinaryTree, 0)
	findAncestorHelper(root, node, find, &ancestor)

	if len(ancestor) == 0 {
		return ans
	}

	last := ancestor[len(ancestor)-1]
	if last.Left != nil {
		ans = append(ans, last.Left)
	}

	if last.Right != nil {
		ans = append(ans, last.Right)
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
	/**
		         -1
		        / \
		       0        // a
		     /
		    1           // b
		   /
	      2            // c
		 /
		3           // d
		/     //
		4    // e
		/ \
		5  6  // f , g



	*/

	ans := findBrotherNodes(&root, &f)
	expected := []*BinaryTree{&f, &g}
	fmt.Println(ans, reflect.DeepEqual(ans, expected)) // output two node

}
