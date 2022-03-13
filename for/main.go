package main

import (
	"fmt"
)

func main() {
	//普通循环
	var i int
	for i = 0; i < 10; i++ {
		println(i)
	}

	//初始化，并循环
	for j := 0; j < 10; j++ {
		println(j)
	}

	//省略初始化
	var h int32 = 10

	for ; h < 10; h++ {
		println(h)
	}

	//省略后置子句, 死循环
	for h < 10 {
		println(h)
	}

	//省略初始化和后置子句，类似if
	for h == 10 {
		println(h)
	}

	//循环数组
	arr := [4]int{1, 2, 3, 4}

	for index, val := range arr {
		fmt.Printf("%d => %d \n", index, val)
	}

	//循环切片
	sliceTest := []int{3, 4, 5, 6}
	sliceTest = append(sliceTest, 5)
	for k, v := range sliceTest {
		fmt.Printf("%d => %d", k, v)
	}

	//循环map
	//var mapTest map[string]string
	mapTest := make(map[string]string)
	mapTest["a"] = "A"
	mapTest["b"] = "B"
	mapTest["c"] = "C"
	mapTest["d"] = "D"

	delete(mapTest, "d")

	for k, v := range mapTest {
		fmt.Printf("%s => %s \n", k, v)
	}

}
