package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang")

	//输出管道
	stdout0, err := cmd0.StdoutPipe()

	if err != nil {
		fmt.Printf("error, can not obtain the stdout pipe for command %s\n", err)
		return
	}

	//创建start
	if err := cmd0.Start(); err != nil {
		fmt.Printf("error, can not start cmd %s\n", err)
		return
	}

	output0 := make([]byte, 30)
	n,err := stdout0.Read(output0)

	if err != nil {
		fmt.Printf("error, can not read data from the pip %s \n", err)

		return
	}

	fmt.Printf("%s\n", output0[:n])
}
