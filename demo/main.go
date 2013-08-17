package main

import (
	"fmt"
	"github.com/elsonwu/restapi"
	"net/url"
)

type User struct {
	restapi.Api
}

func (self *User) GET(params restapi.Params) restapi.Output {

	url := url.URL{}
	query := url.Query()
	query.Add("with_content_user", "1")
	query.Add("with_content_comments", "1")
	params2 := restapi.Params{Query: query}
	return restapi.Call("content", restapi.MethodGet, params2)
}

func (self *User) POST(params restapi.Params) restapi.Output {
	return restapi.SetupOutput(true, restapi.Map{"user": "post", "params": params}, nil)
}

type Content struct {
	restapi.Api
}

func (self *Content) GET(params restapi.Params) restapi.Output {
	return restapi.SetupOutput(true, restapi.Map{"content": "get", "params": params}, nil)
}

func (self *Content) POST(params restapi.Params) restapi.Output {
	return restapi.SetupOutput(true, restapi.Map{"content": "post", "params": params}, nil)
}

func main() {
	fmt.Println("binding:8888")
	restapi.Add("user", &User{})
	restapi.Add("content", &Content{})
	restapi.Run(":8888")
}
