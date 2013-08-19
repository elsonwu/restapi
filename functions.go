package restapi

import (
	"encoding/json"
	// "errors"
	"net/http"
	// "net/url"
	// "fmt"
	"strings"
	// "reflect"
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

type Params struct {
	Query IQuery
}

type Map map[string]interface{}

func bind(apiName string) {
	http.HandleFunc("/"+apiName+"/", func(res http.ResponseWriter, req *http.Request) {
		params := Params{}
		params.Query = req.URL.Query()
		if "POST" == req.Method || "PUT" == req.Method {
			req.ParseForm()
		}

		if "GET" == req.Method || "PUT" == req.Method || "DELETE" == req.Method {
			_id := strings.Trim(strings.TrimPrefix(req.URL.Path, "/"+apiName), "/")
			if "" != _id {
				params.Query.Add("id", _id)
			}
		}

		handler := new(Handler)
		handler.req = req
		handler.res = res

		method := requestApiMethod(req, params)
		if "" == method {
			http.Error(res, "Request does not acceptable", 400)
			return
		}

		output := handler.Call(apiName, method, params)
		if nil == output {
			http.Error(res, "API Not found", 404)
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
}

func requestApiMethod(req *http.Request, params Params) string {
	var method string
	if "GET" == req.Method {
		if "" == params.Query.Get("id") {
			method = MethodList
		} else {
			method = MethodView
		}
	} else if "POST" == req.Method {
		method = MethodCreate
	} else if "PUT" == req.Method {
		if "" == params.Query.Get("id") {
			method = MethodUpdateAll
		} else {
			method = MethodUpdate
		}
	} else if "DELETE" == req.Method {
		if "" == params.Query.Get("id") {
			method = MethodDeleteAll
		} else {
			method = MethodDelete
		}
	}

	return method
}

func Add(apiName string, api IApi) {
	_apis.Add(apiName, api)
	bind(apiName)
}

func Run(bindString string) {
	http.ListenAndServe(bindString, nil)
}
