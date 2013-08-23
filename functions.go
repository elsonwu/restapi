package restapi

import (
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

func Call(apiName, method string, ctx IContext) IOutput {
	api := _apis.Get(apiName)

	if nil == api {
		return nil
	}

	err := api.BeforeRun(ctx)
	if nil != err {
		return Output(false, nil, []string{err.Error()})
	}

	var output IOutput
	switch method {
	case MethodView:
		err := api.BeforeView(ctx)
		if nil != err {
			return Output(false, nil, []string{err.Error()})
		}
		output = api.View(ctx)

	case MethodCreate:
		err := api.BeforeCreate(ctx)
		if nil != err {
			return Output(false, nil, []string{err.Error()})
		}
		output = api.Create(ctx)

	case MethodUpdate:
		err := api.BeforeUpdate(ctx)
		if nil != err {
			return Output(false, nil, []string{err.Error()})
		}
		output = api.Update(ctx)

	case MethodDelete:
		err := api.BeforeDelete(ctx)
		if nil != err {
			return Output(false, nil, []string{err.Error()})
		}
		output = api.Delete(ctx)

	case MethodList:
		err := api.BeforeList(ctx)
		if nil != err {
			return Output(false, nil, []string{err.Error()})
		}
		output = api.List(ctx)

	case MethodUpdateAll:
		err := api.BeforeUpdateAll(ctx)
		if nil != err {
			return Output(false, nil, []string{err.Error()})
		}
		output = api.UpdateAll(ctx)

	case MethodDeleteAll:
		err := api.BeforeDeleteAll(ctx)
		if nil != err {
			return Output(false, nil, []string{err.Error()})
		}
		output = api.DeleteAll(ctx)
	}

	return output
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

		Conf.ResponseFunc(output, ctx, res, req)
	})

	http.ListenAndServe(bindString, nil)
}
