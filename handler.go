package restapi

import (
	"net/http"
)

type Handler struct {
	res http.ResponseWriter
	req *http.Request
}

func (self *Handler) Call(apiName, method string, params Params) Output {
	return Call(apiName, method, params)
}

func (self *Handler) innerCall(api IApi, method string, params Params) Output {
	return innerCall(api, method, params)
}
