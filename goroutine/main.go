package main

import (
	"fmt"
)

func main()  {
	println("this is starting")
	names := [] string {"a", "b", "c", "d"}
	for _, name := range names {
		go func(who string) {
			go fmt.Printf("this name is %s! \n", who)
		}(name)

		//time.Sleep(time.Second)
	}

	println("this is end")

	channel := make(chan int, 10)

	channel <- 1
	channel <- 2
	channel <- 3


}
