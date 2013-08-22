package main

import (
	// "fmt"
	"github.com/elsonwu/restapi"
)

type Content struct {
	restapi.Api
}

func (self *Content) View(ctx restapi.IContext) restapi.IOutput {
	return restapi.Output(true, ctx, nil)
}

func (self *Content) Create(ctx restapi.IContext) restapi.IOutput {
	return restapi.Output(true, restapi.Map{"content": "post", "ctx": ctx}, nil)
}
