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
	if api, ok := self[apiName]; ok {
		return api
	}

	return nil
}

type IApi interface {
	View(IContext) IOutput
	List(IContext) IOutput
	Create(IContext) IOutput
	Update(IContext) IOutput
	Delete(IContext) IOutput
	UpdateAll(IContext) IOutput
	DeleteAll(IContext) IOutput

	BeforeRun(IContext) error
	BeforeView(IContext) error
	BeforeList(IContext) error
	BeforeCreate(IContext) error
	BeforeUpdate(IContext) error
	BeforeDelete(IContext) error
	BeforeUpdateAll(IContext) error
	BeforeDeleteAll(IContext) error
}

type Api struct{}

func (self *Api) View(ctx IContext) IOutput {
	return nil
}

func (self *Api) Create(ctx IContext) IOutput {
	return nil
}

func (self *Api) Update(ctx IContext) IOutput {
	return nil
}

func (self *Api) Delete(ctx IContext) IOutput {
	return nil
}

func (self *Api) List(ctx IContext) IOutput {
	return nil
}

func (self *Api) UpdateAll(ctx IContext) IOutput {
	return nil
}

func (self *Api) DeleteAll(ctx IContext) IOutput {
	return nil
}

func (self *Api) BeforeRun(IContext) error {
	return nil
}

func (self *Api) BeforeView(IContext) error {
	return nil
}

func (self *Api) BeforeList(IContext) error {
	return nil
}

func (self *Api) BeforeCreate(IContext) error {
	return nil
}

func (self *Api) BeforeUpdate(IContext) error {
	return nil
}

func (self *Api) BeforeDelete(IContext) error {
	return nil
}

func (self *Api) BeforeUpdateAll(IContext) error {
	return nil
}

func (self *Api) BeforeDeleteAll(IContext) error {
	return nil
}
