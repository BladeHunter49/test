package main

import (
	"fmt"
)

func main() {
	//for
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println("sum 1~100:", sum)

	//range
	str := "Rock"
	for i := range str {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	for j, chr := range str {
		fmt.Printf("str[%d]=%c\n", j, chr)
		// fmt.Println("j:", j, "chr", byte(chr))  Println打印字符问题
	}

}
