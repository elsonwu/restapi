package restapi

import (
// "net/http"
)

type IApi interface {
	GET(Params) Output
	POST(Params) Output
	PUT(Params) Output
	DELETE(Params) Output
	SetOwner(owner *Handler)
	Owner() *Handler
}

type Api struct {
	owner *Handler
}

func (self *Api) SetOwner(owner *Handler) {
	self.owner = owner
}

func (self *Api) Owner() *Handler {
	return self.owner
}

func (self *Api) GET(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support get method"})
}

func (self *Api) POST(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support post method"})
}

func (self *Api) PUT(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support put method"})
}

func (self *Api) DELETE(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support delete method"})
}
