package main

import (
	"demo/package/p1"
	. "demo/package1/p1"
	"fmt"
)

func init()  {
	fmt.Println("this is main package init")
}

func main() {
	Test4()

	p1.Test1()
}
