package main

import "fmt"

// 树的广度遍历, 按层次输出二叉树
// 解决的过程是, 在按层次输出二叉树的时候, 记录当前遍历节点所在的层次
// 如果当前的层次
type binaryTreeLevel struct {
	left  *binaryTreeLevel
	right *binaryTreeLevel
	data  int
}

type helperBfsLevel struct {
	node  *binaryTreeLevel // 树节点
	level int              // 记录 node 节点所在的层次
}

func bfsLevel(root *binaryTreeLevel) (ans [][]int) {

	if root == nil {
		return
	}

	helper := make([]helperBfsLevel, 0)
	helper = append(helper, helperBfsLevel{
		node:  root,
		level: 0,
	})
	ans = append(ans, make([]int, 0))
	front := 0

	for len(helper) != 0 {
		node := helper[0]
		helper = helper[1:]
		// 如果当前节点所在的层次大于当前节点的层次, 换行
		if node.level > front {
			ans = append(ans, make([]int, 0))

		}
		ans[node.level] = append(ans[node.level], node.node.data)
		// 更新层次计数
		front = node.level

		// 左节点入队列
		if node.node.left != nil {
			helper = append(helper, helperBfsLevel{
				node:  node.node.left,
				level: node.level + 1,
			})
		}

		// 右节点入队列
		if node.node.right != nil {
			helper = append(helper, helperBfsLevel{
				node:  node.node.right,
				level: node.level + 1,
			})
		}
	}

	// 返回按层次遍历的结果
	return ans
}
func main() {
	a := binaryTreeLevel{data: 1}
	b := binaryTreeLevel{data: 2}
	c := binaryTreeLevel{data: 3}
	d := binaryTreeLevel{data: 4}

	root := binaryTreeLevel{data: 0}
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
	// :output [[0], [1 ,3], [2, 4]]
	*/

	ans := bfsLevel(&root)
	fmt.Println(ans)
}
