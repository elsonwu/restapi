package main

import (
	// "fmt"
	"github.com/elsonwu/restapi"
	"net/http"
)

func main() {
	restapi.Add("user", &User{})
	restapi.Add("content", &Content{})
	//customized responseFunc
	restapi.Conf.ResponseFunc = func(output restapi.IOutput, ctx restapi.IContext, res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello elson"))
	}
	restapi.Run(":8888")
}
