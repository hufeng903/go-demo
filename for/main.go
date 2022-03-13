package main

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
}
