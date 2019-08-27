package main

import (
	"fmt"
)

func arr(x int) {
	defer func() {
		if rev := recover(); rev != nil {
			fmt.Println(rev)
		}
	}()
	var a [10]int
	a[x] = 123
}

func main() {
	arr(11)
	fmt.Println("recovery test")
}
