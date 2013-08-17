package main

import (
	// "fmt"
	"github.com/elsonwu/restapi"
)

func main() {
	restapi.Add("user", &User{})
	restapi.Add("content", &Content{})
	restapi.Run(":8888")
}
