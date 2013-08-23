package restapi

import (
	"net/http"
	"strings"
)

type RouterFunc func(ctx IContext) (apiName, method string, ok bool)

func DefaultRouterFunc(ctx IContext) (apiName, method string, ok bool) {
	paths := strings.Split(ctx.Req().URL.Path, "/")
	lenPaths := len(paths)
	if 1 == lenPaths {
		http.Error(ctx.Res(), "API Not found", 404)
		return
	}

	if 2 < lenPaths && ("GET" == ctx.Req().Method || "PUT" == ctx.Req().Method || "DELETE" == ctx.Req().Method) {
		if "" != paths[2] {
			ctx.Query().Add("id", paths[2])
		}
	}

	method = requestApiMethod(ctx)
	if "" == method {
		http.Error(ctx.Res(), "Request does not acceptable", 400)
		return
	}

	apiName = paths[1]
	ok = true
	return
}

func requestApiMethod(ctx IContext) string {
	var method string
	if "GET" == ctx.Req().Method {
		if "" == ctx.Query().Get("id") {
			method = MethodList
		} else {
			method = MethodView
		}
	} else if "POST" == ctx.Req().Method {
		method = MethodCreate
	} else if "PUT" == ctx.Req().Method {
		if "" == ctx.Query().Get("id") {
			method = MethodUpdateAll
		} else {
			method = MethodUpdate
		}
	} else if "DELETE" == ctx.Req().Method {
		if "" == ctx.Query().Get("id") {
			method = MethodDeleteAll
		} else {
			method = MethodDelete
		}
	}

	return method
}
