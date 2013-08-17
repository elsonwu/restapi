package main

import (
	"github.com/elsonwu/restapi"
)

type Content struct {
	restapi.Api
}

func (self *Content) GET(params restapi.Params) restapi.Output {
	return restapi.SetupOutput(true, restapi.Map{"content": "get", "params": params}, nil)
}

func (self *Content) POST(params restapi.Params) restapi.Output {
	return restapi.SetupOutput(true, restapi.Map{"content": "post", "params": params}, nil)
}
