package restapi

import (
	"net/http"
	"reflect"
)

type IHandler interface {
	Get(apiName string) IApi
	Call(apiName, method string, params Params) IOutput
	Req() *http.Request
}

type Handler struct {
	req *http.Request
}

func (self *Handler) Req() *http.Request {
	return self.req
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
		if MethodGet == method {
			return api.GET(params)
		} else if MethodPost == method {
			return api.POST(params)
		} else if MethodPut == method {
			return api.PUT(params)
		} else if MethodDelete == method {
			return api.DELETE(params)
		}
	}

	return Output(false, nil, []string{"API(" + apiName + ") does not exist"})
}
