package main

import (
	"demo/package/p1"
	. "demo/package1/p1"
	"errors"
	"fmt"
)

func init() {
	fmt.Println("this is main package init")
}

func main() {
	Test4()

	p1.Test1()
}

func div(a int, b int) (ret int, err error) {
	if a == 0 {
		err = errors.New("this is error")
	}

	ret = a + b

	return
}
