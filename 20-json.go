package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	//struct创建JSON
	s := IT{"itcast", []string{"go", "python"}, true, 666.6}
	// buf, err := json.Marshal(s)
	bufs, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("buf=", string(bufs))

	//map创建JSON
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"go", "python"}
	m["isok"] = true
	m["price"] = 666.6
	bufm, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("buf=", string(bufm))

	jsonbuf := ` {"Company": "itcast",
	"Subjects": ["go","python"],
	"Isok": true,
 	"Price": 666.6}`
	//JSON解析到struct,取某个字段另定义一个IT2只包含所要字段使用即可！优先使用
	var jxs IT
	err1 := json.Unmarshal([]byte(jsonbuf), &jxs) //地址传递！
	if err1 != nil {
		fmt.Println("err=", err1)
	}
	fmt.Printf("jxs=%+v\n", jxs)

	//JSON解析到map，取某个字段需要空接口类型断言，类型是万能指针，空切片
	jxm := make(map[string]interface{}, 4)
	err2 := json.Unmarshal([]byte(jsonbuf), &jxm)
	if err2 != nil {
		fmt.Println("err=", err2)
	}
	fmt.Printf("jxm=%+v\n", jxm)

	var str string
	for _, value := range jxm {
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("类型为string，值为%v\n", str)
		case bool:
			bo := data
			fmt.Printf("类型为bool，值为%v\n", bo)
		}
	}
}
