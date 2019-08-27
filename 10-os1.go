package main

import (
	"fmt"
	"os"
)

func main() {

	list := os.Args
	l := len(list)
	fmt.Println("input:", list, "\ninput len:", l)

	for i := range list {
		fmt.Printf("list[%d]:%s\n", i, list[i])
	}
}
