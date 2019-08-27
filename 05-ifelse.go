package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Please input a num:")
	fmt.Scan(&a)

	if a > 42 {
		fmt.Printf("%d > 42", a)
	} else if a < 42 {
		fmt.Printf("%d < 42", a)
	} else if a == 42 {
		fmt.Printf("%d = 42", a)
	} else {
		fmt.Println("不可能！")
	}

	// switch 42 {
	switch {
	case a > 90:
		fmt.Println("\n优秀")
		// break  默认包含，可不写
	case a > 80:
		fmt.Println("\n良好")
		break
	default:
		fmt.Println("\n垃圾")
	}
}
