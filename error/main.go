package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("new error1")
	if err1 != nil {
		fmt.Println(err1)
	}
	err2 := fmt.Errorf("new error2") // 就是对errors.New的封装
	if err2 != nil {
		fmt.Println(err2)
	}
}
