package restapi

import (
	"strings"
)

var _apis apis

func init() {
	_apis = make(apis)
}

type apis map[string]IApi

func (self apis) Add(apiName string, api IApi) {
	apiName = strings.ToLower(apiName)
	self[apiName] = api
}

func (self apis) Get(apiName string) IApi {
	apiName = strings.ToLower(apiName)
	api, ok := self[apiName]
	if ok {
		return api
	}

	return nil
}

type IApi interface {
	GET(Params) IOutput
	POST(Params) IOutput
	PUT(Params) IOutput
	DELETE(Params) IOutput
	SetOwner(owner IHandler)
	Owner() IHandler
}

type Api struct {
	owner IHandler
}

func (self *Api) SetOwner(owner IHandler) {
	self.owner = owner
}

func (self *Api) Owner() IHandler {
	return self.owner
}

func (self *Api) GET(params Params) IOutput {
	return SetupOutput(false, Map{}, []string{"Does not support GET method"}, 0)
}

func (self *Api) POST(params Params) IOutput {
	return SetupOutput(false, Map{}, []string{"Does not support POST method"}, 0)
}

func (self *Api) PUT(params Params) IOutput {
	return SetupOutput(false, Map{}, []string{"Does not support PUT method"}, 0)
}

func (self *Api) DELETE(params Params) IOutput {
	return SetupOutput(false, Map{}, []string{"Does not support DELETE method"}, 0)
}
