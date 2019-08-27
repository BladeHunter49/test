package main

import (
	"fmt"
	"time"
)

func main() {
	//1
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("2s!")
	//2
	time.Sleep(2 * time.Second)
	fmt.Println("2s!!")
	//3  2s后产生一个事件，往channel中写内容
	<-time.After(2 * time.Second)
	fmt.Println("2s!!!")
}
