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
	View(Params) IOutput
	List(Params) IOutput
	Create(Params) IOutput
	Update(Params) IOutput
	Delete(Params) IOutput
	UpdateAll(Params) IOutput
	DeleteAll(Params) IOutput
	SetOwner(IHandler)
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

func (self *Api) View(params Params) IOutput {
	return nil
}

func (self *Api) Create(params Params) IOutput {
	return nil
}

func (self *Api) Update(params Params) IOutput {
	return nil
}

func (self *Api) Delete(params Params) IOutput {
	return nil
}

func (self *Api) List(params Params) IOutput {
	return nil
}

func (self *Api) UpdateAll(params Params) IOutput {
	return nil
}

func (self *Api) DeleteAll(params Params) IOutput {
	return nil
}
