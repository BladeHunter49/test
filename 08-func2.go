package main

import (
	"fmt"
)

//回调函数：函数参数有一个是函数类型！多态实现
//调用同一个接口，不同的表现，可以先定义再补充实现功能(拓展优势)
type Functype func(int, int) int

func Calc(a, b int, fTest Functype) (result int) {
	result = fTest(a, b)
	return
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func main() {
	r := Calc(3, 5, Add)
	fmt.Println("r1 =", r)

	r = Calc(1, 3, Minus)
	fmt.Println("r2 =", r)

	r = Calc(2, 3, Mul)
	fmt.Println("r3 =", r)
}
