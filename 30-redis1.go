package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect err:", err)
		return //结束程序，执行defer
	}
	defer con.Close()

	_, err = con.Do("SETEX", "gokey", "5", "test")
	if err != nil {
		fmt.Println("SET error", err)
	}

	s, err := redis.String(con.Do("GET", "gokey"))
	if err != nil {
		fmt.Println("GET err", err)
	} else {
		fmt.Println("GET value", s)
	}

	time.Sleep(3 * time.Second)
	key_exit, err := redis.Bool(con.Do("EXISTS", "gokey"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", key_exit)
	}

	con.Send("SET", "foo", "bar")
	con.Send("GET", "foo")
	con.Flush()
	r1, _ := redis.String(con.Receive()) //一次只能接收一个命令的return
	r2, _ := redis.String(con.Receive())
	fmt.Println("GET value", r1, r2)
}
