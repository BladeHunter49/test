package main

import (
	"fmt"
)

func Add01(a, b int) int {
	return a + b
}

//给一个接受者rev绑定一个函数
type long int

func (rev long) Add02(oth long) long {
	return rev + oth
}

func main() {
	var result01 int
	result01 = Add01(1, 2)
	fmt.Println("result01=", result01)

	//	var result02 long   int不行，类型判定严格
	var a long = 3
	result02 := a.Add02(4)
	fmt.Println("result02=", result02)
}
