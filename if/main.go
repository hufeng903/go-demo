package main

func main() {
	var i int
	i = 100

	//if-else普通的通用方式
	if i < 10 {
		println("i less than 10")
	} else {
		println("i more than 10")
	}

	//if-elseif-else 通用方式
	if i < 100 {
		println("i less than 100")
	} else if i > 50 {
		println("i more than 50")
	} else {
		println("i more than 50 and less than 100")
	}

	//初始化赋值，少写几行代码
	if a := 120 - i; a < 100 {
		println("i less than 20")
	} else {
		println("i more than 20")
	}
}
