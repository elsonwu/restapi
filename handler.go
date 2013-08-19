package restapi

import (
	"net/http"
	"reflect"
)

type IHandler interface {
	Get(apiName string) IApi
	Call(apiName, method string, params Params) IOutput
	Req() *http.Request
	Res() http.ResponseWriter
}

type Handler struct {
	req *http.Request
	res http.ResponseWriter
}

func (self *Handler) Req() *http.Request {
	return self.req
}

func (self *Handler) Res() http.ResponseWriter {
	return self.res
}

func (self *Handler) Get(apiName string) IApi {
	api := _apis.Get(apiName)

	if nil != api {
		//create a new api to handle this call.
		tempApi := reflect.New(reflect.ValueOf(api).Elem().Type()).Interface()
		if api, ok := tempApi.(IApi); ok {
			api.SetOwner(self)
			return api
		}
	}

	return nil
}

func (self *Handler) Call(apiName, method string, params Params) IOutput {

	api := self.Get(apiName)
	if nil != api {
		if MethodView == method {
			return api.View(params)
		} else if MethodCreate == method {
			return api.Create(params)
		} else if MethodUpdate == method {
			return api.Update(params)
		} else if MethodDelete == method {
			return api.Delete(params)
		} else if MethodList == method {
			return api.List(params)
		} else if MethodUpdateAll == method {
			return api.UpdateAll(params)
		} else if MethodDeleteAll == method {
			return api.DeleteAll(params)
		}
	}

	return nil
}
