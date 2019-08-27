package main

import (
	"fmt"
)

func fibo(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("end")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	//从chan中读取数据
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true
	}()

	//产生数据，写入chan
	fibo(ch, quit)
}
