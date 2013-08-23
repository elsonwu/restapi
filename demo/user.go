package main

import (
	"errors"
	"github.com/elsonwu/restapi"
)

type User struct {
	restapi.Api
}

func (self *User) BeforeRun(ctx restapi.IContext) error {
	if "" == ctx.Query().Get("key") {
		return errors.New("key is missing")
	}

	return nil
}

func (self *User) View(ctx restapi.IContext) restapi.IOutput {
	return restapi.Output(true, "user view data", nil)
}
