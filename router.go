package restapi

import (
	"net/http"
	"strings"
)

type RouterFunc func(ctx IContext, res http.ResponseWriter, req *http.Request) (apiName, method string, ok bool)

func DefaultRouterFunc(ctx IContext, res http.ResponseWriter, req *http.Request) (apiName, method string, ok bool) {
	paths := strings.Split(req.URL.Path, "/")
	lenPaths := len(paths)
	if 1 == lenPaths {
		http.Error(res, "API Not found", 404)
		return
	}

	if 2 < lenPaths && ("GET" == req.Method || "PUT" == req.Method || "DELETE" == req.Method) {
		if "" != paths[2] {
			ctx.Query().Add("id", paths[2])
		}
	}

	method = requestApiMethod(req, ctx)
	if "" == method {
		http.Error(res, "Request does not acceptable", 400)
		return
	}

	apiName = paths[1]
	ok = true
	return
}
