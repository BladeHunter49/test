package main

import (
	"fmt"
)

func main() {
	a := [9]int{6, 4, -3, 5, -2, -1, 0, 1, -9}
	fmt.Println("before:", a)
	low := 0
	mid := 0
	heigh := len(a) - 1

	for mid <= heigh {
		cur := a[mid]
		if cur < 0 {
			a[mid], a[low] = a[low], a[mid]
			mid++
			low++
		} else if cur == 0 {
			mid++
		} else {
			a[mid], a[heigh] = a[heigh], a[mid]
			heigh--
		}
	}

	fmt.Println("after:", a)
}
