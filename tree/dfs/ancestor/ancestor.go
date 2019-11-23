// 在做前端开发的过程中, 有个 api 叫做或者指定类名结点的祖先的 api
// 那么这个 api 到底是怎么实现的呢?
// 我们来看一看 基于二叉树的祖先结点的查找吧
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

// 查找目标节点的祖先节点
// 解决思路是:
// 前序遍历树的时候, 将树的节点放入一个队列, 后序遍历的时候, 将树的节点出队列
// 如果找到需要查找的节点, 将节点出队列
// 如果子树为空, 推出递归
func LeastCommonAncestors(root *BinaryTree, find *BinaryTree) (ancestor []*BinaryTree) {
	if root == nil {
		return nil
	}

	vec := make([]*BinaryTree, 0)
	// 查找节点 find 的祖先节点
	helperAncestor(root, find, vec, &ancestor)
	return ancestor
}

// DFS 遍历二叉树, 前序遍历是将树的节点入队列, 后序遍历时将树的节点把当前的节点出队列
func helperAncestor(root *BinaryTree, node *BinaryTree, vec []*BinaryTree, find *[]*BinaryTree) {
	if root == nil {
		return
	}

	if root == node {
		// 如果找到需要查找的节点, 将当前的公共前缀祖先赋值给 find
		*find = vec[:]
		return
	}

	vec = append(vec, root)
	// 递归遍历左子树
	helperAncestor(root.Left, node, vec, find)
	// 递归遍历右子树
	helperAncestor(root.Right, node, vec, find)
	// 将当前的头节点移除从栈顶弹出去
	vec = vec[0 : len(vec)-1]

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
	// f 节点的祖先为 : root, a, b, c, d, e
	ancestors := LeastCommonAncestors(&root, &f)
	expected := []*BinaryTree{&root, &a, &b, &c, &d, &e}
	fmt.Println(ancestors, reflect.DeepEqual(ancestors, expected))
}
