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

func (p *Person) Setinfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}
func main() {
	p1 := Person{"Wayne", 'm', 18}
	p1.Printinfo()

	var p2 Person
	p2.Setinfo("Jon", 'f', 14)
	p2.Printinfo()
}
