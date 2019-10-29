package main

import (
	"bytes"
	"fmt"
)

// 逆波兰表达式的思想就是, 设置一个栈粗放操作数, 从左到右扫描,
// 每读到一个操作数就将其进栈, 没读到一个运算符, 就从栈顶取出两个操作数完成运算符代表的运算操作, 并把操作的结果作为一个新的操作数重新入栈
// 此过程一直到操作到后缀表达式读完, 最后栈顶的操作数就是该后缀表达式的运算结果
func main() {

	polish := PreReversePolish([]byte("1+(3-2/1)*2"))
	fmt.Println(polish)
	fmt.Println(ReversePolish(polish)) // example = 3
}

// 这里是将中缀表达式转后缀的表达式
func PreReversePolish(data []byte) (str string) {
	vec := make([]byte, 0, len(data))
	buff := bytes.NewBuffer(nil)
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			vec = append(vec, '(')
			continue
		}
		if data[i] == ')' {
			for da := vec[len(vec)-1]; len(vec) > 0; da = vec[len(vec)-1] {
				vec = vec[0 : len(vec)-1]
				if da == '(' {
					break
				}
				buff.WriteByte(da)
			}
			continue
		}

		switch data[i] {
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			vec = append(vec, data[i])
			break
		default:
			buff.WriteByte(data[i])
			break
		}
	}

	for len(vec) > 0 {
		buff.WriteByte(vec[len(vec)-1])
		vec = vec[0 : len(vec)-1]
	}

	return buff.String()
}

// 逆波兰表达式
func ReversePolish(str string) int {
	vec := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		switch str[i] {
		// 如果当前的符号操作的符号, 直接 进行对应的操作
		case '+':
			a, b := vec[len(vec)-2], vec[len(vec)-1]
			a = a + b
			vec = vec[0 : len(vec)-2]
			vec = append(vec, a)
		case '-':
			a, b := vec[len(vec)-2], vec[len(vec)-1]
			a = a - b
			vec = vec[0 : len(vec)-2]
			vec = append(vec, a)
		case '*':
			a, b := vec[len(vec)-2], vec[len(vec)-1]
			a = a * b
			vec = vec[0 : len(vec)-2]
			vec = append(vec, a)
		case '/':
			a, b := vec[len(vec)-2], vec[len(vec)-1]
			a = a / b
			vec = vec[0 : len(vec)-2]
			vec = append(vec, a)

		default:
			// 这是将 ascii 转为对应的数值
			vec = append(vec, str[i]-'0')
		}
	}
	// 栈顶元素就是最终求得的结果
	return int(vec[0])
}
