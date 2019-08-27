package main

import (
	"flag"
	"fmt"
)

//自定义解析参数，实现Set和String方法,类比linux命令
type Hello string

func (p *Hello) Set(s string) error {
	v := fmt.Sprintf("hello %v", s)
	*p = Hello(v)
	return nil
}

func (p *Hello) String() string {
	return fmt.Sprintf("%f", *p)
}

// 第一个参数 ：接收flagname的实际值的 有无&
// 第二个参数 ：flag名称为flagname
// 第三个参数 ：flagname默认值为1234
// 第四个参数 ：flagname的提示信息
func main() {
	nameP := flag.String("name", "username", "姓名")
	ageP := flag.Int("age", 18, "年龄")
	sexP := flag.Bool("set", true, "man")

	var email string
	flag.StringVar(&email, "email", "defult@qq.com", "邮箱")

	var hello Hello
	flag.Var(&hello, "hello", "hello参数")

	flag.Parse()
	others := flag.Args()
	fmt.Println("name:", *nameP)
	fmt.Println("age:", *ageP)
	fmt.Println("sex:", *sexP)
	fmt.Println("hello:", hello)
	fmt.Println("email:", email)
	fmt.Println("other:", others)
}
