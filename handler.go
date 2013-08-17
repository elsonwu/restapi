package restapi

import (
	"net/http"
)

type Handler struct {
	res http.ResponseWriter
	req *http.Request
}

func (self *Handler) Call(apiName, method string, params Params, rs IRuntimes) Output {
	return Call(apiName, method, params, rs)
}

func (self *Handler) innerCall(api IApi, method string, params Params, rs IRuntimes) Output {
	return innerCall(api, method, params, rs)
}
