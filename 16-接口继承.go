package main

import (
	"fmt"
	"reflect"
)

type Humaner interface {
	sayhi()
}

type Woman interface {
	Humaner
	sing(lrc string)
}

type Student struct {
	name string
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student %s say hi\n", tmp.name)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("sing", lrc)
}

func main() {
	var i Woman
	s := &Student{"Bran"}
	i = s
	// v := Student{"Tom"}
	var v interface{} = Student{"Tom"}
	i.sayhi() //来自继承的Humaner
	i.sing("lalala")

	fmt.Println("type:", reflect.TypeOf(i))
	fmt.Println("value:", reflect.ValueOf(i))
	fmt.Println("type:", reflect.TypeOf(v))
	fmt.Println("value:", reflect.ValueOf(v))
}
