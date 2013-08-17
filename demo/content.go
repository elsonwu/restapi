package main

import (
	"github.com/elsonwu/restapi"
)

type Content struct {
	restapi.Api
}

func (self *Content) GET(params restapi.Params) restapi.IOutput {
	return restapi.SetupOutput(true, restapi.Map{"content": "get", "params": params}, nil, 0)
}

func (self *Content) POST(params restapi.Params) restapi.IOutput {
	return restapi.SetupOutput(true, restapi.Map{"content": "post", "params": params}, nil, 0)
}
