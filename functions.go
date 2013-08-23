package restapi

import (
	"net/http"
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

func Call(apiName, method string, ctx IContext) IOutput {

	err := Filter.Emit("beforeCall", ctx)
	if nil != err {
		return Output(false, nil, []string{err.Error()})
	}

	api := _apis.Get(apiName)

	if nil == api {
		return nil
	}

	err = api.BeforeRun(ctx)
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
		err := Filter.Emit("beforeHandleRequest", ctx)
		if nil != err {
			http.Error(res, err.Error(), 404)
			return
		}

		if "POST" == req.Method || "PUT" == req.Method {
			req.ParseForm()
		}

		apiName, method, ok := Conf.RouterFunc(ctx)
		if !ok {
			return
		}

		if "" == apiName {
			http.Error(res, "API Not found", 404)
			return
		}

		if "" == method {
			http.Error(res, "Bad request method", 400)
			return
		}

		output := Call(apiName, method, ctx)
		if nil == output {
			http.Error(res, "API no response", 500)
			return
		}

		Conf.ResponseFunc(output, ctx)
	})

	http.ListenAndServe(bindString, nil)
}
