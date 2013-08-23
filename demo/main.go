package main

import (
	// "fmt"
	"github.com/elsonwu/restapi"
	//"net/http"
	"errors"
)

func main() {
	restapi.Add("user", &User{})
	restapi.Add("content", &Content{})
	restapi.Filter.On("beforeHandleRequest", func(ctx restapi.IContext) error {
		return errors.New("Sorry, stop")
	})
	//customized responseFunc
	// restapi.Conf.ResponseFunc = func(output restapi.IOutput, ctx restapi.IContext) {
	// 	ctx.Res().Write([]byte("hello elson"))
	// }

	// restapi.Conf.RouterFunc = func(ctx restapi.IContext) (apiName, method string, ok bool) {
	// 	return "content", "View", true
	// }
	restapi.Run(":8888")
}
