package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	MethodView      string = "View"
	MethodCreate    string = "Create"
	MethodUpdate    string = "Update"
	MethodDelete    string = "Delete"
	MethodList      string = "List"
	MethodUpdateAll string = "UpdateAll"
	MethodDeleteAll string = "DeleteAll"
)

type IQuery interface {
	Add(string, string)
	Del(string)
	Get(string) string
	Set(string, string)
}

type Map map[string]interface{}

func requestApiMethod(req *http.Request, ctx IContext) string {
	var method string
	if "GET" == req.Method {
		if "" == ctx.Query().Get("id") {
			method = MethodList
		} else {
			method = MethodView
		}
	} else if "POST" == req.Method {
		method = MethodCreate
	} else if "PUT" == req.Method {
		if "" == ctx.Query().Get("id") {
			method = MethodUpdateAll
		} else {
			method = MethodUpdate
		}
	} else if "DELETE" == req.Method {
		if "" == ctx.Query().Get("id") {
			method = MethodDeleteAll
		} else {
			method = MethodDelete
		}
	}

	return method
}

func Get(apiName string) IApi {
	return _apis.Get(apiName)
}

func Call(apiName, method string, ctx IContext) IOutput {
	api := Get(apiName)

	if nil != api {
		if MethodView == method {
			return api.View(ctx)
		} else if MethodCreate == method {
			return api.Create(ctx)
		} else if MethodUpdate == method {
			return api.Update(ctx)
		} else if MethodDelete == method {
			return api.Delete(ctx)
		} else if MethodList == method {
			return api.List(ctx)
		} else if MethodUpdateAll == method {
			return api.UpdateAll(ctx)
		} else if MethodDeleteAll == method {
			return api.DeleteAll(ctx)
		} else {
			fmt.Println("API ", apiName, " method ", method, " does not exist")
		}
	}

	return nil
}

func Add(apiName string, api IApi) {
	_apis.Add(apiName, api)
}

func Run(bindString string) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		ctx := &Context{query: req.URL.Query(), req: req, res: res}
		if "POST" == req.Method || "PUT" == req.Method {
			req.ParseForm()
		}

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

		method := requestApiMethod(req, ctx)
		if "" == method {
			http.Error(res, "Request does not acceptable", 400)
			return
		}

		output := Call(paths[1], method, ctx)
		if nil == output {
			http.Error(res, "API Not found, api error:"+paths[1], 404)
			return
		}

		o := map[string]interface{}{}
		o["result"] = output.Result()
		o["data"] = output.Data()
		o["errors"] = output.Errors()
		data, _ := json.Marshal(o)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(data))
	})

	http.ListenAndServe(bindString, nil)
}
