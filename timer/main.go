package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5*time.Second)
	fmt.Printf("start current time %v \n", time.Now())

	expireTime := <- timer.C

	fmt.Printf("expire time %v \n", expireTime)

	fmt.Printf("end current time %v \n", time.Now())
}
