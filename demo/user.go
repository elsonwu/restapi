package main

import (
	"fmt"
	"github.com/elsonwu/restapi"
	// "net/http"
	// "time"
)

type User struct {
	restapi.Api
}

func (self *User) View(params restapi.Params) restapi.IOutput {
	fmt.Println("id:" + params.Query.Get("id"))
	api := self.Owner().Get("content")
	return api.View(params)
}

func (self *User) Create(params restapi.Params) restapi.IOutput {
	query := params.Query
	query.Add("with_content_user", "1")
	query.Add("with_content_comments", "1")
	return self.Owner().Call("content", restapi.MethodView, params)
}

func (self *User) List(params restapi.Params) restapi.IOutput {
	return restapi.Output(true, restapi.Map{"xxx": "yyy"}, nil)
}
