package main

import (
	. "fmt"
)

func test01(a int) (sum int) {
	for i := 1; i <= a; i++ {
		sum += i
	}
	return
}
func test02(i int) int {
	if i == 1 {
		return 1
	}
	return i + test02(i-1)
}

func main() {
	var sum1, sum2 int

	sum1 = test01(99)
	Println("sum1:", sum1)

	sum2 = test02(100)
	Println("sum2:", sum2)
}
