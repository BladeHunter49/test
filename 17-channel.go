package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string) {
	for _, data := range str { //一个一个字符打印
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Print("\n")
}

func person1() {
	Printer("hello")
	ch <- 123 //给管道写数据
}

func person2() {
	<-ch //管道没数据会阻塞
	Printer("world")
}

func main() {
	go person1()
	go person2()

	for {
	}
}
