package main

import (
	// "fmt"
	"github.com/elsonwu/restapi"
	// "time"
)

type User struct {
	restapi.Api
}

func (self *User) GET(params restapi.Params) restapi.IOutput {
	query := params.Query
	query.Add("with_content_user", "1")
	query.Add("with_content_comments", "1")
	return self.Owner().Call("content", restapi.MethodGet, params)
}

func (self *User) POST(params restapi.Params) restapi.IOutput {
	return restapi.SetupOutput(true, restapi.Map{"user": "post", "params": params}, nil, 0)
}
