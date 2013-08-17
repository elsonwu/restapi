package main

import (
	"fmt"
	"github.com/elsonwu/restapi"
	// "time"
)

type User struct {
	restapi.Api
}

func (self *User) GET(params restapi.Params) restapi.Output {
	// query := params.Query
	// query.Add("with_content_user", "1")
	// query.Add("with_content_comments", "1")
	// return restapi.Call("content", restapi.MethodGet, params, rs)
	fmt.Println("before:", self.Owner().Req())
	// time.Sleep(15 * time.Second)
	fmt.Println("after:", self.Owner().Req())
	return restapi.SetupOutput(true, restapi.Map{"user": "get", "params": params}, nil, 0)
}

func (self *User) POST(params restapi.Params) restapi.Output {
	return restapi.SetupOutput(true, restapi.Map{"user": "post", "params": params}, nil, 0)
}
