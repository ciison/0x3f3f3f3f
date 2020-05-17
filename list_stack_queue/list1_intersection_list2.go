package main

import (
	"fmt"
	"sort"
)

// 求两个列表的交集
func main() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 8, 8, 9, 10, 13, 16}
	l2 := []int{2, 3, 4, 5, 6, 7, 6, 7, 7, 7, 8}
	sort.Ints(l1)
	sort.Ints(l2)
	fmt.Println(List1IntersectionList2(l1, l2))
}

// 求两个有序列表的交集
// 1. 如果两个集合的元素相同, 将元素添加到新的列表中, 然后需要做后一个元素的值和前一个元素的不能相同
// 2. 如果某个列表的当前值小于另一个列表的当前值, 将列表的当前下标加一, 然后也需要做去重的处理, 因为没有意义循环浪费过多的判断
// 最后 ans 的结果就是当前两个列表的并集
func List1IntersectionList2(l1, l2 []int) (ans []int) {

	ans = make([]int, 0)

	for i, j := 0, 0; i < len(l1) && j < len(l2); {
		// 如果两个集合的元素相等, 将集合的元素添加到 ans
		if l1[i] == l2[j] {
			ans = append(ans, l1[i])
			i++
			j++

			// 下面是作去重的判断
			for i < len(l1) && l1[i] == l1[i-1] {
				i++
			}

			for j < len(l2) && l2[j] == l2[j-1] {
				j++
			}
			continue
		}

		// 如果 l1[i] 小于l2[j] 需要将 i 的值向后增加
		if l1[i] < l2[j] {
			i++
			// 这里需要做去重的处理
			for i < len(l1) && l1[i] == l1[i-1] {
				i++
			}
			continue
		}

		if l1[i] > l2[j] {
			j++
			for j < len(l2) && l2[j] == l2[j-1] {
				j++
			}
			continue
		}
	}
	return ans
}
