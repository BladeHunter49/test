package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("input a,b:")
	fmt.Scan(&a)
	fmt.Scanf("%d", &b)
	a, b = b, a
	fmt.Printf("after switch a,b:%d,%d", a, b)

}
