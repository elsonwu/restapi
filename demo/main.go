package main

import (
	"fmt"
	"github.com/elsonwu/restapi"
)

func main() {
	fmt.Println("binding:8888")
	restapi.Add("user", &User{})
	restapi.Add("content", &Content{})
	restapi.Run(":8888")
}
