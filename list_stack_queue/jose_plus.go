package main

import "fmt"

func main() {
	// 约瑟夫问题, 使用的是了 list 的解法
	data := []int{1, 2, 3, 4, 5}
	evil := 1
	fmt.Println(JosePlus(data, evil))
}

func JosePlus(data []int, evil int) int {
	if len(data) == 1 {
		return data[0]
	}
	for len(data) != 1 {
		index := 0
		index++
		index = index % len(data)
		if index%evil == 0 {
			fmt.Println("cls: ", data[index])
			data = append(data[0:index], data[index+1:]...)
		}

	}
	return data[0]
}
