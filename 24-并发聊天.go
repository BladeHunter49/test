package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var Onlinemap map[string]Client
var message = make(chan string)

func Handleconn(conn net.Conn) {
	defer conn.Close()

	cliAddr := conn.RemoteAddr().String()
	cli := Client{make(chan string), cliAddr, cliAddr}
	Onlinemap[cliAddr] = cli
	message <- "[" + cliAddr + "]" + cli.Name + "==>login"

	isinTime := make(chan bool)
	var isQuit = make(chan bool)

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("conn.Read error", err)
				return
			}
			msg := string(buf[:n-1])
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("online list:\n"))
				for _, tmp := range Onlinemap {
					conn.Write([]byte(tmp.Addr + ":" + tmp.Name + "\n"))
				}
			} else if len(msg) > 7 && msg[:7] == "rename|" {
				newname := strings.Split(msg, "|")[1]
				cli.Name = newname
				Onlinemap[cliAddr] = cli
				message <- "[" + cliAddr + "]" + "rename sucess!"
			} else {
				message <- "[" + cliAddr + "]" + cli.Name + ":" + msg
			}

			isinTime <- true
		}
	}()

	go WriteMsgToClient(cli, conn)
	for {
		select {
		case <-isQuit:
			delete(Onlinemap, cliAddr) //del from map
			message <- "[" + cliAddr + "]" + cli.Name + "==>logout"
			return
		case <-isinTime:

		case <-time.After(60 * time.Second):
			delete(Onlinemap, cliAddr) //del from map
			message <- "[" + cliAddr + "]" + cli.Name + "==>timeout leave"
			return

		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func Delivermsg() {
	Onlinemap = make(map[string]Client)
	for {
		msg := <-message
		for _, clt := range Onlinemap {
			clt.C <- msg
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen error", err)
		return
	}
	defer listener.Close()

	go Delivermsg()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			continue
		}
		go Handleconn(conn)
	}
}
