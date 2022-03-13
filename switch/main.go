package main

import "fmt"

func main() {
	var lang string

	lang = "js"

	//常用switch, 一个case 多个条件, 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case
	switch lang {
	case "java":
		println("this case is java")
	case "go", "js":
		println("this case is go")
		println("this case also is js")
	case "php":
		println("this case is php")
	default:
		println("this case is python")
	}

	//输出结果
	//this case is go
	//this case also is js

	//fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	switch lang {
	case "java":
		println("this case is java")
	case "go", "js":
		println("this case is go")
		println("this case also is js")
		fallthrough
	case "php":
		println("this case is php")
	default:
		println("this case is python")
	}

	//输出结果
	//this case is go
	//this case also is js
	//this case is php

	//类型判断
	var x interface{}

	switch i := x.(type) {
	case int:
		fmt.Printf("x type is %T", i)
	case string:
		fmt.Printf("x type is %T", i)
	case bool, nil:
		fmt.Printf("x type is %T", i)
	}

	//输出结果
	//x type is <nil>
}
