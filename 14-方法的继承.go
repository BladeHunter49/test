package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

func (temp Person) Printinfo() {
	fmt.Println("temp = ", temp)
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s := Student{Person{"jon", 'm', 14}, 123, "42th"}
	s.Printinfo()
}
