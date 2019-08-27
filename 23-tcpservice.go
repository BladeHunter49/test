package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err =", err)
			continue
		}

		buf := make([]byte, 1024)
		for {
			n, err1 := conn.Read(buf)
			if err1 != nil {
				fmt.Println("err1 =", err1)
				continue
			}
			fmt.Println("buf =", string(buf[:n]))
		}
		defer conn.Close()
	}
}
