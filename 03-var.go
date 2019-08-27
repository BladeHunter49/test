package main

import (
	"fmt"
)

func main() {
	a := 10
	const b int = 5
	a += 1
	fmt.Printf("a，b:%v ,%T \n", a, b)
	fmt.Println("b:", b)

	var c1 byte = 'a'
	var s1 string = "abc"
	c2 := 'b'
	s2 := "nova"
	fmt.Println("len(s1)", len(s1))
	fmt.Println("len(s2)", len(s2))
	fmt.Println("s1:", s1, "\ns2:", s2)
	fmt.Printf("c1:%d\nc2:%c", c1, c2)
}
