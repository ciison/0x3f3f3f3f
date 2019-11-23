package dfs

import "fmt"

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	data  int
}

// 查找树的最近公共祖先
// 解决思路是:
// 前序遍历树的时候, 将树的节点放入一个队列, 后序遍历的时候, 将树的节点出队列
// 如果找到需要查找的节点, 将节点出队列
// 如果子树为空, 推出递归
func LeastCommonAncestors(root *BinaryTree, nodeleft, noderight *BinaryTree) *BinaryTree {
	if root == nil {
		return nil
	}
	findLeft := make([]*BinaryTree, 0)
	findRight := make([]*BinaryTree, 0)

	vec := make([]*BinaryTree, 0)
	// 查找节点 nodeleft 的公共前缀子树
	helperAncestor(root, nodeleft, vec, &findLeft)
	// 将当前 的vec 清空
	vec = make([]*BinaryTree, 0)
	// 查找节点 noderight 公共前缀子树
	helperAncestor(root, noderight, vec, &findRight)

	var ancestor *BinaryTree

	for i, j := 0, 0; i < len(findLeft) && j < len(findRight); i, j = i+1, j+1 {
		// 如果祖先相同, 赋值
		if findLeft[i] == findRight[j] {
			ancestor = findLeft[i]
		}
	}
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

	ancestors := LeastCommonAncestors(&root, &f, &g)
	fmt.Println(ancestors)
}
