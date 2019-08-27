package main

import (
	"fmt"
)

var x int

func test03() func() int {
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := test03()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func main1() {
	a := 42

	func() {
		a += 1
		fmt.Println("a=", a)
	}()

	f1 := func() {
		a += 1
		fmt.Println("a=", a)
	}
	f1()

	m := func(i, j int) (max int) {
		if i >= j {
			max = i
		} else {
			max = j
		}
		return
	}(11, 12)
	fmt.Println("max=", m)

}
