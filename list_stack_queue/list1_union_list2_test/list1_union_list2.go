package test

import "fmt"

// 实现两个有序链表的并集的操作

//
func main() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,21}
	l2 := []int{7, 8, 9, 10, 11, 12, 13}
	ans := List1UnionList2(l1, l2)
	fmt.Println(ans )
}

// 需要注意的是, 集合不能有重复的数字
// 如果两个链表中有相同的元素, 合并第一个, 其他的跳过
func List1UnionList2(l1, l2 [] int) (ans [] int) {
	ans = make([]int, 0)

	i := 0
	j := 0

	for i < len(l1) && j < len(l2) {
		// 如果集合中的两个元素的值相同的情况下, 合并一个进行了
		if l1[i] == l2[j] {
			ans = append(ans, l1[i])
			i++
			j++
			// 这里是做集合中元素的去重
			for i < len(l1) && l1[i] == l1[i-1] {
				i++
			}
 			//
			for j < len(l2) && l2[j] == l2[ j-1] {
				j++
			}

			continue
		}

		// 这里优先考虑的是小元素
		if l1[i] < l2[j] {
			ans = append(ans, l1[i])
			i++
			for i < len(l1) && l1[i] == l1[i-1] {
				i++
			}
			continue
		}

		if l1[i] > l2[j] {
			ans = append(ans, l2[j])
			j++
			for j < len(l2) && l2[j] == l2[ j-1] {
				j++
			}

			continue
		}
	}

	// 如果 链表中的元素的 l2 大于 l1 这里也是需要合并进去
	// 这里同样也是需要排序相同的元素
	for j < len(l2) {
		ans = append(ans, l2[j])
		j++
		for j < len(l2) && l2[j] == l2[ j-1] {
			j++
		}

		continue

	}

	for i < len(l1) {
		ans = append(ans, l1[i])
		i++
		for i < len(l1) && l1[i] == l1[i-1] {
			i++
		}
		continue
	}
	return ans
}
