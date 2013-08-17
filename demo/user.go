package main

import (
	"github.com/elsonwu/restapi"
)

type User struct {
	restapi.Api
}

func (self *User) GET(params restapi.Params, rs restapi.IRuntimes) restapi.Output {
	query := params.Query
	query.Add("with_content_user", "1")
	query.Add("with_content_comments", "1")
	return restapi.Call("content", restapi.MethodGet, params, rs)
}

func (self *User) POST(params restapi.Params, rs restapi.IRuntimes) restapi.Output {
	return restapi.SetupOutput(true, restapi.Map{"user": "post", "params": params}, nil)
}
