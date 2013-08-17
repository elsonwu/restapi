package restapi

import (
	"net/http"
	"reflect"
)

type Handler struct {
	req *http.Request
}

func (self *Handler) Req() *http.Request {
	return self.req
}

func (self *Handler) Call(apiName, method string, params Params) Output {
	api := _apis.Get(apiName)

	if nil != api {
		//create a new api to handle this call.
		tempApi := reflect.New(reflect.ValueOf(api).Elem().Type()).Interface()
		if api, ok := tempApi.(IApi); ok {
			api.SetOwner(self)

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
	}

	return SetupOutput(false, nil, []string{"API(" + apiName + ") does not exist"}, 0)
}
