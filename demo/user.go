package main

import (
	// "fmt"
	"github.com/elsonwu/restapi"
	// "net/http"
	// "time"
)

type User struct {
	restapi.Api
}

func (self *User) View(ctx restapi.IContext) restapi.IOutput {
	return restapi.Output(true, ctx.Query(), nil)
	// fmt.Println("id:" + ctx.Query().Get("id"))
	// api := restapi.Get("content")
	// return api.View(ctx)
}

func (self *User) Create(ctx restapi.IContext) restapi.IOutput {
	query := ctx.Query()
	query.Add("with_content_user", "1")
	query.Add("with_content_comments", "1")
	return restapi.Call("content", restapi.MethodView, ctx)
}

func (self *User) List(ctx restapi.IContext) restapi.IOutput {
	return restapi.Output(true, []restapi.Map{restapi.Map{"xxx": "yyy"}, restapi.Map{"eeee": "ssss"}}, nil)
}
