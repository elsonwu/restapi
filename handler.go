package restapi

import (
	"net/http"
	"reflect"
)

type Handler struct {
	res http.ResponseWriter
	req *http.Request
}

func (self *Handler) Res() http.ResponseWriter {
	return self.res
}

func (self *Handler) Req() *http.Request {
	return self.req
}

func (self *Handler) Call(apiName, method string, params Params) Output {
	api := _apis.Get(apiName)
	if v, ok := api.(IApi); ok {
		tempApi := reflect.New(reflect.ValueOf(v).Elem().Type()).Interface()
		if v, ok := tempApi.(IApi); ok {
			v.SetOwner(self)
			return self.innerCall(v, method, params)
		}
	}

	return SetupOutput(false, nil, []string{"api - " + apiName + " does not exist"})
}

func (self *Handler) innerCall(api IApi, method string, params Params) Output {
	var output Output
	if MethodGet == method {
		output = api.GET(params)
	} else if MethodPost == method {
		output = api.POST(params)
	} else if MethodPut == method {
		output = api.PUT(params)
	} else if MethodDelete == method {
		output = api.DELETE(params)
	}

	return output
}
