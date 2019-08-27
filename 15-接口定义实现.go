package main

import (
	"fmt"
)

type Humaner interface {
	sayhi()
}

type Student struct {
	name string
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student %s say hi\n", tmp.name)
}

type Teacher struct {
	name string
}

func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher %s say hi\n", tmp.name)
}

func Whosayhi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"Bran"}
	t := &Teacher{"Mr.W"}
	Whosayhi(s)
	Whosayhi(t)
	//多态
}

func main01() {
	var i Humaner
	s := &Student{"Bran"}
	i = s
	i.sayhi()

	t := &Teacher{"Mr.W"}
	i = t
	//重点，同一个i.sayhi()不同表现，接收者为指针
	i.sayhi()
}
