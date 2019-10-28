package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 给定一个表 L 和另一个 表 P, 它们包含以上升的序列的整数, 操作 printLots(L, P) 将
// 打印 L 中那些由 P 所指定的位置上的元素,
// 例如: 如果 P = 1, 3, 4, 6
// 那么, L 中位于第一, 第三, 第四, 和第六 位置上的元素被打印出来
func main() {
	const (
		LEN = 10
	)
	rand.Seed(time.Now().UnixNano())
	l:=make([]int, LEN, LEN )
	p := make([]int , 3 )
	for i:= 0 ; i< LEN; i++ {
		l[i] =rand.Intn(LEN +100)

	}
	sort.Ints(l)
	p[0] = 1
	p[1] = 0
	p[2] = 3
	fmt.Println("l:",l)
	fmt.Println("p:", p )


	printLots(l, p )
}


func printLots(l, p []int ) {
	if len(p) <=0 {
		return
	}

	for _, val:=range p {
		// 如果 l 的长度小于 val, 直接退出就行了
		if len(l) < val  {
			return
		}
		if val <= 0 {
			continue
		}
		fmt.Print(l[val-1], " ")
	}
}