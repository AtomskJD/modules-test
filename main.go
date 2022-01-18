package main

import (
	"fmt"
	"github.com/atomskjd/modules-test/bye"
	"github.com/atomskjd/modules-test/hello"
)

func main() {
	fmt.Println("Simple packeges test")
	fmt.Println(hello.En())
	fmt.Println(bye.En())
	fmt.Println(hello.Gr())
	fmt.Println(bye.Gr())
	fmt.Println(hello.Jp())
	fmt.Println(bye.Jp())
}
