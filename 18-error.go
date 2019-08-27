package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("分母为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	//var err1 error
	// err1 := fmt.Errorf("%s", "this is err1")
	// fmt.Println("err1: ", err1)

	// err2 := errors.New("this is err2")
	// fmt.Println("err2: ", err2)

	result, err := MyDiv(5, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
